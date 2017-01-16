package ast

import "stone/token"

type NumberLiteral struct {
	astLeaf
}

func NewNumberLiteral (token token.Token) *NumberLiteral {
	return &NumberLiteral{astLeaf{token}}
}

func (self *NumberLiteral) Value() int {
	return self.token.GetNumber()
}
