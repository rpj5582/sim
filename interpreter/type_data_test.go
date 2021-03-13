package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeDataNewTypeData(t *testing.T) {
	typeData := NewTypeData("int", "0", TypeInfoSignedInteger)
	assert.Equal(t, TypeData{zeroValue: NewValue("int", "0"), typeInfo: TypeInfoSignedInteger, implicitCastMap: make(map[string]struct{})}, typeData)
}

func TestTypeDataGetTypeName(t *testing.T) {
	typeData := NewTypeData("int", "0", TypeInfoSignedInteger)
	typeName := typeData.GetTypeName()
	assert.Equal(t, "int", typeName)
}

func TestTypeDataIsEmpy(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		typeData := NewTypeData("", "", TypeInfoNone)
		isEmpty := typeData.IsEmpty()
		assert.True(t, isEmpty)
	})

	t.Run("not empty", func(t *testing.T) {
		typeData := NewTypeData("int", "0", TypeInfoSignedInteger)
		isEmpty := typeData.IsEmpty()
		assert.False(t, isEmpty)
	})
}

func TestTypeDataAddImplicitCast(t *testing.T) {
	context := NewParseContext(0, 0)

	typeDataInt := NewTypeData("int", "0", TypeInfoSignedInteger)
	typeDataInt64 := NewTypeData("int64", "0", TypeInfoSignedInteger)

	t.Run("self implicit cast", func(t *testing.T) {
		err := typeDataInt.AddImplicitCast(context, typeDataInt)
		assert.EqualError(t, err, SelfImplicitCastErr{TypeName: typeDataInt.zeroValue.typeName}.Error())
	})

	t.Run("success", func(t *testing.T) {
		err := typeDataInt.AddImplicitCast(context, typeDataInt64)
		assert.NoError(t, err)
	})

	t.Run("implicit cast exists", func(t *testing.T) {
		err := typeDataInt.AddImplicitCast(context, typeDataInt64)
		assert.EqualError(t, err, ImplicitCastExistsErr{OriginalTypeName: typeDataInt.zeroValue.typeName, CastedTypeName: typeDataInt64.zeroValue.typeName}.Error())
	})
}

func TestTypeDataCanImplicitlyCast(t *testing.T) {
	typeDataInt := NewTypeData("int", "0", TypeInfoSignedInteger)
	typeDataInt64 := NewTypeData("int64", "0", TypeInfoSignedInteger)

	err := typeDataInt.AddImplicitCast(NewParseContext(0, 0), typeDataInt64)
	assert.NoError(t, err)

	assert.False(t, typeDataInt64.CanImplicitlyCast(typeDataInt))
	assert.True(t, typeDataInt.CanImplicitlyCast(typeDataInt64))
}

func TestTypeDataIsSignedInteger(t *testing.T) {
	typeDataInt := NewTypeData("int", "0", TypeInfoSignedInteger)
	typeDataUInt := NewTypeData("uint", "0", TypeInfoUnsignedInteger)

	assert.True(t, typeDataInt.IsSignedInteger())
	assert.False(t, typeDataUInt.IsSignedInteger())
}

func TestTypeDataIsUnsignedInteger(t *testing.T) {
	typeDataInt := NewTypeData("int", "0", TypeInfoSignedInteger)
	typeDataUInt := NewTypeData("uint", "0", TypeInfoUnsignedInteger)

	assert.False(t, typeDataInt.IsUnsignedInteger())
	assert.True(t, typeDataUInt.IsUnsignedInteger())
}

func TestTypeDataIsFloatingPoint(t *testing.T) {
	typeDataInt := NewTypeData("int", "0", TypeInfoSignedInteger)
	typeDataFloat := NewTypeData("float", "0", TypeInfoFloatingPoint)

	assert.False(t, typeDataInt.IsFloatingPoint())
	assert.True(t, typeDataFloat.IsFloatingPoint())
}

func TestTypeDataIsBool(t *testing.T) {
	typeDataInt := NewTypeData("int", "0", TypeInfoSignedInteger)
	typeDataFloat := NewTypeData("bool", "0", TypeInfoBool)

	assert.False(t, typeDataInt.IsBool())
	assert.True(t, typeDataFloat.IsBool())
}

func TestTypeDataIsString(t *testing.T) {
	typeDataInt := NewTypeData("int", "0", TypeInfoSignedInteger)
	typeDataFloat := NewTypeData("string", "0", TypeInfoString)

	assert.False(t, typeDataInt.IsString())
	assert.True(t, typeDataFloat.IsString())
}
