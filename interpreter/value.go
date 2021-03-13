package interpreter

import "strconv"

// Value represents any Sim value. It is useful for translating Sim types to Go types.
type Value struct {
	typeName string
	data     string
	err      error
}

// NewValue returns a new Value type wrapping the given data.
func NewValue(typeName string, data string) Value {
	return Value{
		typeName: typeName,
		data:     data,
	}
}

// NewErrorValue returns a new Value type wrapping the given error.
func NewErrorValue(err error) Value {
	return Value{
		err: err,
	}
}

// GetTypeFromLiteral returns the type of the given literal string.
// Only some types can be represented as a string; other types will have to be casted.
// If no type could be deduced, then this function returns an empty string.
func GetTypeFromLiteral(context ParseContext, literal string) string {
	if context.TypeData.IsEmpty() {
		_, err := NewValue("int", literal).GetInt(context)
		if err == nil {
			return "untyped int"
		}

		_, err = NewValue("float", literal).GetFloat(context)
		if err == nil {
			return "untyped float"
		}
	}

	// We can't assume a literal's specific integer type,
	// so use the provided context to check if the associated type
	// is the correct integer type, and if so, return that type.
	// Otherwise, return an "untyped int", which the caller can deal with.
	if context.TypeData.IsUnsignedInteger() || context.TypeData.IsSignedInteger() {
		_, err := NewValue("uint", literal).GetUint(context)
		if err == nil {
			if context.TypeData.IsUnsignedInteger() {
				return context.TypeData.zeroValue.typeName
			}
		}

		_, err = NewValue("int", literal).GetInt(context)
		if err == nil {
			if context.TypeData.IsSignedInteger() {
				return context.TypeData.zeroValue.typeName
			}
		}
	}

	// We can't assume a literal's specific floating point type,
	// so use the provided context to check if the associated type
	// is the correct floating point type, and if so, return that type.
	// Otherwise, return an "untyped float", which the caller can deal with.
	if context.TypeData.IsFloatingPoint() {
		_, err := NewValue("float", literal).GetFloat(context)
		if err == nil {
			if context.TypeData.IsFloatingPoint() {
				return context.TypeData.zeroValue.typeName
			}
		}
	}

	_, err := NewValue("bool", literal).GetBool(context)
	if err == nil {
		return "bool"
	}

	_, err = NewValue("string", literal).GetString(context)
	if err == nil {
		return "string"
	}

	return ""
}

// GetType returns the type name of the value,
// or returns an error if it is an error value.
func (v Value) GetType() (string, error) {
	if v.err != nil {
		return "", v.err
	}

	return v.typeName, nil
}

// GetRawData returns the value's raw data, or the error
// if the value is storing an error.
func (v Value) GetRawData() (string, error) {
	if v.err != nil {
		return "", v.err
	}

	return v.data, nil
}

// GetString returns the value as a Go string,
// or returns an error if the data is not a string type.
func (v Value) GetString(context ParseContext) (string, error) {
	if v.err != nil {
		return "", v.err
	}

	if len(v.data) < 2 {
		v.err = DataTypeErr{Context: context, TypeName: v.typeName}
		return "", v.err
	}

	if v.data[0] != '"' || v.data[len(v.data)-1] != '"' {
		v.err = DataTypeErr{Context: context, TypeName: v.typeName}
		return "", v.err
	}

	return v.data, nil
}

// GetInt returns the value as a Go int32,
// or returns an error if the data is not an integer type.
func (v Value) GetInt(context ParseContext) (int32, error) {
	if v.err != nil {
		return 0, v.err
	}

	num, err := strconv.ParseInt(v.data, 10, 32)
	if err != nil {
		v.err = DataTypeErr{Context: context, TypeName: v.typeName}
		return 0, v.err
	}

	return int32(num), nil
}

// GetUint returns the value as a Go uint32,
// or returns an error if the data is not an integer type.
func (v Value) GetUint(context ParseContext) (uint32, error) {
	if v.err != nil {
		return 0, v.err
	}

	num, err := strconv.ParseUint(v.data, 10, 32)
	if err != nil {
		v.err = DataTypeErr{Context: context, TypeName: v.typeName}
		return 0, v.err
	}

	return uint32(num), nil
}

// GetByte returns the value as a Go byte,
// or returns an error if the data is not a byte type.
func (v Value) GetByte(context ParseContext) (byte, error) {
	if v.err != nil {
		return 0, v.err
	}

	num, err := strconv.ParseInt(v.data, 10, 8)
	if err != nil {
		v.err = DataTypeErr{Context: context, TypeName: v.typeName}
		return 0, v.err
	}

	return byte(num), nil
}

// GetFloat returns the value as a Go float32,
// or returns an error if the data is not a floating point type.
func (v Value) GetFloat(context ParseContext) (float32, error) {
	if v.err != nil {
		return 0, v.err
	}

	num, err := strconv.ParseFloat(v.data, 32)
	if err != nil {
		v.err = DataTypeErr{Context: context, TypeName: v.typeName}
		return 0, v.err
	}

	return float32(num), nil
}

// GetBool returns the value as a Go bool,
// or returns an error if the data is not a boolean type.
func (v Value) GetBool(context ParseContext) (bool, error) {
	if v.err != nil {
		return false, v.err
	}

	boolean, err := strconv.ParseBool(v.data)
	if err != nil {
		v.err = DataTypeErr{Context: context, TypeName: v.typeName}
		return false, v.err
	}

	return boolean, nil
}
