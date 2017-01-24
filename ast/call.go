package ast

import (
	"stone/environment"
)

/*
	抽象语法树——函数调用
	例如 test()()
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

func (self *Call) HasPostfix(index int) bool {
	return self.NumChildren() - index > 0
}

func (self *Call) Eval(env environment.Environment, args... interface{}) interface{} {
	result := self.Child(1).Eval(env, env.Get(self.Name()))

	for index := 2; self.HasPostfix(index); index++ {
		args := self.Child(index).(*Arguments)
		result = args.Eval(env, result)
	}
	return result
}
