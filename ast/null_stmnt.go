package ast

import (
	"stone/environment"
)

type NullStmnt struct {
	astList
}

func NewNullStmnt (list []ASTree) *NullStmnt {
	return &NullStmnt{astList{list}}
}

func (self *NullStmnt) Eval(env environment.Environment) interface{} {
	return nil
}
