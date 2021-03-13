package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rpj5582/sim/interpreter"
	"github.com/rpj5582/sim/parser"
)

// ExpressionEvaluator evaluates all expressions in a Sim program.
type ExpressionEvaluator struct{}

// NewExpressionEvaluator returns a new instance of an ExpressionEvaluator.
func NewExpressionEvaluator() *ExpressionEvaluator {
	return new(ExpressionEvaluator)
}

// Evaluate actually evaluates the expression by visiting the expression's children.
// The returned Value contains methods to access the data as the appropriate data type.
func (e *ExpressionEvaluator) Evaluate(context interpreter.ParseContext, visitor antlr.ParseTreeVisitor, ctx parser.IExpressionContext) interpreter.Value {
	expressionValue := ctx.Accept(visitor)

	// If the expression value is an error type, then pass that up.
	if err, ok := expressionValue.(error); ok {
		return interpreter.NewErrorValue(err)
	}

	// If the expression value is already a Value type, then pass that up.
	if value, ok := expressionValue.(interpreter.Value); ok {
		return value
	}

	// If the expression value is a string type, convert it to a Value then pass it up.
	if value, ok := expressionValue.(string); ok {
		return interpreter.NewValue(interpreter.GetTypeFromLiteral(context, value), value)
	}

	// If the expression value isn't a proper Value or string data, return an error.
	return interpreter.NewErrorValue(
		interpreter.InvalidValueErr{
			Context: context,
			Data:    expressionValue,
		},
	)
}
