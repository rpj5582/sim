package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rpj5582/sim/interpreter"
	"github.com/rpj5582/sim/parser"
)

// ControlFlow is an enum defining what the control flow is once a statement is evaluated.
// This is useful for statements that cause jumps, such as return, break, and continue.
type ControlFlow int

const (
	// ControlFlowNormal control flow. Process the next statement as normal.
	ControlFlowNormal ControlFlow = 0

	// ControlFlowReturn from the current function.
	ControlFlowReturn ControlFlow = 1

	// ControlFlowBreak from the current loop.
	ControlFlowBreak ControlFlow = 2

	// ControlFlowContinue to the next iteration of the current loop.
	ControlFlowContinue ControlFlow = 3
)

// StatementEvaluator evaluates all statements in a Sim program.
type StatementEvaluator struct{}

// NewStatementEvaluator returns a new instance of a StatementEvaluator.
func NewStatementEvaluator() *StatementEvaluator {
	return new(StatementEvaluator)
}

// Evaluate actually evaluates the statement by visiting the statement's children,
// breaking out any errors that may have been generated into a separate return variable.
// The returned control flow describes how the caller should handle following statements.
func (s *StatementEvaluator) Evaluate(visitor antlr.ParseTreeVisitor, ctx parser.IStatementContext) (ControlFlow, error) {
	result := ctx.Accept(visitor)
	if result == nil {
		return ControlFlowNormal, nil
	}

	if err, ok := result.(error); ok && err != nil {
		return ControlFlowNormal, err
	}

	if controlFlow, ok := result.(ControlFlow); ok {
		return controlFlow, nil
	}

	start := ctx.GetStart()
	line := start.GetLine()
	column := start.GetColumn()

	return ControlFlowNormal, interpreter.InvalidControlFlowErr{
		Context: interpreter.NewParseContext(line, column),
		Data:    result,
	}
}
