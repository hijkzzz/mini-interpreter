package ast

import "stone/environment"

type CallFunc struct {
	astList
}

func NewCallFunc(list []ASTree) *CallFunc {
	return &CallFunc{astList{list}}
}

func (self *CallFunc) Eval(env environment.Environment) interface{} {

}
