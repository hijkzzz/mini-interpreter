package ast

import "stone/token"

type OP struct {
	astLeaf
}

func NewOP(token token.Token) *OP {
	return &OP{astLeaf{token}}
}
