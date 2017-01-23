package ast

import (
	"stone/environment"
)

type CallFunc struct {
	astList
}

func NewCallFunc(list []ASTree) *CallFunc {
	return &CallFunc{astList{list}}
}

func (self *CallFunc) Name() string {
	return self.Child(0).(*Name).Token().GetText()
}

func (self *CallFunc) Arguments() *Arguments {
	return self.Child(1).(*Arguments)
}

func (self *CallFunc) Eval(env environment.Environment, args... interface{}) interface{} {
	return self.Arguments().Eval(env, env.Get(self.Name()))
}
