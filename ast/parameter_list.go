package ast

import "stone/environment"

type ParameterList struct {
	astList
}

func NewParameterList(list []ASTree) *ParameterList {
	return &ParameterList{astList{list}}
}

func (self *ParameterList) Name(i int) string {
	return self.Child(i).(ASTLeaf).Token().GetText()
}

func (self *ParameterList) Size(i int) int {
	return self.NumChildren()
}

func (self *ParameterList) Eval(env environment.Environment) interface{} {

}
