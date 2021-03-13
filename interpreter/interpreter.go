package interpreter

import (
	"fmt"
	"io"
)

type scope struct {
	varNames []string
}

// SimInterpreter interprents Sim by simulating a runtime environment,
// keeping track of declared types, variables, scopes, etc.
type SimInterpreter struct {
	types  map[string]TypeData
	vars   map[string]Variable
	scopes []*scope

	output io.ReadWriter
}

func getBasicTypes() map[string]TypeData {
	types := make(map[string]TypeData)

	types["int"] = TypeData{
		zeroValue: NewValue("int", "0"),
		typeInfo:  TypeInfoSignedInteger,
		implicitCastMap: map[string]struct{}{
			"int32": {},
			"int64": {},
		},
	}

	types["int8"] = TypeData{
		zeroValue: NewValue("int8", "0"),
		typeInfo:  TypeInfoSignedInteger,
		implicitCastMap: map[string]struct{}{
			"int":   {},
			"int16": {},
			"int32": {},
			"int64": {},
		},
	}

	types["int16"] = TypeData{
		zeroValue: NewValue("int16", "0"),
		typeInfo:  TypeInfoSignedInteger,
		implicitCastMap: map[string]struct{}{
			"int":   {},
			"int32": {},
			"int64": {},
		},
	}

	types["int32"] = TypeData{
		zeroValue: NewValue("int32", "0"),
		typeInfo:  TypeInfoSignedInteger,
		implicitCastMap: map[string]struct{}{
			"int":   {},
			"int64": {},
		},
	}

	types["int64"] = TypeData{
		zeroValue:       NewValue("int64", "0"),
		typeInfo:        TypeInfoSignedInteger,
		implicitCastMap: map[string]struct{}{},
	}

	types["uint"] = TypeData{
		zeroValue: NewValue("uint", "0"),
		typeInfo:  TypeInfoUnsignedInteger,
		implicitCastMap: map[string]struct{}{
			"uint32": {},
			"uint64": {},
		},
	}

	types["byte"] = TypeData{
		zeroValue: NewValue("byte", "0"),
		typeInfo:  TypeInfoUnsignedInteger,
		implicitCastMap: map[string]struct{}{
			"uint":   {},
			"uint8":  {},
			"uint16": {},
			"uint32": {},
			"uint64": {},
		},
	}

	types["uint8"] = TypeData{
		zeroValue: NewValue("uint8", "0"),
		typeInfo:  TypeInfoUnsignedInteger,
		implicitCastMap: map[string]struct{}{
			"byte":   {},
			"uint":   {},
			"uint16": {},
			"uint32": {},
			"uint64": {},
		},
	}

	types["uint16"] = TypeData{
		zeroValue: NewValue("uint16", "0"),
		typeInfo:  TypeInfoUnsignedInteger,
		implicitCastMap: map[string]struct{}{
			"uint":   {},
			"uint32": {},
			"uint64": {},
		},
	}

	types["uint32"] = TypeData{
		zeroValue: NewValue("uint32", "0"),
		typeInfo:  TypeInfoUnsignedInteger,
		implicitCastMap: map[string]struct{}{
			"uint":   {},
			"uint64": {},
		},
	}

	types["uint64"] = TypeData{
		zeroValue:       NewValue("uint64", "0"),
		typeInfo:        TypeInfoUnsignedInteger,
		implicitCastMap: map[string]struct{}{},
	}

	types["float"] = TypeData{
		zeroValue: NewValue("float", "0.0"),
		typeInfo:  TypeInfoFloatingPoint,
		implicitCastMap: map[string]struct{}{
			"float32": {},
			"float64": {},
		},
	}

	types["float32"] = TypeData{
		zeroValue: NewValue("float32", "0.0"),
		typeInfo:  TypeInfoFloatingPoint,
		implicitCastMap: map[string]struct{}{
			"float":   {},
			"float64": {},
		},
	}

	types["float64"] = TypeData{
		zeroValue:       NewValue("float64", "0.0"),
		typeInfo:        TypeInfoFloatingPoint,
		implicitCastMap: map[string]struct{}{},
	}

	types["bool"] = TypeData{
		zeroValue:       NewValue("bool", "false"),
		typeInfo:        TypeInfoBool,
		implicitCastMap: map[string]struct{}{},
	}

	types["string"] = TypeData{
		zeroValue:       NewValue("string", "\"\""),
		typeInfo:        TypeInfoString,
		implicitCastMap: map[string]struct{}{},
	}

	return types
}

// NewSimInterpreter creates a new SimInterpreter instance.
func NewSimInterpreter(output io.ReadWriter) *SimInterpreter {

	return &SimInterpreter{
		types:  getBasicTypes(),
		vars:   make(map[string]Variable),
		scopes: []*scope{{}}, // Always have a global scope
		output: output,
	}
}

// GetTypeData returns the type data for the provided type name,
// or an error if that type name doesn't exist.
func (interpreter *SimInterpreter) GetTypeData(context ParseContext, typeName string) (TypeData, error) {
	if typeData, ok := interpreter.types[typeName]; ok {
		return typeData, nil
	}

	return TypeData{}, UnknownTypeErr{Context: context, TypeName: typeName}
}

// PushScope adds a new scope to the scope stack.
// Scopes keep track of when to remove variables.
func (interpreter *SimInterpreter) PushScope() {
	interpreter.scopes = append(interpreter.scopes, &scope{})
}

// PopScope removes the most recently added scope from the scope stack,
// along with all of the variables declared in that scope.
func (interpreter *SimInterpreter) PopScope(context ParseContext) error {
	if len(interpreter.scopes) <= 1 {
		return ExitGlobalScopeErr{Context: context}
	}

	currScope := interpreter.currentScope()

	// Remove all variables that were declared in the current scope from the variable map
	for i := 0; i < len(currScope.varNames); i++ {
		delete(interpreter.vars, currScope.varNames[i])
	}

	interpreter.scopes = interpreter.scopes[:len(interpreter.scopes)-1]

	return nil
}

// AddVar adds a new variable to the variable map, owned by the current scope.
func (interpreter *SimInterpreter) AddVar(context ParseContext, variable Variable) error {
	if variable.value.err != nil {
		return variable.value.err
	}

	// If the variable is still an untyped int or float, switch them to the concrete types.
	if variable.value.typeName == "untyped int" {
		variable.value.typeName = "int"
	}

	if variable.value.typeName == "untyped float" {
		variable.value.typeName = "float"
	}

	typeData, ok := interpreter.types[variable.value.typeName]
	if !ok {
		return UnknownTypeErr{Context: context, TypeName: variable.value.typeName}
	}

	oldTypeData := context.TypeData
	defer func() {
		context.TypeData = oldTypeData
	}()

	context.TypeData = typeData

	if _, ok := interpreter.vars[variable.name]; ok {
		return VarExistsErr{Context: context, VarName: variable.name}
	}

	if variable.value.data == "" {
		variable.value.data = interpreter.types[variable.value.typeName].zeroValue.data
	}

	if ok := interpreter.validateValue(context, variable.value); !ok {
		return MismatchedTypeAssignErr{Context: context, Var: variable}
	}

	currScope := interpreter.currentScope()

	currScope.varNames = append(currScope.varNames, variable.name)
	interpreter.vars[variable.name] = variable

	return nil
}

// GetVar returns the variable given its name.
func (interpreter *SimInterpreter) GetVar(context ParseContext, varName string) (Variable, error) {
	variable, ok := interpreter.vars[varName]
	if !ok {
		return Variable{}, UnknownVarErr{Context: context, VarName: varName}
	}

	return variable, nil
}

// SetVarValue sets the value of a variable.
// Setting a variable cannot change its underlying type.
func (interpreter *SimInterpreter) SetVarValue(context ParseContext, varName string, value Value) error {
	variable, ok := interpreter.vars[varName]
	if !ok {
		return UnknownVarErr{Context: context, VarName: varName}
	}

	// If the value is still an untyped int or float, switch them to the concrete types.
	if value.typeName == "untyped int" {
		value.typeName = "int"
	}

	if value.typeName == "untyped float" {
		value.typeName = "float"
	}

	_, ok = interpreter.types[value.typeName]
	if !ok {
		return InvalidTypeErr{Context: context, TypeName: value.typeName, VarName: varName}
	}

	varTypeData, ok := interpreter.types[variable.value.typeName]
	if !ok {
		return UnknownTypeErr{Context: context, TypeName: variable.value.typeName}
	}

	oldTypeData := context.TypeData
	defer func() {
		context.TypeData = oldTypeData
	}()

	context.TypeData = varTypeData

	if value.typeName != variable.value.typeName {
		return MismatchedTypeAssignErr{Context: context, Var: variable}
	}

	if ok := interpreter.validateValue(context, value); !ok {
		return MismatchedTypeAssignErr{Context: context, Var: variable}
	}

	interpreter.vars[varName] = NewVariable(variable.name, value)

	return nil
}

// PrintLine adds the given output to the end of the output stream followed by a newline.
func (interpreter *SimInterpreter) PrintLine(output interface{}) {
	fmt.Fprintln(interpreter.output, output)
}

// GetAllVars returns the map of all variables the interpreter currently knows about keyed by variable name.
func (interpreter *SimInterpreter) GetAllVars() map[string]Variable {
	varsCopy := make(map[string]Variable)

	for k, v := range interpreter.vars {
		varsCopy[k] = v
	}

	return varsCopy
}

// ResolveUnaryOperations resolves a single value using the given operator.
// Some values do not implement all binary operators, such as bool for example.
func (interpreter *SimInterpreter) ResolveUnaryOperations(context ParseContext, val Value, operator string) (Value, error) {
	typeName, err := val.GetType()
	if err != nil {
		return NewErrorValue(err), err
	}

	typeData, err := interpreter.GetTypeData(context, typeName)
	if err != nil {
		return NewErrorValue(err), err
	}

	if typeData.IsSignedInteger() {
		num, err := val.GetInt(context)
		if err != nil {
			return NewErrorValue(err), err
		}

		switch operator {
		case "-":
			return NewValue(typeName, fmt.Sprintf("%d", -num)), nil
		default:
			err := UnknownOperatorErr{Context: context, Operator: operator}
			return NewErrorValue(err), err
		}
	}

	if typeData.IsUnsignedInteger() {
		num, err := val.GetUint(context)
		if err != nil {
			return NewErrorValue(err), err
		}

		switch operator {
		case "-":
			return NewValue(typeName, fmt.Sprintf("%d", -num)), nil
		default:
			err := UnknownOperatorErr{Context: context, Operator: operator}
			return NewErrorValue(err), err
		}
	}

	if typeData.IsFloatingPoint() {
		num, err := val.GetFloat(context)
		if err != nil {
			return NewErrorValue(err), err
		}

		switch operator {
		case "-":
			return NewValue(typeName, fmt.Sprintf("%g", -num)), nil
		default:
			err := UnknownOperatorErr{Context: context, Operator: operator}
			return NewErrorValue(err), err
		}
	}

	err = InvalidOperationErr{Context: context, TypeNames: []string{typeName}}
	return NewErrorValue(err), err
}

// ResolveBinaryOperations resolves two values using the given operator.
// Some values do not implement all binary operators, such as bool for example.
func (interpreter *SimInterpreter) ResolveBinaryOperations(leftContext ParseContext, rightContext ParseContext, leftVal Value, rightVal Value, operator string) (Value, error) {
	leftTypeName, err := leftVal.GetType()
	if err != nil {
		return NewErrorValue(err), err
	}

	rightTypeName, err := rightVal.GetType()
	if err != nil {
		return NewErrorValue(err), err
	}

	if leftTypeName != rightTypeName {
		return interpreter.handleMismatchedTypesBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, rightTypeName, operator)
	}

	if leftTypeName == "untyped int" {
		return interpreter.handleSignedIntegerBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
	}

	if leftTypeName == "untyped float" {
		return interpreter.handleFloatingPointBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
	}

	leftTypeData, err := interpreter.GetTypeData(leftContext, leftTypeName)
	if err != nil {
		return NewErrorValue(err), err
	}

	if leftTypeData.IsSignedInteger() {
		return interpreter.handleSignedIntegerBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
	}

	if leftTypeData.IsUnsignedInteger() {
		return interpreter.handleUnsignedIntegerBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
	}

	if leftTypeData.IsFloatingPoint() {
		return interpreter.handleFloatingPointBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
	}

	if leftTypeData.IsBool() {
		return interpreter.handleBooleanBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
	}

	if leftTypeData.IsString() {
		return interpreter.handleStringBinaryOperations(leftContext, rightContext, leftVal, rightVal, operator)
	}

	err = InvalidOperationErr{Context: leftContext, TypeNames: []string{leftTypeName, rightTypeName}}
	return NewErrorValue(err), err
}

func (interpreter *SimInterpreter) handleSignedIntegerBinaryOperations(leftContext, rightContext ParseContext, leftVal, rightVal Value, typeName string, operator string) (Value, error) {
	left, err := leftVal.GetInt(leftContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	right, err := rightVal.GetInt(rightContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	switch operator {
	case "+":
		return NewValue(typeName, fmt.Sprintf("%d", left+right)), nil
	case "-":
		return NewValue(typeName, fmt.Sprintf("%d", left-right)), nil
	case "*":
		return NewValue(typeName, fmt.Sprintf("%d", left*right)), nil
	case "/":
		if right == 0 {
			err := DivideByZeroErr{Context: rightContext}
			return NewErrorValue(err), err
		}

		return NewValue(typeName, fmt.Sprintf("%d", left/right)), nil
	case "%":
		if right == 0 {
			err := DivideByZeroErr{Context: rightContext}
			return NewErrorValue(err), err
		}

		return NewValue(typeName, fmt.Sprintf("%d", left%right)), nil
	case ">":
		return NewValue("bool", fmt.Sprintf("%t", left > right)), nil
	case "<":
		return NewValue("bool", fmt.Sprintf("%t", left < right)), nil
	case ">=":
		return NewValue("bool", fmt.Sprintf("%t", left >= right)), nil
	case "<=":
		return NewValue("bool", fmt.Sprintf("%t", left <= right)), nil
	case "==":
		return NewValue("bool", fmt.Sprintf("%t", left == right)), nil
	case "!=":
		return NewValue("bool", fmt.Sprintf("%t", left != right)), nil
	default:
		err := UnknownOperatorErr{Context: leftContext, Operator: operator}
		return NewErrorValue(err), err
	}
}

func (interpreter *SimInterpreter) handleUnsignedIntegerBinaryOperations(leftContext, rightContext ParseContext, leftVal, rightVal Value, typeName string, operator string) (Value, error) {
	left, err := leftVal.GetUint(leftContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	right, err := rightVal.GetUint(rightContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	switch operator {
	case "+":
		return NewValue(typeName, fmt.Sprintf("%d", left+right)), nil
	case "-":
		return NewValue(typeName, fmt.Sprintf("%d", left-right)), nil
	case "*":
		return NewValue(typeName, fmt.Sprintf("%d", left*right)), nil
	case "/":
		if right == 0 {
			err := DivideByZeroErr{Context: rightContext}
			return NewErrorValue(err), err
		}

		return NewValue(typeName, fmt.Sprintf("%d", left/right)), nil
	case "%":
		if right == 0 {
			err := DivideByZeroErr{Context: rightContext}
			return NewErrorValue(err), err
		}

		return NewValue(typeName, fmt.Sprintf("%d", left%right)), nil
	case ">":
		return NewValue("bool", fmt.Sprintf("%t", left > right)), nil
	case "<":
		return NewValue("bool", fmt.Sprintf("%t", left < right)), nil
	case ">=":
		return NewValue("bool", fmt.Sprintf("%t", left >= right)), nil
	case "<=":
		return NewValue("bool", fmt.Sprintf("%t", left <= right)), nil
	case "==":
		return NewValue("bool", fmt.Sprintf("%t", left == right)), nil
	case "!=":
		return NewValue("bool", fmt.Sprintf("%t", left != right)), nil
	default:
		err := UnknownOperatorErr{Context: leftContext, Operator: operator}
		return NewErrorValue(err), err
	}
}

func (interpreter *SimInterpreter) handleFloatingPointBinaryOperations(leftContext, rightContext ParseContext, leftVal, rightVal Value, typeName string, operator string) (Value, error) {
	left, err := leftVal.GetFloat(leftContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	right, err := rightVal.GetFloat(rightContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	switch operator {
	case "+":
		return NewValue(typeName, fmt.Sprintf("%g", left+right)), nil
	case "-":
		return NewValue(typeName, fmt.Sprintf("%g", left-right)), nil
	case "*":
		return NewValue(typeName, fmt.Sprintf("%g", left*right)), nil
	case "/":
		if right == 0 {
			err := DivideByZeroErr{Context: rightContext}
			return NewErrorValue(err), err
		}

		return NewValue(typeName, fmt.Sprintf("%g", left/right)), nil
	case ">":
		return NewValue("bool", fmt.Sprintf("%t", left > right)), nil
	case "<":
		return NewValue("bool", fmt.Sprintf("%t", left < right)), nil
	case ">=":
		return NewValue("bool", fmt.Sprintf("%t", left >= right)), nil
	case "<=":
		return NewValue("bool", fmt.Sprintf("%t", left <= right)), nil
	case "==":
		return NewValue("bool", fmt.Sprintf("%t", left == right)), nil
	case "!=":
		return NewValue("bool", fmt.Sprintf("%t", left != right)), nil
	default:
		err := UnknownOperatorErr{Context: leftContext, Operator: operator}
		return NewErrorValue(err), err
	}
}

func (interpreter *SimInterpreter) handleBooleanBinaryOperations(leftContext, rightContext ParseContext, leftVal, rightVal Value, typeName string, operator string) (Value, error) {
	left, err := leftVal.GetBool(leftContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	right, err := rightVal.GetBool(rightContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	switch operator {
	case "==":
		return NewValue(typeName, fmt.Sprintf("%t", left == right)), nil
	case "!=":
		return NewValue(typeName, fmt.Sprintf("%t", left != right)), nil
	default:
		err := UnknownOperatorErr{Context: leftContext, Operator: operator}
		return NewErrorValue(err), err
	}
}

func (interpreter *SimInterpreter) handleStringBinaryOperations(leftContext, rightContext ParseContext, leftVal, rightVal Value, operator string) (Value, error) {
	left, err := leftVal.GetString(leftContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	right, err := rightVal.GetString(rightContext)
	if err != nil {
		return NewErrorValue(err), err
	}

	switch operator {
	case ">":
		return NewValue("bool", fmt.Sprintf("%t", left > right)), nil
	case "<":
		return NewValue("bool", fmt.Sprintf("%t", left < right)), nil
	case ">=":
		return NewValue("bool", fmt.Sprintf("%t", left >= right)), nil
	case "<=":
		return NewValue("bool", fmt.Sprintf("%t", left <= right)), nil
	case "==":
		return NewValue("bool", fmt.Sprintf("%t", left == right)), nil
	case "!=":
		return NewValue("bool", fmt.Sprintf("%t", left != right)), nil
	default:
		err := UnknownOperatorErr{Context: leftContext, Operator: operator}
		return NewErrorValue(err), err
	}
}

func (interpreter *SimInterpreter) handleMismatchedTypesBinaryOperations(leftContext, rightContext ParseContext, leftVal, rightVal Value, leftTypeName string, rightTypeName string, operator string) (Value, error) {
	if leftTypeName == "untyped int" {
		if rightTypeName == "untyped float" {
			rightTypeName = "float"
		}

		rightTypeData, err := interpreter.GetTypeData(rightContext, rightTypeName)
		if err != nil {
			return NewErrorValue(err), err
		}

		if rightTypeData.IsSignedInteger() {
			leftTypeName = rightTypeName
			return interpreter.handleSignedIntegerBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
		}

		if rightTypeData.IsUnsignedInteger() {
			leftTypeName = rightTypeName
			return interpreter.handleUnsignedIntegerBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
		}

		if rightTypeData.IsFloatingPoint() {
			leftTypeName = rightTypeName
			return interpreter.handleFloatingPointBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
		}
	}

	if rightTypeName == "untyped int" {
		if leftTypeName == "untyped float" {
			leftTypeName = "float"
		}

		leftTypeData, err := interpreter.GetTypeData(leftContext, leftTypeName)
		if err != nil {
			return NewErrorValue(err), err
		}

		if leftTypeData.IsSignedInteger() {
			rightTypeName = leftTypeName
			return interpreter.handleSignedIntegerBinaryOperations(leftContext, rightContext, leftVal, rightVal, rightTypeName, operator)
		}

		if leftTypeData.IsUnsignedInteger() {
			rightTypeName = leftTypeName
			return interpreter.handleUnsignedIntegerBinaryOperations(leftContext, rightContext, leftVal, rightVal, rightTypeName, operator)
		}

		if leftTypeData.IsFloatingPoint() {
			rightTypeName = leftTypeName
			return interpreter.handleFloatingPointBinaryOperations(leftContext, rightContext, leftVal, rightVal, rightTypeName, operator)
		}
	}

	if leftTypeName == "untyped float" {
		rightTypeData, err := interpreter.GetTypeData(rightContext, rightTypeName)
		if err != nil {
			return NewErrorValue(err), err
		}

		if rightTypeData.IsFloatingPoint() {
			leftTypeName = rightTypeName
			return interpreter.handleFloatingPointBinaryOperations(leftContext, rightContext, leftVal, rightVal, leftTypeName, operator)
		}
	}

	if rightTypeName == "untyped float" {
		leftTypeData, err := interpreter.GetTypeData(leftContext, leftTypeName)
		if err != nil {
			return NewErrorValue(err), err
		}

		if leftTypeData.IsFloatingPoint() {
			rightTypeName = leftTypeName
			return interpreter.handleFloatingPointBinaryOperations(leftContext, rightContext, leftVal, rightVal, rightTypeName, operator)
		}
	}

	err := InvalidOperationErr{Context: leftContext, TypeNames: []string{leftTypeName, rightTypeName}}
	return NewErrorValue(err), err
}

func (interpreter *SimInterpreter) validateValue(context ParseContext, value Value) bool {
	return GetTypeFromLiteral(context, value.data) == value.typeName
}

// Helper function to return the current scope
func (interpreter *SimInterpreter) currentScope() *scope {
	return interpreter.scopes[len(interpreter.scopes)-1]
}
