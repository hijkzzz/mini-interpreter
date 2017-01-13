package token

type NumToken struct {
	BaseToken
	value int
}

func NewNumToken(lineNumber int, value int) *NumToken {
	return &NumToken{BaseToken{lineNumber}, value}
}

func (self *NumToken) IsNumber() bool {
	return true
}

func (self *NumToken) GetNumber() int {
	return self.value
}
