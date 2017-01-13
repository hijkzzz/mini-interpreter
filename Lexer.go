package stone

import (
	"stone/token"
	"regexp"
	"bufio"
	"io"
)

var regexPat string =
	`\s*(?:(//.*)|([0-9]+)|("(?:\\"|\\\\|\\n|[^"])*")|[A-Z_a-z][A-Z_a-z0-9]*|==|<=|>=|&&|\|\||[[:punct:]])|\s+`

type Lexer struct {
	queue []token.Token
	hasMore bool
	pattern *regexp.Regexp
	reader *bufio.Scanner
}

func NewLexer(r io.Reader) *Lexer{
	return &Lexer{nil, true, regexp.MustCompile(regexPat), bufio.NewScanner(r)}
}

func (self *Lexer) Read() token.Token {

}

func (self *Lexer) Peek() token.Token {

}

func (self *Lexer) fillQueue() bool {

}

func (self *Lexer) readLine() {

}

func (self *Lexer) addToken() {

}

func (self *Lexer) toStringLiteral(s string) string {

}
