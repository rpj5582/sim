package visitor

import (
	"errors"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rpj5582/sim/interpreter"
	"github.com/rpj5582/sim/parser"
	"github.com/stretchr/testify/assert"
)

type mockVisitor struct {
	antlr.BaseParseTreeVisitor
}

func newMockVisitor() *mockVisitor {
	return new(mockVisitor)
}

type mockExpressionContext struct {
	*parser.ExpressionContext
	resultFunc func() interface{}
}

func newMockExpressionContext(resultFunc func() interface{}) *mockExpressionContext {
	return &mockExpressionContext{
		ExpressionContext: parser.NewEmptyExpressionContext(),
		resultFunc:        resultFunc,
	}
}

func (m *mockExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	return m.resultFunc()
}

func TestNewExpressionEvaluator(t *testing.T) {
	expected := &ExpressionEvaluator{}
	evaluator := NewExpressionEvaluator()
	assert.Equal(t, expected, evaluator)
}

func TestExpressionEvaluatorEvaluate(t *testing.T) {
	evaluator := NewExpressionEvaluator()

	t.Run("evaluate an error", func(t *testing.T) {
		err := errors.New("test error")
		expected := interpreter.NewErrorValue(err)

		result := evaluator.Evaluate(interpreter.NewParseContext(0, 0), newMockVisitor(), newMockExpressionContext(func() interface{} {
			return err
		}))

		assert.Equal(t, expected, result)
	})

	t.Run("evaluate a value", func(t *testing.T) {
		expected := interpreter.NewValue("int", "10")

		result := evaluator.Evaluate(interpreter.NewParseContext(0, 0), newMockVisitor(), newMockExpressionContext(func() interface{} {
			return expected
		}))

		assert.Equal(t, expected, result)
	})

	t.Run("evaluate a data string", func(t *testing.T) {
		expected := interpreter.NewValue("int", "10")

		context := interpreter.NewParseContext(0, 0)
		context.TypeData = interpreter.NewTypeData("int", "0", interpreter.TypeInfoSignedInteger)
		result := evaluator.Evaluate(context, newMockVisitor(), newMockExpressionContext(func() interface{} {
			return "10"
		}))

		assert.Equal(t, expected, result)
	})

	t.Run("evaluate unknown type", func(t *testing.T) {
		expected := interpreter.NewErrorValue(interpreter.InvalidValueErr{Data: 10})

		result := evaluator.Evaluate(interpreter.NewParseContext(0, 0), newMockVisitor(), newMockExpressionContext(func() interface{} {
			return 10
		}))

		assert.Equal(t, expected, result)
	})
}
