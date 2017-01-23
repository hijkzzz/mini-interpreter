package ast

import "stone/environment"

type Arguments struct {
	astList
}

func NewArguments(list []ASTree) *Arguments {
	return &Arguments{astList{list}}
}

func (self *Arguments) Size() int {
	return self.NumChildren()
}

func (self *Arguments) Eval(callerEnv environment.Environment, args... interface{}) interface{} {
	fnc, ok := args[0].(*Function)
	if !ok {
		panic("bad function")
	}
	params := fnc.Parameters()
	if params.Size() != self.Size() {
		panic("bad number of arguments")
	}

	newEnv := fnc.MakeEnv()
	for num, a := range self.Children() {
		params.Eval(newEnv, num, a.Eval(callerEnv))
	}
	return fnc.Body().Eval(newEnv)
}
