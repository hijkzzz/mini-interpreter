package ast

import (
	"stone/environment"
	"fmt"
)

type StoneObject struct {
	env environment.Environment
}

func NewStoneObject(e environment.Environment) *StoneObject {
	return &StoneObject{e}
}

func (self *StoneObject) String() string {
	return fmt.Sprintf("<object %v>", self)
}

func (self *StoneObject) Read(name string) interface{} {
	return self.GetEnv(name).Get(name)
}

func (self *StoneObject) Write(name string, value interface{}) {
	self.GetEnv(name).SetNew(name, value)
}

func (self *StoneObject) GetEnv(name string) environment.Environment {
	e := self.env.Where(name)
	if e == self.env {
		return e
	} else {
		panic("class member access error")
	}
}
