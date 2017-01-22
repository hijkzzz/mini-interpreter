package ast

import "stone/environment"

type Arguments struct {
	astList
}

func NewArguments(list []ASTree) *Arguments {
	return &Arguments{astList{list}}
}

func (self *Arguments) Size(i int) int {
	return self.NumChildren()
}

func (self *Arguments) Eval(env environment.Environment) interface{} {

}
