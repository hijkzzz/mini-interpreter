package ast

import (
	"stone/lexer/token"
	"strconv"
)

type ASTLeaf interface {
	ASTree
	Token() token.Token
}

type astLeaf struct {
	token token.Token
}

func (self *astLeaf) Child(i int) ASTree {
	panic("index out of range")
}

func (self *astLeaf) NumChildren() int{
	return 0
}

func (self *astLeaf) Children() []ASTree {
	return nil
}

func (self *astLeaf) Location() string {
	return "at line " + strconv.Itoa(self.token.GetLineNumber())
}

func (self *astLeaf) Token() token.Token {
	return self.token
}
