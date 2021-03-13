// Code generated from /home/ryanj/sim/SimParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SimParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimParserListener is a complete listener for a parse tree produced by SimParser.
type SimParserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterBlockStatement is called when entering the BlockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterIfStatement is called when entering the IfStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterInfiniteLoopStatement is called when entering the InfiniteLoopStatement production.
	EnterInfiniteLoopStatement(c *InfiniteLoopStatementContext)

	// EnterConditionalLoopStatement is called when entering the ConditionalLoopStatement production.
	EnterConditionalLoopStatement(c *ConditionalLoopStatementContext)

	// EnterLoopStatement is called when entering the LoopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterFunctionStatement is called when entering the FunctionStatement production.
	EnterFunctionStatement(c *FunctionStatementContext)

	// EnterDeclarationStatement is called when entering the DeclarationStatement production.
	EnterDeclarationStatement(c *DeclarationStatementContext)

	// EnterAssignmentStatement is called when entering the AssignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterReturnStatement is called when entering the ReturnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterPrintStatement is called when entering the PrintStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterBreakStatement is called when entering the BreakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the ContinueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterNegateExpression is called when entering the NegateExpression production.
	EnterNegateExpression(c *NegateExpressionContext)

	// EnterAddSubExpression is called when entering the AddSubExpression production.
	EnterAddSubExpression(c *AddSubExpressionContext)

	// EnterInequalityExpression is called when entering the InequalityExpression production.
	EnterInequalityExpression(c *InequalityExpressionContext)

	// EnterAndExpression is called when entering the AndExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterNotExpression is called when entering the NotExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterVariableExpression is called when entering the VariableExpression production.
	EnterVariableExpression(c *VariableExpressionContext)

	// EnterEqualityExpression is called when entering the EqualityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterOrExpression is called when entering the OrExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterParensExpression is called when entering the ParensExpression production.
	EnterParensExpression(c *ParensExpressionContext)

	// EnterMulDivModExpression is called when entering the MulDivModExpression production.
	EnterMulDivModExpression(c *MulDivModExpressionContext)

	// EnterAssignment_op is called when entering the assignment_op production.
	EnterAssignment_op(c *Assignment_opContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitBlockStatement is called when exiting the BlockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitIfStatement is called when exiting the IfStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitInfiniteLoopStatement is called when exiting the InfiniteLoopStatement production.
	ExitInfiniteLoopStatement(c *InfiniteLoopStatementContext)

	// ExitConditionalLoopStatement is called when exiting the ConditionalLoopStatement production.
	ExitConditionalLoopStatement(c *ConditionalLoopStatementContext)

	// ExitLoopStatement is called when exiting the LoopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitFunctionStatement is called when exiting the FunctionStatement production.
	ExitFunctionStatement(c *FunctionStatementContext)

	// ExitDeclarationStatement is called when exiting the DeclarationStatement production.
	ExitDeclarationStatement(c *DeclarationStatementContext)

	// ExitAssignmentStatement is called when exiting the AssignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitReturnStatement is called when exiting the ReturnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitPrintStatement is called when exiting the PrintStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitBreakStatement is called when exiting the BreakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the ContinueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitNegateExpression is called when exiting the NegateExpression production.
	ExitNegateExpression(c *NegateExpressionContext)

	// ExitAddSubExpression is called when exiting the AddSubExpression production.
	ExitAddSubExpression(c *AddSubExpressionContext)

	// ExitInequalityExpression is called when exiting the InequalityExpression production.
	ExitInequalityExpression(c *InequalityExpressionContext)

	// ExitAndExpression is called when exiting the AndExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitNotExpression is called when exiting the NotExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitVariableExpression is called when exiting the VariableExpression production.
	ExitVariableExpression(c *VariableExpressionContext)

	// ExitEqualityExpression is called when exiting the EqualityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitOrExpression is called when exiting the OrExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitParensExpression is called when exiting the ParensExpression production.
	ExitParensExpression(c *ParensExpressionContext)

	// ExitMulDivModExpression is called when exiting the MulDivModExpression production.
	ExitMulDivModExpression(c *MulDivModExpressionContext)

	// ExitAssignment_op is called when exiting the assignment_op production.
	ExitAssignment_op(c *Assignment_opContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
