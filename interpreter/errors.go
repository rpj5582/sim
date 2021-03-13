package interpreter

import "fmt"

// ExitGlobalScopeErr describes an attempt to pop off the global scope.
type ExitGlobalScopeErr struct {
	Context ParseContext
}

func (e ExitGlobalScopeErr) Error() string {
	return fmt.Sprintf("%s: attempted to exit the global scope", e.Context.String())
}

// UnknownTypeErr is returned when a type is referenced that has not yet been declared.
type UnknownTypeErr struct {
	Context  ParseContext
	TypeName string
}

func (e UnknownTypeErr) Error() string {
	return fmt.Sprintf("%s: type %s is not declared in this scope", e.Context.String(), e.TypeName)
}

// InvalidTypeErr is returned when a variable somehow has a type that has not yet been declared.
type InvalidTypeErr struct {
	Context  ParseContext
	TypeName string
	VarName  string
}

func (e InvalidTypeErr) Error() string {
	return fmt.Sprintf("%s: %s has an invalid type %q", e.Context.String(), e.VarName, e.TypeName)
}

// UnknownVarErr is returned when a variable is referenced that has not yet been declared.
type UnknownVarErr struct {
	Context ParseContext
	VarName string
}

func (e UnknownVarErr) Error() string {
	return fmt.Sprintf("%s: var %s is not declared in this scope", e.Context.String(), e.VarName)
}

// VarExistsErr is returned when a variable has already been declared but is trying to be declared again.
type VarExistsErr struct {
	Context ParseContext
	VarName string
}

func (e VarExistsErr) Error() string {
	return fmt.Sprintf("%s: var %s is already declared in this scope", e.Context.String(), e.VarName)
}

// MismatchedTypeAssignErr is returned when a variable is being assigned a value,
// and that value's type is mismatched with the variable's type.
type MismatchedTypeAssignErr struct {
	Context ParseContext
	Var     Variable
}

func (e MismatchedTypeAssignErr) Error() string {
	return fmt.Sprintf("%s: cannot assign %s to %s of type %s", e.Context.String(), e.Var.value.data, e.Var.name, e.Var.value.typeName)
}

// DataTypeErr is returned when a value's data mismatches its expected type.
type DataTypeErr struct {
	Context  ParseContext
	TypeName string
}

func (e DataTypeErr) Error() string {
	return fmt.Sprintf("%s: value is not of type %s", e.Context.String(), e.TypeName)
}

// InvalidValueErr is returned when the result of an expression is not a Value type or another error.
type InvalidValueErr struct {
	Context ParseContext
	Data    interface{}
}

func (e InvalidValueErr) Error() string {
	return fmt.Sprintf("%s: %q is not a proper value type", e.Context.String(), e.Data)
}

// InvalidControlFlowErr is returned when the result of a statement is not a ControlFlow type or another error.
type InvalidControlFlowErr struct {
	Context ParseContext
	Data    interface{}
}

func (e InvalidControlFlowErr) Error() string {
	return fmt.Sprintf("%s: %q is not a proper control flow type", e.Context.String(), e.Data)
}

// ImplicitCastExistsErr is returned when an implicit cast was already defined for the type.
type ImplicitCastExistsErr struct {
	Context          ParseContext
	OriginalTypeName string
	CastedTypeName   string
}

func (e ImplicitCastExistsErr) Error() string {
	return fmt.Sprintf("%s: implicit cast from %s to %s already exists", e.Context.String(), e.OriginalTypeName, e.CastedTypeName)
}

// SelfImplicitCastErr is returned when a type tried to add an implicit cast to itself.
type SelfImplicitCastErr struct {
	Context  ParseContext
	TypeName string
}

func (e SelfImplicitCastErr) Error() string {
	return fmt.Sprintf("%s: cannot add an implicit cast for type %s to itself", e.Context.String(), e.TypeName)
}

// InvalidOperationErr is returned when an operation cannot be completed between two types.
type InvalidOperationErr struct {
	Context   ParseContext
	TypeNames []string
}

func (e InvalidOperationErr) Error() string {
	if len(e.TypeNames) == 0 {
		return "invalid operation"
	}

	if len(e.TypeNames) == 1 {
		return fmt.Sprintf("%s: invalid operation for type %s", e.Context.String(), e.TypeNames[0])
	}

	return fmt.Sprintf("%s: invalid operation between types %s and %s", e.Context.String(), e.TypeNames[0], e.TypeNames[1])
}

// UnknownOperatorErr is returned an unknown operator was provided.
type UnknownOperatorErr struct {
	Context  ParseContext
	Operator string
}

func (e UnknownOperatorErr) Error() string {
	return fmt.Sprintf("%s: unknown operator %s", e.Context.String(), e.Operator)
}

// DivideByZeroErr is returned when an attempt was made to divide by zero.
type DivideByZeroErr struct {
	Context ParseContext
}

func (e DivideByZeroErr) Error() string {
	return fmt.Sprintf("%s: divide by zero", e.Context.String())
}
