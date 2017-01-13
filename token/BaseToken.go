package token

type BaseToken struct {
	lineNumber int
}

func (self *BaseToken) GetLineNumber() int {
	return self.lineNumber
}

func (self *BaseToken) GetNumber() int {
	panic("not number token")
}

func (self *BaseToken) GetText() string {
	return ""
}

func (self *BaseToken) IsIdentifier() bool {
	return false
}

func (self *BaseToken) IsNumber() bool {
	return false
}

func (self *BaseToken) IsString() bool {
	return false
}

var EOF Token = &BaseToken{-1}
var EOL string = "\\n"
