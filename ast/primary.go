package ast

import (
	"stone/environment"
)

type Primary struct {
	astList
}

func NewPrimary(list []ASTree) *Primary {
	return &Primary{astList{list}}
}

func (self *Primary) Name() string {
	return self.Child(0).(*Name).Token().GetText()
}

func (self *Primary) HasPostfix(index int) bool {
	return self.NumChildren() - index > 0
}

func (self *Primary) Eval(env environment.Environment, args... interface{}) interface{} {
	result := self.Child(1).Eval(env, env.Get(self.Name()))

	for index := 2; self.HasPostfix(index); index++ {
		args := self.Child(index).(*Arguments)
		result = args.Eval(env, result)
	}
	return result
}
