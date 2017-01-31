package token

/*
	单词接口
	数字，字符串，标识符三种类型
 */

var EOF Token = &token{-1}
var EOL string = "\\n"

type Token interface {
	GetLineNumber() int
	GetNumber() int
	GetText() string
	IsIdentifier() bool
	IsNumber() bool
	IsString() bool
}

type token struct {
	lineNumber int
}

func (self *token) GetLineNumber() int {
	return self.lineNumber
}

func (self *token) GetNumber() int {
	panic("not number token")
}

func (self *token) GetText() string {
	return ""
}

func (self *token) IsIdentifier() bool {
	return false
}

func (self *token) IsNumber() bool {
	return false
}

func (self *token) IsString() bool {
	return false
}

