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
	return index < self.NumChildren()
}

func (self *Primary) Eval(env environment.Environment, args... interface{}) interface{} {
	return self.Eval2(0, env, args)
}

func (self *Primary) Eval2(offset int, env environment.Environment, args... interface{}) interface{} {
	result := self.Child(0).(*Name).Eval(env)

	for index := 1; self.HasPostfix(index + offset); index++ {
		a := self.Child(index)
		if args, ok := a.(*Arguments); ok {
			result = args.Eval(env, result)
		} else if dot, ok := a.(*Dot); ok {
			result = dot.Eval(env, result)
		} else if arrref, ok := a.(*ArrayRef); ok{
			result = arrref.Eval(env, result)
		} else {
			panic("error postfix")
		}
	}
	return result
}
