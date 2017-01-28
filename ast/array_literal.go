package ast

import "stone/environment"

type ArrayLiteral struct {
	astList
}

func NewArrayLiteral(list []ASTree) *ArrayLiteral {
	return &ArrayLiteral{astList{list}}
}

func (self *ArrayLiteral) Size() int {
	return self.NumChildren()
}

func (self *ArrayLiteral) Eval(env environment.Environment, args... interface{}) interface{} {
	res := make([]interface{}, self.NumChildren())
	for index, t := range self.Children() {
		res[index] = t.Eval(env)
	}
	return res
}
