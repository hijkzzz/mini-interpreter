package lexer

import (
	"stone/token"

	"testing"
	"os"
	"fmt"
)

func Test_Lexer_Read(t *testing.T) {
	fin, err := os.Open("lexer_test")
	if err != nil {
		panic(err)
	}

	l := NewLexer(fin)
	for t := l.Read(); t != token.EOF; t = l.Read() {
		fmt.Printf("%d %s\n", t.GetLineNumber(), t.GetText())
	}
}
