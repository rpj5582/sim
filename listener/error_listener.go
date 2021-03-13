package listener

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SimErrorListener struct {
	*antlr.DefaultErrorListener

	err error
}

func NewSimErrorListener(diag *antlr.DiagnosticErrorListener) *SimErrorListener {
	return new(SimErrorListener)
}

func (l *SimErrorListener) GetError() error {
	return l.err
}

func (l *SimErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.err = fmt.Errorf("syntax error: line %d:%d %s", line, column, msg)
}
