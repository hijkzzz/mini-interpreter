package ast

import "stone/environment"

type VMFunction struct {
	Function
	entry int
}

func NewVMFunction(name string,
	parameters *ParameterList,
	body *BlockStmnt,
	env environment.Environment,
	entry int) *VMFunction {
	return &VMFunction{Function{name, parameters, body, env}, entry}
}

func (self *VMFunction) Entry() int {
	return self.entry
}
