package ast

import "stone/token"

type NumberLiteral struct {
	astLeaf
}

func NewNumberLiteral(t token.Token) *NumberLiteral {
	return &NumberLiteral{astLeaf{t}}
}

func (self *NumberLiteral) Value () int {
	return self.token.GetNumber()
}
