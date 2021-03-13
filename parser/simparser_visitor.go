// Code generated from /home/ryanj/sim/SimParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SimParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SimParser.
type SimParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by SimParser#BlockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by SimParser#IfStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by SimParser#InfiniteLoopStatement.
	VisitInfiniteLoopStatement(ctx *InfiniteLoopStatementContext) interface{}

	// Visit a parse tree produced by SimParser#ConditionalLoopStatement.
	VisitConditionalLoopStatement(ctx *ConditionalLoopStatementContext) interface{}

	// Visit a parse tree produced by SimParser#LoopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by SimParser#FunctionStatement.
	VisitFunctionStatement(ctx *FunctionStatementContext) interface{}

	// Visit a parse tree produced by SimParser#DeclarationStatement.
	VisitDeclarationStatement(ctx *DeclarationStatementContext) interface{}

	// Visit a parse tree produced by SimParser#AssignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by SimParser#ReturnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by SimParser#PrintStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by SimParser#BreakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by SimParser#ContinueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by SimParser#NegateExpression.
	VisitNegateExpression(ctx *NegateExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#AddSubExpression.
	VisitAddSubExpression(ctx *AddSubExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#InequalityExpression.
	VisitInequalityExpression(ctx *InequalityExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#AndExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#NotExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#VariableExpression.
	VisitVariableExpression(ctx *VariableExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#EqualityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#OrExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#ParensExpression.
	VisitParensExpression(ctx *ParensExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#MulDivModExpression.
	VisitMulDivModExpression(ctx *MulDivModExpressionContext) interface{}

	// Visit a parse tree produced by SimParser#assignment_op.
	VisitAssignment_op(ctx *Assignment_opContext) interface{}

	// Visit a parse tree produced by SimParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
