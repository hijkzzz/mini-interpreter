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
	result := self.Child(0).(*Name).Eval(env)

	for index := 1; self.HasPostfix(index); index++ {
		a := self.Child(index)
		if args, ok := a.(Arguments); ok {
			result = args.Eval(env, result)
		} else if dot, ok := a.(Dot); ok {
			result = dot.Eval(env, result)
		} else if arrref, ok := a.(ArrayRef); ok{
			
		} else {
			panic("error postfix")
		}
	}
	return result
}
