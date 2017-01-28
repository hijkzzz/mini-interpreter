package ast

import (
	"stone/environment"
)

type NullStmnt struct {
	astList
}

func NewNullStmnt () *NullStmnt {
	return &NullStmnt{astList{nil}}
}

func (self *NullStmnt) Eval(env environment.Environment, args... interface{}) interface{} {
	return nil
}

func (self *NullStmnt) String() string {
	return "(null)"
}
