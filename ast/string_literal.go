package ast

import "stone/token"

type StringLiteral struct {
	astLeaf
}

func NewStringLiteral(token token.Token) *StringLiteral {
	return &StringLiteral{astLeaf{token}}
}

func (self *StringLiteral) Value() string {
	return self.Token().GetText()
}
