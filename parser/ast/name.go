package ast

import "stone/lexer/token"

type Name struct {
	astLeaf
}

func NewName(t token.Token) *Name {
	return &Name{astLeaf{t}}
}

func (self *Name) Name () string {
	return self.token.GetText()
}
