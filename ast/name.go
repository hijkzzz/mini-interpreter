package ast

import (
	"stone/token"
	"stone/environment"
)

type Name struct {
	astLeaf
}

func NewName(token token.Token) *Name {
	return &Name{astLeaf{token}}
}

func (self *Name) Name() string {
	return self.token.GetText()
}

func (self *Name) Eval(env environment.Environment) interface{} {
	value := env.Get(self.Name())
	if value == nil {
		panic("undefined name " + self.Name())
	} else {
		return value
	}
}
