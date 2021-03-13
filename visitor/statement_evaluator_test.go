package visitor

import (
	"errors"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rpj5582/sim/interpreter"
	"github.com/rpj5582/sim/parser"
	"github.com/stretchr/testify/assert"
)

type mockStatementContext struct {
	*parser.StatementContext
	resultFunc func() interface{}
	token      antlr.Token
}

func newMockStatementContext(resultFunc func() interface{}) *mockStatementContext {
	lexer := parser.NewSimLexer(antlr.NewInputStream(""))
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	token := tokenStream.GetTokenSource().NextToken()

	return &mockStatementContext{
		StatementContext: parser.NewEmptyStatementContext(),
		resultFunc:       resultFunc,
		token:            token,
	}
}

func (m *mockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	return m.resultFunc()
}

func (m *mockStatementContext) GetStart() antlr.Token {
	return m.token
}

func TestNewStatementEvaluator(t *testing.T) {
	expected := &StatementEvaluator{}
	evaluator := NewStatementEvaluator()
	assert.Equal(t, expected, evaluator)
}

func TestStatementEvaluatorEvaluate(t *testing.T) {
	statement := NewStatementEvaluator()

	t.Run("evaluate nil", func(t *testing.T) {
		controlFlow, err := statement.Evaluate(newMockVisitor(), newMockStatementContext(func() interface{} {
			return nil
		}))
		assert.Equal(t, ControlFlowNormal, controlFlow)
		assert.NoError(t, err)
	})

	t.Run("evaluate error", func(t *testing.T) {
		controlFlow, err := statement.Evaluate(newMockVisitor(), newMockStatementContext(func() interface{} {
			return errors.New("test error")
		}))
		assert.Equal(t, ControlFlowNormal, controlFlow)
		assert.EqualError(t, err, "test error")
	})

	t.Run("evaluate control flow", func(t *testing.T) {
		controlFlow, err := statement.Evaluate(newMockVisitor(), newMockStatementContext(func() interface{} {
			return ControlFlowBreak
		}))
		assert.Equal(t, ControlFlowBreak, controlFlow)
		assert.NoError(t, err)
	})

	t.Run("evaluate unknown statement result", func(t *testing.T) {
		controlFlow, err := statement.Evaluate(newMockVisitor(), newMockStatementContext(func() interface{} {
			return 2
		}))
		assert.Equal(t, ControlFlowNormal, controlFlow)
		assert.EqualError(t, err, interpreter.InvalidControlFlowErr{Context: interpreter.NewParseContext(1, 0), Data: 2}.Error())
	})
}
