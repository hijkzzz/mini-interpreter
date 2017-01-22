package parser

import (
	"testing"
	"os"
	"fmt"
	"stone/lexer"
	"stone/token"
	"stone/environment"
)

func Test_Parser(t *testing.T) {
	fin, err := os.Open("parser_test")
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(fin)
	p := NewParser(l)
	for l.Peek(0) != token.EOF {
		fmt.Println(p.program().String())
	}
}

func Test_Eval(t *testing.T) {
	fin, err := os.Open("eval_test")
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(fin)
	p := NewParser(l)
	e := environment.NewNestedEnv()
	for l.Peek(0) != token.EOF {
		program := p.program()
		result := program.Eval(e)
		fmt.Println(result)
	}
}
