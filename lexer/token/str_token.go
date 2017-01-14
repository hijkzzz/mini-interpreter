package token

type StrToken struct {
	BaseToken
	str string
}

func NewStrToken(lineNumber int, str string) *StrToken {
	return &StrToken{BaseToken{lineNumber}, str}
}

func (self *StrToken) IsString() bool {
	return true
}

func (self *StrToken) GetText() string {
	return self.str
}
