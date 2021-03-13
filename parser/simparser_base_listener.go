// Code generated from /home/ryanj/sim/SimParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SimParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimParserListener is a complete listener for a parse tree produced by SimParser.
type BaseSimParserListener struct{}

var _ SimParserListener = &BaseSimParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSimParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSimParserListener) ExitStart(ctx *StartContext) {}

// EnterBlockStatement is called when production BlockStatement is entered.
func (s *BaseSimParserListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production BlockStatement is exited.
func (s *BaseSimParserListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterIfStatement is called when production IfStatement is entered.
func (s *BaseSimParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production IfStatement is exited.
func (s *BaseSimParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterInfiniteLoopStatement is called when production InfiniteLoopStatement is entered.
func (s *BaseSimParserListener) EnterInfiniteLoopStatement(ctx *InfiniteLoopStatementContext) {}

// ExitInfiniteLoopStatement is called when production InfiniteLoopStatement is exited.
func (s *BaseSimParserListener) ExitInfiniteLoopStatement(ctx *InfiniteLoopStatementContext) {}

// EnterConditionalLoopStatement is called when production ConditionalLoopStatement is entered.
func (s *BaseSimParserListener) EnterConditionalLoopStatement(ctx *ConditionalLoopStatementContext) {}

// ExitConditionalLoopStatement is called when production ConditionalLoopStatement is exited.
func (s *BaseSimParserListener) ExitConditionalLoopStatement(ctx *ConditionalLoopStatementContext) {}

// EnterLoopStatement is called when production LoopStatement is entered.
func (s *BaseSimParserListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production LoopStatement is exited.
func (s *BaseSimParserListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterFunctionStatement is called when production FunctionStatement is entered.
func (s *BaseSimParserListener) EnterFunctionStatement(ctx *FunctionStatementContext) {}

// ExitFunctionStatement is called when production FunctionStatement is exited.
func (s *BaseSimParserListener) ExitFunctionStatement(ctx *FunctionStatementContext) {}

// EnterDeclarationStatement is called when production DeclarationStatement is entered.
func (s *BaseSimParserListener) EnterDeclarationStatement(ctx *DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production DeclarationStatement is exited.
func (s *BaseSimParserListener) ExitDeclarationStatement(ctx *DeclarationStatementContext) {}

// EnterAssignmentStatement is called when production AssignmentStatement is entered.
func (s *BaseSimParserListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production AssignmentStatement is exited.
func (s *BaseSimParserListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterReturnStatement is called when production ReturnStatement is entered.
func (s *BaseSimParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production ReturnStatement is exited.
func (s *BaseSimParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterPrintStatement is called when production PrintStatement is entered.
func (s *BaseSimParserListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production PrintStatement is exited.
func (s *BaseSimParserListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterBreakStatement is called when production BreakStatement is entered.
func (s *BaseSimParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production BreakStatement is exited.
func (s *BaseSimParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production ContinueStatement is entered.
func (s *BaseSimParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production ContinueStatement is exited.
func (s *BaseSimParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterNegateExpression is called when production NegateExpression is entered.
func (s *BaseSimParserListener) EnterNegateExpression(ctx *NegateExpressionContext) {}

// ExitNegateExpression is called when production NegateExpression is exited.
func (s *BaseSimParserListener) ExitNegateExpression(ctx *NegateExpressionContext) {}

// EnterAddSubExpression is called when production AddSubExpression is entered.
func (s *BaseSimParserListener) EnterAddSubExpression(ctx *AddSubExpressionContext) {}

// ExitAddSubExpression is called when production AddSubExpression is exited.
func (s *BaseSimParserListener) ExitAddSubExpression(ctx *AddSubExpressionContext) {}

// EnterInequalityExpression is called when production InequalityExpression is entered.
func (s *BaseSimParserListener) EnterInequalityExpression(ctx *InequalityExpressionContext) {}

// ExitInequalityExpression is called when production InequalityExpression is exited.
func (s *BaseSimParserListener) ExitInequalityExpression(ctx *InequalityExpressionContext) {}

// EnterAndExpression is called when production AndExpression is entered.
func (s *BaseSimParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production AndExpression is exited.
func (s *BaseSimParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseSimParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseSimParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterNotExpression is called when production NotExpression is entered.
func (s *BaseSimParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production NotExpression is exited.
func (s *BaseSimParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterVariableExpression is called when production VariableExpression is entered.
func (s *BaseSimParserListener) EnterVariableExpression(ctx *VariableExpressionContext) {}

// ExitVariableExpression is called when production VariableExpression is exited.
func (s *BaseSimParserListener) ExitVariableExpression(ctx *VariableExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseSimParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseSimParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterOrExpression is called when production OrExpression is entered.
func (s *BaseSimParserListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production OrExpression is exited.
func (s *BaseSimParserListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterParensExpression is called when production ParensExpression is entered.
func (s *BaseSimParserListener) EnterParensExpression(ctx *ParensExpressionContext) {}

// ExitParensExpression is called when production ParensExpression is exited.
func (s *BaseSimParserListener) ExitParensExpression(ctx *ParensExpressionContext) {}

// EnterMulDivModExpression is called when production MulDivModExpression is entered.
func (s *BaseSimParserListener) EnterMulDivModExpression(ctx *MulDivModExpressionContext) {}

// ExitMulDivModExpression is called when production MulDivModExpression is exited.
func (s *BaseSimParserListener) ExitMulDivModExpression(ctx *MulDivModExpressionContext) {}

// EnterAssignment_op is called when production assignment_op is entered.
func (s *BaseSimParserListener) EnterAssignment_op(ctx *Assignment_opContext) {}

// ExitAssignment_op is called when production assignment_op is exited.
func (s *BaseSimParserListener) ExitAssignment_op(ctx *Assignment_opContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseSimParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseSimParserListener) ExitEos(ctx *EosContext) {}
