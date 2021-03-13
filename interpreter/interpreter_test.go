package interpreter

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func runTestInterpreterNoValue(t *testing.T, typeName string) {
	var buf bytes.Buffer
	interpreter := NewSimInterpreter(&buf)

	a := NewVariable("a", NewValue(typeName, ""))

	err := interpreter.AddVar(NewParseContext(0, 0), a)
	assert.NoError(t, err)

	expectedTypes := getBasicTypes()
	a.value.data = expectedTypes[typeName].zeroValue.data

	expectedVars := map[string]Variable{a.name: a}
	expectedScopes := []*scope{{[]string{a.name}}}

	assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, expectedTypes, "")
}

func assertInterpreterValues(t *testing.T, interpreter *SimInterpreter, output string, expectedVars map[string]Variable, expectedScopes []*scope, expectedTypes map[string]TypeData, expectedOutput string) {
	assert.Equal(t, expectedVars, interpreter.vars)
	assert.Equal(t, expectedScopes, interpreter.scopes)
	assert.Equal(t, expectedTypes, interpreter.types)
	assert.Equal(t, expectedOutput, output)
}

func TestNewSimInterpreter(t *testing.T) {
	var buf bytes.Buffer
	interpreter := NewSimInterpreter(&buf)
	assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
}

func TestInterpreterGetTypeData(t *testing.T) {
	var buf bytes.Buffer
	interpreter := NewSimInterpreter(&buf)
	context := NewParseContext(0, 0)

	t.Run("unknown type", func(t *testing.T) {
		typeData, err := interpreter.GetTypeData(context, "unknown")
		assert.Empty(t, typeData)
		assert.EqualError(t, err, UnknownTypeErr{TypeName: "unknown"}.Error())

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})

	t.Run("success", func(t *testing.T) {
		typeData, err := interpreter.GetTypeData(context, "bool")

		assert.Equal(t, NewTypeData("bool", "false", TypeInfoBool), typeData)
		assert.NoError(t, err)

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})
}

func TestInterpreterPushScope(t *testing.T) {
	var buf bytes.Buffer
	interpreter := NewSimInterpreter(&buf)

	interpreter.PushScope()

	assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}, {}}, getBasicTypes(), "")
}

func TestInterpreterPopScope(t *testing.T) {
	var buf bytes.Buffer
	interpreter := NewSimInterpreter(&buf)
	context := NewParseContext(0, 0)

	t.Run("pop global scope", func(t *testing.T) {
		err := interpreter.PopScope(context)
		assert.EqualError(t, err, ExitGlobalScopeErr{}.Error())

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})

	t.Run("success", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("int", "10"))
		b := NewVariable("b", NewValue("int", "10"))

		interpreter.PushScope()
		err := interpreter.AddVar(context, a)
		assert.NoError(t, err)

		interpreter.PushScope()
		err = interpreter.AddVar(context, b)
		assert.NoError(t, err)

		err = interpreter.PopScope(context)
		assert.NoError(t, err)

		expectedVars := map[string]Variable{a.name: a}
		expectedScopes := []*scope{{}, {[]string{a.name}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
	})
}

func TestInterpreterAddVar(t *testing.T) {
	context := NewParseContext(0, 0)

	t.Run("error", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		err := interpreter.AddVar(context, NewVariable("a", NewErrorValue(errors.New("test error"))))
		assert.EqualError(t, err, "test error")

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})

	t.Run("untyped int", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		err := interpreter.AddVar(context, NewVariable("a", NewValue("untyped int", "10")))
		assert.NoError(t, err)

		expectedVariable := NewVariable("a", NewValue("int", "10"))

		expectedVars := map[string]Variable{expectedVariable.name: expectedVariable}
		expectedScopes := []*scope{{[]string{"a"}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
	})

	t.Run("untyped float", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		err := interpreter.AddVar(context, NewVariable("a", NewValue("untyped float", "10")))
		assert.NoError(t, err)

		expectedVariable := NewVariable("a", NewValue("float", "10"))

		expectedVars := map[string]Variable{expectedVariable.name: expectedVariable}
		expectedScopes := []*scope{{[]string{"a"}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
	})

	t.Run("unknown type", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		err := interpreter.AddVar(context, NewVariable("a", NewValue("unknown-type", "10")))
		assert.EqualError(t, err, UnknownTypeErr{TypeName: "unknown-type"}.Error())

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})

	t.Run("var exists", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("int", "10"))

		err := interpreter.AddVar(context, a)
		assert.NoError(t, err)

		err = interpreter.AddVar(context, a)
		assert.EqualError(t, err, VarExistsErr{VarName: a.name}.Error())

		expectedVars := map[string]Variable{a.name: a}
		expectedScopes := []*scope{{[]string{a.name}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
	})

	t.Run("mismatched type", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("bool", "10"))

		err := interpreter.AddVar(context, a)
		assert.EqualError(t, err, MismatchedTypeAssignErr{Var: a}.Error())

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})

	t.Run("success", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("int", "10"))

		err := interpreter.AddVar(context, a)
		assert.NoError(t, err)

		expectedVars := map[string]Variable{a.name: a}
		expectedScopes := []*scope{{[]string{"a"}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
	})

	t.Run("nil value", func(t *testing.T) {
		runTestInterpreterNoValue(t, "int")
		runTestInterpreterNoValue(t, "int8")
		runTestInterpreterNoValue(t, "int16")
		runTestInterpreterNoValue(t, "int32")
		runTestInterpreterNoValue(t, "int64")

		runTestInterpreterNoValue(t, "byte")
		runTestInterpreterNoValue(t, "uint")
		runTestInterpreterNoValue(t, "uint8")
		runTestInterpreterNoValue(t, "uint16")
		runTestInterpreterNoValue(t, "uint32")
		runTestInterpreterNoValue(t, "uint64")

		runTestInterpreterNoValue(t, "float")
		runTestInterpreterNoValue(t, "float32")
		runTestInterpreterNoValue(t, "float64")

		runTestInterpreterNoValue(t, "bool")

		runTestInterpreterNoValue(t, "string")
	})
}

func TestInterpreterGetVar(t *testing.T) {
	context := NewParseContext(0, 0)

	t.Run("unknown var", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		v, err := interpreter.GetVar(context, "a")
		assert.Empty(t, v)
		assert.EqualError(t, err, UnknownVarErr{VarName: "a"}.Error())

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})

	t.Run("success", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("int", "10"))

		err := interpreter.AddVar(context, a)
		assert.NoError(t, err)

		v, err := interpreter.GetVar(context, a.name)
		assert.NoError(t, err)
		assert.Equal(t, a, v)

		expectedVars := map[string]Variable{a.name: a}
		expectedScope := []*scope{{[]string{"a"}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScope, getBasicTypes(), "")
	})
}

func TestInterpreterSetVarValue(t *testing.T) {
	context := NewParseContext(0, 0)

	t.Run("unknown var", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		err := interpreter.SetVarValue(context, "a", NewValue("int", "10"))
		assert.EqualError(t, err, UnknownVarErr{VarName: "a"}.Error())

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})

	t.Run("unknown type", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("int", "10"))

		err := interpreter.AddVar(context, a)
		assert.NoError(t, err)

		// delete the variable's type after it has been added to simulate this error condition
		delete(interpreter.types, a.value.typeName)

		err = interpreter.SetVarValue(context, a.name, NewValue(a.value.typeName, "20"))
		assert.EqualError(t, err, InvalidTypeErr{TypeName: a.value.typeName, VarName: a.name}.Error())

		expectedVars := map[string]Variable{a.name: a}
		expectedScopes := []*scope{{[]string{"a"}}}

		expectedTypes := getBasicTypes()
		delete(expectedTypes, a.value.typeName)

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, expectedTypes, "")
	})

	t.Run("variable and value have mismatched types", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("int", "10"))

		err := interpreter.AddVar(context, a)
		assert.NoError(t, err)

		err = interpreter.SetVarValue(context, a.name, NewValue("bool", "true"))
		assert.EqualError(t, err, MismatchedTypeAssignErr{Var: a}.Error())

		expectedVars := map[string]Variable{a.name: a}
		expectedScopes := []*scope{{[]string{"a"}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
	})

	t.Run("value and literal have mismatched types", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("int", "10"))

		err := interpreter.AddVar(context, a)
		assert.NoError(t, err)

		err = interpreter.SetVarValue(context, a.name, NewValue(a.value.typeName, "true"))
		assert.EqualError(t, err, MismatchedTypeAssignErr{Var: a}.Error())

		expectedVars := map[string]Variable{a.name: a}
		expectedScopes := []*scope{{[]string{"a"}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
	})

	t.Run("success", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		a := NewVariable("a", NewValue("int", "10"))

		err := interpreter.AddVar(context, a)
		assert.NoError(t, err)

		newAVal := NewValue(a.value.typeName, "20")

		err = interpreter.SetVarValue(context, a.name, newAVal)
		assert.NoError(t, err)

		expectedVars := map[string]Variable{a.name: NewVariable(a.name, newAVal)}
		expectedScopes := []*scope{{[]string{"a"}}}

		assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
	})
}

func TestInterpreterPrintLine(t *testing.T) {
	var buf bytes.Buffer
	interpreter := NewSimInterpreter(&buf)

	interpreter.PrintLine("test")

	assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "test\n")
}

func TestInterpreterGetAllVars(t *testing.T) {
	var buf bytes.Buffer
	interpreter := NewSimInterpreter(&buf)
	context := NewParseContext(0, 0)

	a := NewVariable("a", NewValue("int", "10"))
	b := NewVariable("b", NewValue("int", "20"))

	err := interpreter.AddVar(context, a)
	assert.NoError(t, err)
	err = interpreter.AddVar(context, b)
	assert.NoError(t, err)

	vars := interpreter.GetAllVars()

	expectedVars := map[string]Variable{a.name: a, b.name: b}
	expectedScopes := []*scope{{[]string{"a", "b"}}}

	assert.Equal(t, expectedVars, vars)

	assertInterpreterValues(t, interpreter, buf.String(), expectedVars, expectedScopes, getBasicTypes(), "")
}

func TestInterpreterResolveUnaryOperations(t *testing.T) {
	context := NewParseContext(0, 0)

	t.Run("error", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		value, err := interpreter.ResolveUnaryOperations(context, NewErrorValue(errors.New("test error")), "")
		assert.Equal(t, NewErrorValue(err), value)
		assert.EqualError(t, err, "test error")
	})

	t.Run("unknown type", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		value, err := interpreter.ResolveUnaryOperations(context, NewValue("unknown", "0"), "")
		assert.Equal(t, NewErrorValue(err), value)
		assert.EqualError(t, err, UnknownTypeErr{TypeName: "unknown"}.Error())
	})

	t.Run("success", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		expectedTypes := getBasicTypes()

		for _, typeData := range expectedTypes {
			value, err := interpreter.ResolveUnaryOperations(context, typeData.zeroValue, "unknown")
			if typeData.IsBool() || typeData.IsString() {
				assert.Equal(t, NewErrorValue(err), value)
				assert.EqualError(t, err, InvalidOperationErr{TypeNames: []string{typeData.zeroValue.typeName}}.Error())
			} else {
				assert.Equal(t, NewErrorValue(err), value)
				assert.EqualError(t, err, UnknownOperatorErr{Operator: "unknown"}.Error())
			}

			if !typeData.IsBool() && !typeData.IsString() {
				value, err = interpreter.ResolveUnaryOperations(context, typeData.zeroValue, "-")
				assert.NoError(t, err)

				if typeData.IsFloatingPoint() {
					assert.Equal(t, NewValue(typeData.zeroValue.typeName, "-0"), value)
				} else {
					assert.Equal(t, NewValue(typeData.zeroValue.typeName, "0"), value)
				}
			}
		}

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})
}

func TestInterpreterResolveBinaryOperations(t *testing.T) {
	context := NewParseContext(0, 0)

	t.Run("mismatched types", func(t *testing.T) {
		t.Run("error", func(t *testing.T) {
			var buf bytes.Buffer
			interpreter := NewSimInterpreter(&buf)

			value, err := interpreter.ResolveBinaryOperations(context, context, NewValue("int", "0"), NewValue("bool", "false"), "==")
			assert.Equal(t, NewErrorValue(err), value)
			assert.EqualError(t, err, InvalidOperationErr{TypeNames: []string{"int", "bool"}}.Error())
		})

		t.Run("untyped int", func(t *testing.T) {
			var buf bytes.Buffer
			interpreter := NewSimInterpreter(&buf)

			value, err := interpreter.ResolveBinaryOperations(context, context, NewValue("untyped int", "1"), NewValue("int", "2"), "+")
			assert.Equal(t, NewValue("int", "3"), value)
			assert.NoError(t, err)

			value, err = interpreter.ResolveBinaryOperations(context, context, NewValue("int", "1"), NewValue("untyped int", "2"), "+")
			assert.Equal(t, NewValue("int", "3"), value)
			assert.NoError(t, err)

			value, err = interpreter.ResolveBinaryOperations(context, context, NewValue("untyped int", "1"), NewValue("uint", "2"), "+")
			assert.Equal(t, NewValue("uint", "3"), value)
			assert.NoError(t, err)

			value, err = interpreter.ResolveBinaryOperations(context, context, NewValue("uint", "1"), NewValue("untyped int", "2"), "+")
			assert.Equal(t, NewValue("uint", "3"), value)
			assert.NoError(t, err)

			value, err = interpreter.ResolveBinaryOperations(context, context, NewValue("untyped int", "1"), NewValue("untyped float", "2"), "+")
			assert.Equal(t, NewValue("float", "3"), value)
			assert.NoError(t, err)

			value, err = interpreter.ResolveBinaryOperations(context, context, NewValue("untyped float", "1"), NewValue("untyped int", "2"), "+")
			assert.Equal(t, NewValue("float", "3"), value)
			assert.NoError(t, err)
		})

		t.Run("untyped float", func(t *testing.T) {
			var buf bytes.Buffer
			interpreter := NewSimInterpreter(&buf)

			value, err := interpreter.ResolveBinaryOperations(context, context, NewValue("untyped float", "1"), NewValue("float", "2"), "+")
			assert.Equal(t, NewValue("float", "3"), value)
			assert.NoError(t, err)

			value, err = interpreter.ResolveBinaryOperations(context, context, NewValue("float", "1"), NewValue("untyped float", "2"), "+")
			assert.Equal(t, NewValue("float", "3"), value)
			assert.NoError(t, err)
		})
	})

	t.Run("untyped int", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		value, err := interpreter.ResolveBinaryOperations(context, context, NewValue("untyped int", "1"), NewValue("untyped int", "2"), "+")
		assert.Equal(t, NewValue("untyped int", "3"), value)
		assert.NoError(t, err)
	})

	t.Run("untyped float", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		value, err := interpreter.ResolveBinaryOperations(context, context, NewValue("untyped float", "1"), NewValue("untyped float", "2"), "+")
		assert.Equal(t, NewValue("untyped float", "3"), value)
		assert.NoError(t, err)
	})

	t.Run("unknown types", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		value, err := interpreter.ResolveBinaryOperations(context, context, NewValue("unknown", "0"), NewValue("unknown", "false"), "==")
		assert.Equal(t, NewErrorValue(err), value)
		assert.EqualError(t, err, UnknownTypeErr{TypeName: "unknown"}.Error())
	})

	t.Run("success", func(t *testing.T) {
		var buf bytes.Buffer
		interpreter := NewSimInterpreter(&buf)

		expectedTypes := getBasicTypes()

		division := "/"
		modulo := "%"
		arithmeticOperatorsNoDivision := []string{"+", "-", "*"}
		inequalityOperatorsNoEquals := []string{">", "<"}
		inequalityOperatorsEquals := []string{">=", "<="}
		equals := "=="
		notEquals := "!="

		for _, typeData := range expectedTypes {
			value, err := interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, typeData.zeroValue, "unknown")
			assert.Equal(t, NewErrorValue(err), value)
			assert.EqualError(t, err, UnknownOperatorErr{Operator: "unknown"}.Error())

			value, err = interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, typeData.zeroValue, equals)
			assert.NoError(t, err)
			assert.Equal(t, NewValue("bool", "true"), value)

			value, err = interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, typeData.zeroValue, notEquals)
			assert.NoError(t, err)
			assert.Equal(t, NewValue("bool", "false"), value)

			if !typeData.IsBool() {
				for _, operator := range inequalityOperatorsNoEquals {
					value, err := interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, typeData.zeroValue, operator)
					assert.NoError(t, err)
					assert.Equal(t, NewValue("bool", "false"), value)
				}

				for _, operator := range inequalityOperatorsEquals {
					value, err := interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, typeData.zeroValue, operator)
					assert.NoError(t, err)
					assert.Equal(t, NewValue("bool", "true"), value)
				}

				if !typeData.IsString() {
					value, err := interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, typeData.zeroValue, division)
					assert.Equal(t, NewErrorValue(err), value)
					assert.EqualError(t, err, DivideByZeroErr{}.Error())

					one := NewValue(typeData.zeroValue.typeName, "1")

					value, err = interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, one, division)
					assert.NoError(t, err)
					assert.Equal(t, NewValue(typeData.zeroValue.typeName, "0"), value)

					for _, operator := range arithmeticOperatorsNoDivision {
						value, err := interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, typeData.zeroValue, operator)
						assert.NoError(t, err)
						assert.Equal(t, NewValue(typeData.zeroValue.typeName, "0"), value)
					}

					if !typeData.IsFloatingPoint() {
						value, err := interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, typeData.zeroValue, modulo)
						assert.Equal(t, NewErrorValue(err), value)
						assert.EqualError(t, err, DivideByZeroErr{}.Error())

						one := NewValue(typeData.zeroValue.typeName, "1")

						value, err = interpreter.ResolveBinaryOperations(context, context, typeData.zeroValue, one, modulo)
						assert.NoError(t, err)
						assert.Equal(t, NewValue(typeData.zeroValue.typeName, "0"), value)
					}
				}
			}
		}

		assertInterpreterValues(t, interpreter, buf.String(), map[string]Variable{}, []*scope{{}}, getBasicTypes(), "")
	})
}
