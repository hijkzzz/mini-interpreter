package ast

import (
	"stone/token"
	"stone/environment"
)

type NumberLiteral struct {
	astLeaf
}

func NewNumberLiteral (token token.Token) *NumberLiteral {
	return &NumberLiteral{astLeaf{token}}
}

func (self *NumberLiteral) Value() int {
	return self.token.GetNumber()
}

func (self *NumberLiteral) Eval(env environment.Environment) interface{} {
	return self.Value()
}
