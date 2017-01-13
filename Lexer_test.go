package stone

import (
	"testing"
	"os"
	"stone/token"
	"fmt"
)

func Test_Lexer_Read(t *testing.T) {
	fin, err := os.Open("Lexer_test")
	if err != nil {
		panic(err)
	}

	l := NewLexer(fin)
	for t := l.Read(); t != token.EOF; t = l.Read() {
		fmt.Printf("%d %s\n", t.GetLineNumber(), t.GetText())
	}
}
