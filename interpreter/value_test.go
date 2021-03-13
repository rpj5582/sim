package interpreter

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func runTestErrorValueGetError(t *testing.T, expectedErr error, testFunc reflect.Value) {
	value := NewErrorValue(expectedErr)

	if testFunc.Kind() == reflect.Func {
		output := testFunc.Call([]reflect.Value{reflect.ValueOf(value), reflect.ValueOf(NewParseContext(0, 0))})
		val := output[0].Interface()
		err := output[1].Interface().(error)

		assert.Zero(t, val)
		assert.EqualError(t, err, expectedErr.Error())
	} else {
		t.Fatal("testFunc is not a function type")
	}
}

func runTestValueGetError(t *testing.T, expectedTypeName string, expectedData string, expectedErr error, testFunc reflect.Value) {
	value := NewValue(expectedTypeName, expectedData)

	if testFunc.Kind() == reflect.Func {
		output := testFunc.Call([]reflect.Value{reflect.ValueOf(value), reflect.ValueOf(NewParseContext(0, 0))})
		val := output[0].Interface()
		err := output[1].Interface().(error)

		assert.Zero(t, val)
		assert.EqualError(t, err, expectedErr.Error())
	} else {
		t.Fatal("testFunc is not a function type")
	}
}

func runTestValueGetSuccess(t *testing.T, expectedTypeName string, expectedData string, testFunc reflect.Value) {
	value := NewValue(expectedTypeName, expectedData)

	if testFunc.Kind() == reflect.Func {
		output := testFunc.Call([]reflect.Value{reflect.ValueOf(value), reflect.ValueOf(NewParseContext(0, 0))})
		val := output[0].Interface()

		var err error
		if output[1].Interface() != nil {
			err = output[1].Interface().(error)
		}

		assert.NoError(t, err)
		if expectedTypeName == "float" {
			assert.Equal(t, expectedData, fmt.Sprintf("%.1f", val))
		} else {
			assert.Equal(t, expectedData, fmt.Sprintf("%v", val))
		}
	} else {
		t.Fatal("testFunc is not a function type")
	}
}

func TestValueNewValue(t *testing.T) {
	expectedTypeName := "int"
	expectedData := "10"

	value := NewValue(expectedTypeName, expectedData)

	assert.Equal(t, expectedTypeName, value.typeName)
	assert.Equal(t, expectedData, value.data)
	assert.NoError(t, value.err)
}

func TestValueNewErrorValue(t *testing.T) {
	expectedError := errors.New("test error")

	value := NewErrorValue(expectedError)

	assert.Empty(t, value.typeName)
	assert.Empty(t, value.data)
	assert.Equal(t, expectedError, value.err)
}

func TestValueGetTypeFromLiteral(t *testing.T) {
	t.Run("unknown", func(t *testing.T) {
		context := NewParseContext(0, 0)
		typeName := GetTypeFromLiteral(context, "unknown")
		assert.Equal(t, "", typeName)
	})

	types := getBasicTypes()
	for typeName, typeData := range types {
		if typeData.IsSignedInteger() || typeData.IsUnsignedInteger() {
			t.Run("untyped int", func(t *testing.T) {
				context := NewParseContext(0, 0)

				result := GetTypeFromLiteral(context, "0")
				assert.Equal(t, "untyped int", result)
			})

			t.Run("exact match", func(t *testing.T) {
				context := NewParseContext(0, 0)
				context.TypeData = getBasicTypes()[typeName]

				result := GetTypeFromLiteral(context, "0")
				assert.Equal(t, typeName, result)
			})
		}

		if typeData.IsFloatingPoint() {
			t.Run("untyped float", func(t *testing.T) {
				context := NewParseContext(0, 0)

				result := GetTypeFromLiteral(context, "0.0")
				assert.Equal(t, "untyped float", result)
			})

			t.Run("exact match", func(t *testing.T) {
				context := NewParseContext(0, 0)
				context.TypeData = getBasicTypes()[typeName]

				result := GetTypeFromLiteral(context, "0.0")
				assert.Equal(t, typeName, result)
			})
		}
	}

	t.Run("bool", func(t *testing.T) {
		context := NewParseContext(0, 0)
		typeName := GetTypeFromLiteral(context, "true")
		assert.Equal(t, "bool", typeName)
	})

	t.Run("string", func(t *testing.T) {
		context := NewParseContext(0, 0)
		typeName := GetTypeFromLiteral(context, "\"test\"")
		assert.Equal(t, "string", typeName)
	})
}

func TestValueGetTypeName(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		value := NewErrorValue(errors.New("test error"))
		typeName, err := value.GetType()
		assert.Empty(t, typeName)
		assert.EqualError(t, err, "test error")

		assert.Equal(t, "", value.typeName)
		assert.Equal(t, "", value.data)
		assert.EqualError(t, value.err, "test error")
	})

	t.Run("success", func(t *testing.T) {
		expectedTypeName := "int"
		expectedData := "10"

		value := NewValue(expectedTypeName, expectedData)
		typeName, err := value.GetType()
		assert.NoError(t, err)
		assert.Equal(t, expectedTypeName, typeName)

		assert.Equal(t, expectedTypeName, value.typeName)
		assert.Equal(t, expectedData, value.data)
		assert.NoError(t, value.err)
	})
}

func TestValueGet(t *testing.T) {
	type testType int
	var TestTypeSuccess testType = 0
	var TestTypeError testType = 1
	var TestTypeValueError testType = 2

	type testParams struct {
		testType  testType
		typeName  string
		data      string
		err       error
		funcValue reflect.Value
	}

	invalidData := "invalid data"

	params := []testParams{
		// String with error data
		{testType: TestTypeValueError, typeName: "string", err: InvalidValueErr{Data: invalidData}, funcValue: reflect.ValueOf(Value.GetString)},
		// String without both quotes
		{testType: TestTypeError, typeName: "string", data: "", err: DataTypeErr{TypeName: "string"}, funcValue: reflect.ValueOf(Value.GetString)},
		// String with one quote
		{testType: TestTypeError, typeName: "string", data: "\"test string", err: DataTypeErr{TypeName: "string"}, funcValue: reflect.ValueOf(Value.GetString)},
		// String success
		{testType: TestTypeSuccess, typeName: "string", data: "\"test string\"", err: nil, funcValue: reflect.ValueOf(Value.GetString)},

		// Int with error data
		{testType: TestTypeValueError, typeName: "int", err: InvalidValueErr{Data: invalidData}, funcValue: reflect.ValueOf(Value.GetInt)},
		// Int with mismatched type
		{testType: TestTypeError, typeName: "int", data: "10.0", err: DataTypeErr{TypeName: "int"}, funcValue: reflect.ValueOf(Value.GetInt)},
		// Int success
		{testType: TestTypeSuccess, typeName: "int", data: "10", err: nil, funcValue: reflect.ValueOf(Value.GetInt)},

		// Byte with error data
		{testType: TestTypeValueError, typeName: "byte", err: InvalidValueErr{Data: invalidData}, funcValue: reflect.ValueOf(Value.GetByte)},
		// Byte with mismatched type
		{testType: TestTypeError, typeName: "byte", data: "256", err: DataTypeErr{TypeName: "byte"}, funcValue: reflect.ValueOf(Value.GetByte)},
		// Byte success
		{testType: TestTypeSuccess, typeName: "byte", data: "10", err: nil, funcValue: reflect.ValueOf(Value.GetByte)},

		// Uint with error data
		{testType: TestTypeValueError, typeName: "uint", err: InvalidValueErr{Data: invalidData}, funcValue: reflect.ValueOf(Value.GetUint)},
		// Uint with mismatched type
		{testType: TestTypeError, typeName: "uint", data: "-10", err: DataTypeErr{TypeName: "uint"}, funcValue: reflect.ValueOf(Value.GetUint)},
		// Uint success
		{testType: TestTypeSuccess, typeName: "uint", data: "10", err: nil, funcValue: reflect.ValueOf(Value.GetUint)},

		// Float with error data
		{testType: TestTypeValueError, typeName: "float", err: InvalidValueErr{Data: invalidData}, funcValue: reflect.ValueOf(Value.GetFloat)},
		// Float with mismatched type
		{testType: TestTypeError, typeName: "float", data: "false", err: DataTypeErr{TypeName: "float"}, funcValue: reflect.ValueOf(Value.GetFloat)},
		// Float success
		{testType: TestTypeSuccess, typeName: "float", data: "10.0", err: nil, funcValue: reflect.ValueOf(Value.GetFloat)},

		// Bool with error data
		{testType: TestTypeValueError, typeName: "bool", err: InvalidValueErr{Data: invalidData}, funcValue: reflect.ValueOf(Value.GetBool)},
		// Bool with mismatched type
		{testType: TestTypeError, typeName: "bool", data: "10", err: DataTypeErr{TypeName: "bool"}, funcValue: reflect.ValueOf(Value.GetBool)},
		// Bool success
		{testType: TestTypeSuccess, typeName: "bool", data: "true", err: nil, funcValue: reflect.ValueOf(Value.GetBool)},
	}

	for i := range params {
		switch params[i].testType {
		case TestTypeSuccess:
			runTestValueGetSuccess(t, params[i].typeName, params[i].data, params[i].funcValue)

		case TestTypeError:
			runTestValueGetError(t, params[i].typeName, params[i].data, params[i].err, params[i].funcValue)

		case TestTypeValueError:
			runTestErrorValueGetError(t, params[i].err, params[i].funcValue)

		default:
			t.Fatal("unknown test type")
		}
	}
}
