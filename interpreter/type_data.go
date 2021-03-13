package interpreter

// TypeInfo represents metadata about a type.
type TypeInfo int

const (
	// TypeInfoNone says that a type has no associated metadata.
	TypeInfoNone TypeInfo = 0

	// TypeInfoSignedInteger says that a type is a signed integer.
	TypeInfoSignedInteger TypeInfo = 1

	// TypeInfoUnsignedInteger says that a type is an unsigned integer.
	TypeInfoUnsignedInteger TypeInfo = 2

	// TypeInfoFloatingPoint says that a type is a floating point number.
	TypeInfoFloatingPoint TypeInfo = 3

	// TypeInfoBool says that a type is a boolean.
	TypeInfoBool TypeInfo = 4

	// TypeInfoString says that a type is a string.
	TypeInfoString TypeInfo = 5
)

// TypeData stores common type data, such as the types zero value and casting information.
type TypeData struct {
	zeroValue       Value
	typeInfo        TypeInfo
	implicitCastMap map[string]struct{}
}

// NewTypeData returns a new instance of type data.
func NewTypeData(typeName string, zeroValue string, typeInfo TypeInfo) TypeData {
	return TypeData{
		zeroValue:       NewValue(typeName, zeroValue),
		typeInfo:        typeInfo,
		implicitCastMap: make(map[string]struct{}),
	}
}

// GetTypeName returns the name of the type.
func (t TypeData) GetTypeName() string {
	return t.zeroValue.typeName
}

// IsEmpty checks if the TypeData is empty, representing no type data.
func (t TypeData) IsEmpty() bool {
	return t.zeroValue == NewValue("", "") && t.typeInfo == TypeInfoNone
}

// AddImplicitCast adds an implicit cast from this type to the given type.
// NOTE: This is not a bidirectional cast!
func (t TypeData) AddImplicitCast(context ParseContext, typeData TypeData) error {
	if t.zeroValue.typeName == typeData.zeroValue.typeName {
		return SelfImplicitCastErr{Context: context, TypeName: t.zeroValue.typeName}
	}

	if _, ok := t.implicitCastMap[typeData.zeroValue.typeName]; !ok {
		t.implicitCastMap[typeData.zeroValue.typeName] = struct{}{}
		return nil
	}

	return ImplicitCastExistsErr{Context: context, OriginalTypeName: t.zeroValue.typeName, CastedTypeName: typeData.zeroValue.typeName}
}

// CanImplicitlyCast returns true if this type can be implicitly casted to the given type.
func (t TypeData) CanImplicitlyCast(typeData TypeData) bool {
	if _, ok := t.implicitCastMap[typeData.zeroValue.typeName]; ok {
		return true
	}

	return false
}

// IsSignedInteger returns true if the type is a signed integer type.
func (t TypeData) IsSignedInteger() bool {
	return t.typeInfo == TypeInfoSignedInteger
}

// IsUnsignedInteger returns true if the type is an unsigned integer type.
func (t TypeData) IsUnsignedInteger() bool {
	return t.typeInfo == TypeInfoUnsignedInteger
}

// IsFloatingPoint returns true if the type is a floating point type.
func (t TypeData) IsFloatingPoint() bool {
	return t.typeInfo == TypeInfoFloatingPoint
}

// IsBool returns true if the type is a boolean type.
func (t TypeData) IsBool() bool {
	return t.typeInfo == TypeInfoBool
}

// IsString returns true if the type is a string.
func (t TypeData) IsString() bool {
	return t.typeInfo == TypeInfoString
}
