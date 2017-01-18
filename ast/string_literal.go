package ast

import (
	"stone/token"
	"stone/environment"
)

type StringLiteral struct {
	astLeaf
}

func NewStringLiteral(token token.Token) *StringLiteral {
	return &StringLiteral{astLeaf{token}}
}

func (self *StringLiteral) Value() string {
	return self.Token().GetText()
}

func (self *StringLiteral) Eval(env environment.Environment) interface{} {
	return self.Value()
}
