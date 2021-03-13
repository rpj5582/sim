package visitor

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rpj5582/sim/interpreter"
	"github.com/rpj5582/sim/parser"
)

type SimVisitor struct {
	parser.BaseSimParserVisitor

	interpreter         *interpreter.SimInterpreter
	statementEvaluator  *StatementEvaluator
	expressionEvaluator *ExpressionEvaluator
}

func NewSimVisitor(interpreter *interpreter.SimInterpreter) *SimVisitor {
	return &SimVisitor{
		interpreter:         interpreter,
		statementEvaluator:  NewStatementEvaluator(),
		expressionEvaluator: NewExpressionEvaluator(),
	}
}

func (v *SimVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.StartContext:
		return val.Accept(v)
	}

	return nil
}

func (v *SimVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	statements := ctx.AllStatement()

	var controlFlow ControlFlow
	var err error

	for _, statement := range statements {
		controlFlow, err = v.statementEvaluator.Evaluate(v, statement)
		if err != nil {
			return err
		}

		if controlFlow != ControlFlowNormal {
			break
		}
	}

	return controlFlow
}

func (v *SimVisitor) VisitBlockStatement(ctx *parser.BlockStatementContext) (result interface{}) {
	parseContext := interpreter.NewParseContext(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

	statements := ctx.AllStatement()

	v.interpreter.PushScope()
	defer func() {
		if err := v.interpreter.PopScope(parseContext); err != nil {
			result = err
		}
	}()

	var controlFlow ControlFlow
	var err error

	for _, statement := range statements {
		controlFlow, err = v.statementEvaluator.Evaluate(v, statement)
		if err != nil {
			return err
		}

		if controlFlow != ControlFlowNormal {
			break
		}
	}

	return controlFlow
}

func (v *SimVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	expression := ctx.Expression()
	parseContext := interpreter.NewParseContext(expression.GetStart().GetLine(), expression.GetStart().GetColumn())

	condition, err := v.expressionEvaluator.Evaluate(parseContext, v, expression).GetBool(parseContext)
	if err != nil {
		return err
	}

	var controlFlow ControlFlow
	if condition {
		controlFlow, err = v.statementEvaluator.Evaluate(v, ctx.Statement())
		if err != nil {
			return err
		}
	}

	return controlFlow
}

func (v *SimVisitor) VisitInfiniteLoopStatement(ctx *parser.InfiniteLoopStatementContext) interface{} {
	var controlFlow ControlFlow
	var err error

	for {
		controlFlow, err = v.statementEvaluator.Evaluate(v, ctx.Statement())

		if err != nil {
			return err
		}

		if controlFlow == ControlFlowBreak || controlFlow == ControlFlowReturn {
			break
		}
	}

	return controlFlow
}

func (v *SimVisitor) VisitConditionalLoopStatement(ctx *parser.ConditionalLoopStatementContext) interface{} {
	expression := ctx.Expression()
	parseContext := interpreter.NewParseContext(expression.GetStart().GetLine(), expression.GetStart().GetColumn())

	var iterations uint32
	var condition bool
	var controlFlow ControlFlow
	var err error

	for {
		result := v.expressionEvaluator.Evaluate(parseContext, v, expression)
		condition, err = result.GetBool(parseContext)
		if err != nil {
			maxIterations, err := result.GetUint(parseContext)
			if err != nil {
				return err
			}

			condition = iterations < maxIterations
		}

		if !condition {
			break
		}

		controlFlow, err = v.statementEvaluator.Evaluate(v, ctx.Statement())
		if err != nil {
			return err
		}

		if controlFlow == ControlFlowBreak || controlFlow == ControlFlowReturn {
			break
		}

		iterations++
	}

	return controlFlow
}

func (v *SimVisitor) VisitLoopStatement(ctx *parser.LoopStatementContext) (result interface{}) {
	identifier := ctx.IDENTIFIER()
	minExpression := ctx.GetMin()
	maxExpression := ctx.GetMax()

	identifierContext := interpreter.NewParseContext(identifier.GetSymbol().GetLine(), identifier.GetSymbol().GetColumn())
	minParseContext := interpreter.NewParseContext(minExpression.GetStart().GetLine(), minExpression.GetStart().GetColumn())
	maxParseContext := interpreter.NewParseContext(maxExpression.GetStart().GetLine(), maxExpression.GetStart().GetColumn())

	varName := identifier.GetText()

	// If the variable name is present, min and max should also be present
	minVal := v.expressionEvaluator.Evaluate(minParseContext, v, minExpression)

	minTypeName, err := minVal.GetType()
	if err != nil {
		return err
	}

	min, err := minVal.GetUint(minParseContext)
	if err != nil {
		return err
	}

	max, err := v.expressionEvaluator.Evaluate(maxParseContext, v, maxExpression).GetUint(maxParseContext)
	if err != nil {
		return err
	}

	v.interpreter.PushScope()
	defer func() {
		if err := v.interpreter.PopScope(identifierContext); err != nil {
			result = err
		}
	}()

	// Loop iterators are always unsigned integers
	if err := v.interpreter.AddVar(minParseContext, interpreter.NewVariable(varName, interpreter.NewValue("uint", fmt.Sprintf("%d", min)))); err != nil {
		return err
	}

	for i := min; i < max; i++ {
		controlFlow, err := v.statementEvaluator.Evaluate(v, ctx.Statement())
		if err != nil {
			return err
		}

		if controlFlow == ControlFlowBreak || controlFlow == ControlFlowReturn {
			break
		}

		variable, err := v.interpreter.GetVar(identifierContext, varName)
		if err != nil {
			return err
		}

		result, err := v.interpreter.ResolveBinaryOperations(identifierContext, minParseContext, variable.Value(), interpreter.NewValue(minTypeName, "1"), "+")
		if err != nil {
			return err
		}

		if err := v.interpreter.SetVarValue(identifierContext, varName, result); err != nil {
			return err
		}
	}

	return nil
}

func (v *SimVisitor) VisitDeclarationStatement(ctx *parser.DeclarationStatementContext) interface{} {
	parseContext := interpreter.NewParseContext(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	expression := ctx.Expression()
	var expressionParseContext interpreter.ParseContext
	if expression != nil {
		expressionParseContext = interpreter.NewParseContext(expression.GetStart().GetLine(), expression.GetStart().GetColumn())
	}

	typeName := ctx.GetType_().GetText()
	typeData, err := v.interpreter.GetTypeData(parseContext, typeName)
	if err != nil {
		return err
	}

	parseContext.TypeData = typeData
	expressionParseContext.TypeData = typeData

	varName := ctx.GetVarName().GetText()

	value := interpreter.NewValue(typeName, "")

	if ctx.ASSIGNMENT() != nil {
		value = v.expressionEvaluator.Evaluate(expressionParseContext, v, expression)
	}

	variable := interpreter.NewVariable(varName, value)

	if err := v.interpreter.AddVar(parseContext, variable); err != nil {
		return err
	}

	return nil
}

func (v *SimVisitor) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
	expression := ctx.Expression()
	parseContext := interpreter.NewParseContext(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	expressionParseContext := interpreter.NewParseContext(expression.GetStart().GetLine(), expression.GetStart().GetColumn())

	varName := ctx.IDENTIFIER().GetText()

	variable, err := v.interpreter.GetVar(parseContext, varName)
	if err != nil {
		return err
	}

	result := v.expressionEvaluator.Evaluate(expressionParseContext, v, expression)

	token := ctx.Assignment_op().GetStart()

	if token.GetTokenType() != parser.SimParserASSIGNMENT {
		result, err = v.interpreter.ResolveBinaryOperations(parseContext, expressionParseContext, variable.Value(), result, token.GetText()[:1])
		if err != nil {
			return err
		}
	}

	if err := v.interpreter.SetVarValue(parseContext, varName, result); err != nil {
		return err
	}

	return nil
}

func (v *SimVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	expression := ctx.Expression()
	parseContext := interpreter.NewParseContext(expression.GetStart().GetLine(), expression.GetStart().GetColumn())

	result, err := v.expressionEvaluator.Evaluate(parseContext, v, expression).GetRawData()
	if err != nil {
		return err
	}

	v.interpreter.PrintLine(result)
	return nil
}

func (v *SimVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	return ControlFlowReturn
}

func (v *SimVisitor) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	return ControlFlowBreak
}

func (v *SimVisitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	return ControlFlowContinue
}

func (v *SimVisitor) VisitParensExpression(ctx *parser.ParensExpressionContext) interface{} {
	expression := ctx.Expression()
	parseContext := interpreter.NewParseContext(expression.GetStart().GetLine(), expression.GetStart().GetColumn())
	return v.expressionEvaluator.Evaluate(parseContext, v, expression)
}

func (v *SimVisitor) VisitNegateExpression(ctx *parser.NegateExpressionContext) interface{} {
	expression := ctx.Expression()
	parseContext := interpreter.NewParseContext(expression.GetStart().GetLine(), expression.GetStart().GetColumn())

	val := v.expressionEvaluator.Evaluate(parseContext, v, expression)

	result, err := v.interpreter.ResolveUnaryOperations(parseContext, val, ctx.SUBTRACT().GetText())
	if err != nil {
		return err
	}

	return result
}

func (v *SimVisitor) VisitNotExpression(ctx *parser.NotExpressionContext) interface{} {
	expression := ctx.Expression()
	parseContext := interpreter.NewParseContext(expression.GetStart().GetLine(), expression.GetStart().GetColumn())

	result, err := v.expressionEvaluator.Evaluate(parseContext, v, expression).GetBool(parseContext)
	if err != nil {
		return err
	}

	return fmt.Sprintf("%t", !result)
}

func (v *SimVisitor) VisitMulDivModExpression(ctx *parser.MulDivModExpressionContext) interface{} {
	leftExpression := ctx.GetLeft()
	leftParseContext := interpreter.NewParseContext(leftExpression.GetStart().GetLine(), leftExpression.GetStart().GetColumn())

	rightExpression := ctx.GetRight()
	rightParseContext := interpreter.NewParseContext(rightExpression.GetStart().GetLine(), rightExpression.GetStart().GetColumn())

	left := v.expressionEvaluator.Evaluate(leftParseContext, v, leftExpression)
	right := v.expressionEvaluator.Evaluate(rightParseContext, v, rightExpression)

	result, err := v.interpreter.ResolveBinaryOperations(leftParseContext, rightParseContext, left, right, ctx.GetOp().GetText())
	if err != nil {
		return err
	}

	return result
}

func (v *SimVisitor) VisitAddSubExpression(ctx *parser.AddSubExpressionContext) interface{} {
	leftExpression := ctx.GetLeft()
	leftParseContext := interpreter.NewParseContext(leftExpression.GetStart().GetLine(), leftExpression.GetStart().GetColumn())

	rightExpression := ctx.GetRight()
	rightParseContext := interpreter.NewParseContext(rightExpression.GetStart().GetLine(), rightExpression.GetStart().GetColumn())

	left := v.expressionEvaluator.Evaluate(leftParseContext, v, leftExpression)
	right := v.expressionEvaluator.Evaluate(rightParseContext, v, rightExpression)

	result, err := v.interpreter.ResolveBinaryOperations(leftParseContext, rightParseContext, left, right, ctx.GetOp().GetText())
	if err != nil {
		return err
	}

	return result
}

func (v *SimVisitor) VisitInequalityExpression(ctx *parser.InequalityExpressionContext) interface{} {
	leftExpression := ctx.GetLeft()
	leftParseContext := interpreter.NewParseContext(leftExpression.GetStart().GetLine(), leftExpression.GetStart().GetColumn())

	rightExpression := ctx.GetRight()
	rightParseContext := interpreter.NewParseContext(rightExpression.GetStart().GetLine(), rightExpression.GetStart().GetColumn())

	left := v.expressionEvaluator.Evaluate(leftParseContext, v, leftExpression)
	right := v.expressionEvaluator.Evaluate(rightParseContext, v, rightExpression)

	compare, err := v.interpreter.ResolveBinaryOperations(leftParseContext, rightParseContext, left, right, ctx.GetOp().GetText())
	if err != nil {
		return err
	}

	return compare
}

func (v *SimVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	leftExpression := ctx.GetLeft()
	leftParseContext := interpreter.NewParseContext(leftExpression.GetStart().GetLine(), leftExpression.GetStart().GetColumn())

	rightExpression := ctx.GetRight()
	rightParseContext := interpreter.NewParseContext(rightExpression.GetStart().GetLine(), rightExpression.GetStart().GetColumn())

	left := v.expressionEvaluator.Evaluate(leftParseContext, v, leftExpression)
	right := v.expressionEvaluator.Evaluate(rightParseContext, v, rightExpression)

	compare, err := v.interpreter.ResolveBinaryOperations(leftParseContext, rightParseContext, left, right, ctx.GetOp().GetText())
	if err != nil {
		return err
	}

	return compare
}

func (v *SimVisitor) VisitAndExpression(ctx *parser.AndExpressionContext) interface{} {
	leftExpression := ctx.GetLeft()
	leftParseContext := interpreter.NewParseContext(leftExpression.GetStart().GetLine(), leftExpression.GetStart().GetColumn())

	left, err := v.expressionEvaluator.Evaluate(leftParseContext, v, leftExpression).GetBool(leftParseContext)
	if err != nil {
		return err
	}

	if !left {
		return "false"
	}

	rightExpression := ctx.GetRight()
	rightParseContext := interpreter.NewParseContext(rightExpression.GetStart().GetLine(), rightExpression.GetStart().GetColumn())

	right, err := v.expressionEvaluator.Evaluate(rightParseContext, v, rightExpression).GetBool(rightParseContext)
	if err != nil {
		return err
	}

	if !right {
		return "false"
	}

	return "true"
}

func (v *SimVisitor) VisitOrExpression(ctx *parser.OrExpressionContext) interface{} {
	leftExpression := ctx.GetLeft()
	leftParseContext := interpreter.NewParseContext(leftExpression.GetStart().GetLine(), leftExpression.GetStart().GetColumn())

	left, err := v.expressionEvaluator.Evaluate(leftParseContext, v, leftExpression).GetBool(leftParseContext)
	if err != nil {
		return err
	}

	if left {
		return "true"
	}

	rightExpression := ctx.GetRight()
	rightParseContext := interpreter.NewParseContext(rightExpression.GetStart().GetLine(), rightExpression.GetStart().GetColumn())

	right, err := v.expressionEvaluator.Evaluate(rightParseContext, v, rightExpression).GetBool(rightParseContext)
	if err != nil {
		return err
	}

	if right {
		return "true"
	}

	return "false"
}

func (v *SimVisitor) VisitVariableExpression(ctx *parser.VariableExpressionContext) interface{} {
	parseContext := interpreter.NewParseContext(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

	varName := ctx.GetText()

	variable, err := v.interpreter.GetVar(parseContext, varName)
	if err != nil {
		return err
	}

	return variable.Value()
}

func (v *SimVisitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) interface{} {
	return ctx.GetText()
}
