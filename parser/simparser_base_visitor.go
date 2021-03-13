// Code generated from /home/ryanj/sim/SimParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SimParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSimParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimParserVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitInfiniteLoopStatement(ctx *InfiniteLoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitConditionalLoopStatement(ctx *ConditionalLoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitFunctionStatement(ctx *FunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitDeclarationStatement(ctx *DeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitNegateExpression(ctx *NegateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitAddSubExpression(ctx *AddSubExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitInequalityExpression(ctx *InequalityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitVariableExpression(ctx *VariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitOrExpression(ctx *OrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitParensExpression(ctx *ParensExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitMulDivModExpression(ctx *MulDivModExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitAssignment_op(ctx *Assignment_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}
