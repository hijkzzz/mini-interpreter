package ast

import "stone/environment"

type ClassBody struct {
	astList
}

func NewClassBody(list []ASTree) *ClassBody {
	return &ClassBody{astList{list}}
}

func (self *ClassBody) Eval(env environment.Environment, args... interface{}) interface{} {
	for _, t := range self.Children() {
		t.Eval(env)
	}
	return nil
}
