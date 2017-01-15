package token

type StrToken struct {
	token
	str string
}

func NewStrToken(lineNumber int, str string) *StrToken {
	return &StrToken{token{lineNumber}, str}
}

func (self *StrToken) IsString() bool {
	return true
}

func (self *StrToken) GetText() string {
	return self.str
}
