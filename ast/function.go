package ast

import (
	"stone/environment"
)

type Function struct {
	name string
	parameters *ParameterList
	body *BlockStmnt
	env environment.Environment
}

func NewFunction(name string, parameters *ParameterList, body *BlockStmnt, env environment.Environment) *Function {
	return &Function{name, parameters, body, env}
}

func (self *Function) Parameters() *ParameterList {
	return self.parameters
}

func (self *Function) Body() *BlockStmnt {
	return self.body
}

func (self *Function) MakeEnv() environment.Environment {
	return environment.NewNestedEnv()
}

func (self *Function) String() string {
	return "<func " + self.name  + ">"
}
