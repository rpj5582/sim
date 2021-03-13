package interpreter

// Variable is a representation of data with a fixed data type and a value that can be changed.
type Variable struct {
	name  string
	value Value
}

// NewVariable returns a new instance of a variable.
func NewVariable(name string, value Value) Variable {
	return Variable{
		name:  name,
		value: value,
	}
}

// Value returns the value of the variable.
func (v Variable) Value() Value {
	return v.value
}
