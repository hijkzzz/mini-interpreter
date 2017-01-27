package ast

import (
	"stone/environment"
)

type Func struct {
	astList
}

func NewFunc(list []ASTree) *Func {
	return &Func{astList{list}}
}

func (self *Func) Parameters() *ParameterList {
	return self.Child(0).(*ParameterList)
}

func (self *Func) Body() *BlockStmnt {
	return self.Child(1).(*BlockStmnt)
}

func (self *Func) String() string {
	return "(func " + self.Parameters().String() + " " + self.Body().String() + ")"
}

func (self *Func) Eval(env environment.Environment, args... interface{}) interface{} {
	return NewFunction(self.Parameters().String(), self.Parameters(), self.Body(), env)
}
