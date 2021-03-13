package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rpj5582/sim/interpreter"
	"github.com/rpj5582/sim/listener"
	"github.com/rpj5582/sim/parser"
	"github.com/rpj5582/sim/visitor"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("specify an input file to interpret")
		return
	}

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	lexer := parser.NewSimLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSimParser(stream)
	errListener := listener.NewSimErrorListener(antlr.NewDiagnosticErrorListener(false))
	p.BuildParseTrees = true

	var buf bytes.Buffer
	visitor := visitor.NewSimVisitor(interpreter.NewSimInterpreter(&buf))

	tree := p.Start()

	if err := errListener.GetError(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err, ok := visitor.Visit(tree).(error); ok && err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(buf.String())
}
