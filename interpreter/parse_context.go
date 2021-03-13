package interpreter

import "fmt"

// ParseContext describes context about the parsing that the interpreter
// can use to make correct decisions and display accurate error information.
type ParseContext struct {
	TypeData TypeData

	line   int
	column int
}

// NewParseContext returns a new error context.
func NewParseContext(line int, column int) ParseContext {
	return ParseContext{line: line, column: column}
}

func (e ParseContext) String() string {
	return fmt.Sprintf("line %d:%d", e.line, e.column)
}
