package visitor

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rpj5582/sim/interpreter"
	"github.com/rpj5582/sim/listener"
	"github.com/rpj5582/sim/parser"
	"github.com/stretchr/testify/assert"
)

func TestVisitBlockStatement(t *testing.T) {
	input := `{
		int a = 10
		int b = a
	}

	bool c = true`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"c": interpreter.NewVariable("c", interpreter.NewValue("bool", "true")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestIfStatement(t *testing.T) {
	input := `int a

	if 20 > 10 
	{
		a = 10
	}`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "10")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestInfiniteLoopStatement(t *testing.T) {
	input := `int a
	
	loop
	{
		if a >= 5
			break
		
		if a % 2 == 1 {
			a += 1
			continue
		}

		print(a)
		a += 1
	}`

	var buf bytes.Buffer
	simInterpreter := interpreter.NewSimInterpreter(&buf)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "5")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, "0\n2\n4\n", buf.String())
	assert.Equal(t, expectedVars, vars)
}

func TestConditionalLoopStatement(t *testing.T) {
	input := `int a
	loop a < 5
	{
		if a % 2 == 1 {
			a += 1
			continue
		}

		print(a)
		a += 1
	}`

	var buf bytes.Buffer
	simInterpreter := interpreter.NewSimInterpreter(&buf)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "5")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, "0\n2\n4\n", buf.String())
	assert.Equal(t, expectedVars, vars)
}

func TestLoopStatement(t *testing.T) {
	input := `loop a = 0 to 5
	{
		if a % 2 == 1 {
			continue
		}
		
		print(a)
	}`

	var buf bytes.Buffer
	simInterpreter := interpreter.NewSimInterpreter(&buf)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, "0\n2\n4\n", buf.String())
	assert.Empty(t, vars)
}

func TestVisitDeclarationStatement(t *testing.T) {
	input := `int a = 10
	int b = a`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "10")),
		"b": interpreter.NewVariable("b", interpreter.NewValue("int", "10")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitAssignmentStatement(t *testing.T) {
	input := `int a = 10
	a = 20
	a += 10
	a -= 5
	a *= 5
	a /= 25`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "5")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitPrintStatement(t *testing.T) {
	input := `print(10)
	
	int a = 20
	print(a)`

	var buf bytes.Buffer
	simInterpreter := interpreter.NewSimInterpreter(&buf)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "20")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
	assert.Equal(t, "10\n20\n", buf.String())
}

func TestVisitParensExpression(t *testing.T) {
	input := `int a = (10 * 20)`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "200")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitNegateExpression(t *testing.T) {
	input := `int a = 10
	int b = -a`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "10")),
		"b": interpreter.NewVariable("b", interpreter.NewValue("int", "-10")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitNotExpression(t *testing.T) {
	input := `bool a = false
	bool b = not a`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("bool", "false")),
		"b": interpreter.NewVariable("b", interpreter.NewValue("bool", "true")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitMulDivExpression(t *testing.T) {
	input := `int a = 10 * 20
	a = a / 5`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "40")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitAddSubExpression(t *testing.T) {
	input := `int a = 10 + 20
	a = a - 5`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "25")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitInequalityExpression(t *testing.T) {
	t.Run("mismatched types", func(t *testing.T) {
		input := `int a = 10
		float b = 20
		
		bool c = a > b`

		simInterpreter := interpreter.NewSimInterpreter(nil)

		err := walkTree(t, input, simInterpreter)
		assert.EqualError(t, err, interpreter.InvalidOperationErr{Context: interpreter.NewParseContext(4, 11), TypeNames: []string{"int", "float"}}.Error())

		expectedVars := map[string]interpreter.Variable{
			"a": interpreter.NewVariable("a", interpreter.NewValue("int", "10")),
			"b": interpreter.NewVariable("b", interpreter.NewValue("float", "20")),
		}

		vars := simInterpreter.GetAllVars()

		assert.Equal(t, expectedVars, vars)
	})

	input := `int a = 10
	int b = 20
	int c = 10
	
	bool d = a > b
	bool e = a < b
	bool f = a >= b
	bool g = a <= b

	bool h = a > c
	bool i = a < c
	bool j = a >= c
	bool k = a <= c`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "10")),
		"b": interpreter.NewVariable("b", interpreter.NewValue("int", "20")),
		"c": interpreter.NewVariable("c", interpreter.NewValue("int", "10")),
		"d": interpreter.NewVariable("d", interpreter.NewValue("bool", "false")),
		"e": interpreter.NewVariable("e", interpreter.NewValue("bool", "true")),
		"f": interpreter.NewVariable("f", interpreter.NewValue("bool", "false")),
		"g": interpreter.NewVariable("g", interpreter.NewValue("bool", "true")),
		"h": interpreter.NewVariable("h", interpreter.NewValue("bool", "false")),
		"i": interpreter.NewVariable("i", interpreter.NewValue("bool", "false")),
		"j": interpreter.NewVariable("j", interpreter.NewValue("bool", "true")),
		"k": interpreter.NewVariable("k", interpreter.NewValue("bool", "true")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitEqualityExpression(t *testing.T) {
	operandValues := []string{"0", "0.0", "false"}
	operandTypes := []string{"untyped int", "untyped float", "bool"}

	for i := range operandValues {
		for j := range operandValues {
			leftVal := operandValues[i]
			rightVal := operandValues[j]

			leftType := operandTypes[i]
			rightType := operandTypes[j]

			if i == j || (strings.Contains(leftType, "untyped") && strings.Contains(rightType, "untyped")) {
				input := fmt.Sprintf("bool a = %s == %s", leftVal, rightVal)

				simInterpreter := interpreter.NewSimInterpreter(nil)

				err := walkTree(t, input, simInterpreter)
				assert.NoError(t, err)

				expectedVars := map[string]interpreter.Variable{
					"a": interpreter.NewVariable("a", interpreter.NewValue("bool", "true")),
				}

				vars := simInterpreter.GetAllVars()
				assert.Equal(t, expectedVars, vars)
				continue
			}

			input := fmt.Sprintf("bool a = %s == %s", leftVal, rightVal)

			simInterpreter := interpreter.NewSimInterpreter(nil)

			err := walkTree(t, input, simInterpreter)
			assert.EqualError(t, err, interpreter.InvalidOperationErr{Context: interpreter.NewParseContext(1, 9), TypeNames: []string{leftType, rightType}}.Error())

			vars := simInterpreter.GetAllVars()
			assert.Empty(t, vars)
		}
	}

	// TODO: more types, custom types
}

func TestVisitAndExpression(t *testing.T) {
	input := `bool a = false and false
	bool b = false and true
	bool c = true and false
	bool d = true and true`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("bool", "false")),
		"b": interpreter.NewVariable("b", interpreter.NewValue("bool", "false")),
		"c": interpreter.NewVariable("c", interpreter.NewValue("bool", "false")),
		"d": interpreter.NewVariable("d", interpreter.NewValue("bool", "true")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitOrExpression(t *testing.T) {
	input := `bool a = false or false
	bool b = false or true
	bool c = true or false
	bool d = true or true`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("bool", "false")),
		"b": interpreter.NewVariable("b", interpreter.NewValue("bool", "true")),
		"c": interpreter.NewVariable("c", interpreter.NewValue("bool", "true")),
		"d": interpreter.NewVariable("d", interpreter.NewValue("bool", "true")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitVariableExpression(t *testing.T) {
	input := `int a = 10
	int b = a`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "10")),
		"b": interpreter.NewVariable("b", interpreter.NewValue("int", "10")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func TestVisitLiteralExpression(t *testing.T) {
	input := `int a = 10`

	simInterpreter := interpreter.NewSimInterpreter(nil)

	err := walkTree(t, input, simInterpreter)
	assert.NoError(t, err)

	expectedVars := map[string]interpreter.Variable{
		"a": interpreter.NewVariable("a", interpreter.NewValue("int", "10")),
	}

	vars := simInterpreter.GetAllVars()

	assert.Equal(t, expectedVars, vars)
}

func walkTree(t *testing.T, input string, interpreter *interpreter.SimInterpreter) error {
	inputStream := antlr.NewInputStream(input)

	lexer := parser.NewSimLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSimParser(stream)

	errListener := listener.NewSimErrorListener(antlr.NewDiagnosticErrorListener(false))
	p.AddErrorListener(errListener)

	visitor := NewSimVisitor(interpreter)

	tree := p.Start()

	if err := errListener.GetError(); err != nil {
		return err
	}

	if err, ok := visitor.Visit(tree).(error); ok && err != nil {
		return err
	}

	return nil
}
