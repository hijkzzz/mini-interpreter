package ast

import (
	"stone/environment"
)

/*
	抽象语法树——函数调用节点
 */

type Call struct {
	astList
}

func NewCall(list []ASTree) *Call {
	return &Call{astList{list}}
}

func (self *Call) Name() string {
	return self.Child(0).(*Name).Token().GetText()
}

func (self *Call) Arguments() *Arguments {
	return self.Child(1).(*Arguments)
}

func (self *Call) Eval(env environment.Environment, args... interface{}) interface{} {
	return self.Arguments().Eval(env, env.Get(self.Name()))
}
