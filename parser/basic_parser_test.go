package parser

import (
	"testing"
	"os"
	"fmt"
	"stone/lexer"
	"stone/token"
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
