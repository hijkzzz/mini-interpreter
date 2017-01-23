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

func (self *ParameterList) Size() int {
	return self.NumChildren()
}

func (self *ParameterList) Eval(env environment.Environment, args... interface{}) interface{} {
	index := args[0].(int)
	env.SetNew(self.Name(index), args[1])
	return nil
}
