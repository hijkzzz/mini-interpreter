package ast

import (
	"stone/environment"
)

/*
	用户自定义函数
 */

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
	e := environment.NewNestedEnv()
	e.SetOuter(self.env)
	return e
}

func (self *Function) String() string {
	return "<func " + self.name  + ">"
}
