package parser

import (
	"testing"
	"os"
	"fmt"
	"stone/lexer"
	"stone/token"
	"stone/environment"
)

func Test_BasicParser_Program(t *testing.T) {
	fin, err := os.Open("basic_parser_test")
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(fin)
	p := NewBasicParser(l)
	for l.Peek(0) != token.EOF {
		fmt.Println(p.Program().String())
	}
}

func Test_ASTree_Eval(t *testing.T) {
	fin, err := os.Open("astree_eval_test")
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(fin)
	p := NewBasicParser(l)
	e := environment.NewBasicEnv()
	for l.Peek(0) != token.EOF {
		program := p.Program()
		result := program.Eval(e)
		fmt.Println(result)
	}
}
