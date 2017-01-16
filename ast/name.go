package ast

import "stone/token"

type Name struct {
	astLeaf
}

func NewName(token token.Token) *Name {
	return &Name{astLeaf{token}}
}

func (self *Name) Name() string {
	return self.token.GetText()
}
