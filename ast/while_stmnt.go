package ast

import (
	"reflect"
	"stone/environment"
)

type WhileStmnt struct {
	astList
}

func NewWhileStmnt(list []ASTree) *WhileStmnt {
	return &WhileStmnt{astList{list}}
}

func (self *WhileStmnt) Condition() ASTree {
	return self.Child(0)
}

func (self *WhileStmnt) Body() ASTree {
	return self.Child(1)
}

func (self *WhileStmnt) String() string {
	b := self.Body()
	return "(while " + self.Condition().String() + " " + b.String() + ")"
}

func (self *WhileStmnt) Eval(env environment.Environment, args... interface{}) interface{} {
	var result interface{}
	for true {
		c := self.Condition().Eval(env)
		if reflect.TypeOf(c).Kind() == reflect.Int && c.(int) == 0 {
			return result
		} else {
			result = self.Body().Eval(env)
		}
	}
	panic("runtime error")
}
