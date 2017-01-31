package lexer

import (
	"stone/token"

	"regexp"
	"bufio"
	"io"
	"strconv"
	"bytes"
)

/*
	词法分析器
	源文件被分割为——数字，字符串，标识符序列
 */

// 注释|数字|字符串|标识符（且不许为空）
var regexPat string =
	`\s*(?:(//.*)|([0-9]+)|"((?:\\"|\\\\|\\n|[^"])*)"|([A-Z_a-z][A-Z_a-z0-9]*|==|<=|>=|&&|\|\||[[:punct:]]))|\s+`

type Lexer struct {
	queue []token.Token
	hasMore bool
	pattern *regexp.Regexp
	reader *bufio.Scanner
	lineNo int
}

func NewLexer(r io.Reader) *Lexer{
	return &Lexer{nil, true, regexp.MustCompile(regexPat), bufio.NewScanner(r), 0}
}

func (self *Lexer) Read() token.Token {
	if self.fillQueue(0) {
		defer func() {self.queue = self.queue[1:]}()
		return self.queue[0]
	} else {
		return token.EOF
	}
}

func (self *Lexer) Peek(i int) token.Token {
	if self.fillQueue(i) {
		return self.queue[i]
	} else {
		return token.EOF
	}
}

func (self *Lexer) fillQueue(i int) bool {
	for i >= len(self.queue) {
		if self.hasMore {
			self.readLine()
		} else {
			return false
		}
	}
	return true
}

func (self *Lexer) readLine() {
	if self.reader.Scan() {
		self.lineNo++
		line := self.reader.Text()
		groups := self.pattern.FindStringSubmatch(line)
		for len(line) > 0 {
			if (groups == nil) {
				panic("lexer error at line " + strconv.Itoa(self.lineNo) + " " + line)
			}
			self.addToken(groups)
			line = line[len(groups[0]):]
			groups = self.pattern.FindStringSubmatch(line)
		}
		self.queue = append(self.queue, token.NewIdToken(self.lineNo, token.EOL))
	} else {
		self.hasMore = false
	}
}

func (self *Lexer) addToken(groups []string) {
	if groups[1] != "" {

	} else if groups[2] != "" {
		number, err := strconv.Atoi(groups[2])
		if err != nil {
			panic("Atoi error")
		}
		self.queue = append(self.queue, token.NewNumToken(self.lineNo, number))
	} else if groups[3] != "" {
		self.queue = append(self.queue, token.NewStrToken(self.lineNo, self.toStringLiteral(groups[3])))
	} else if groups[4] != "" {
		self.queue = append(self.queue, token.NewIdToken(self.lineNo, groups[4]))
	}
	//white space
}

func (self *Lexer) toStringLiteral(s string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '\\' && i + 1 < len(s) {
			c2 := s[i + 1]
			if c2 == '"' || c2 == '\\' {
				i++
				c = s[i]
			} else if c2 == 'n' {
				i++
				c = '\n'
			}
		}
		buffer.WriteByte(c)
	}
	return buffer.String()
}

func (self *Lexer) GetLineNumber() string {
	return strconv.Itoa(self.lineNo)
}
