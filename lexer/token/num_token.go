package token

import "strconv"

type NumToken struct {
	token
	value int
}

func NewNumToken(lineNumber int, value int) *NumToken {
	return &NumToken{token{lineNumber}, value}
}

func (self *NumToken) IsNumber() bool {
	return true
}

func (self *NumToken) GetNumber() int {
	return self.value
}

func (self *NumToken) GetText() string {
	return strconv.Itoa(self.value)
}
