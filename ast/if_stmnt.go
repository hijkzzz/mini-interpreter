package ast

import (
	"reflect"
	"stone/environment"
)

type IfStmnt struct {
	astList
}

func NewIfStmnt(list []ASTree) *IfStmnt{
	return &IfStmnt{astList{list}}
}

func (self *IfStmnt) Condition() ASTree{
	return self.Child(0)
}

func (self *IfStmnt) ThenBlock() ASTree {
	return self.Child(1)
}

func (self *IfStmnt) ElseBlock() ASTree {
	if self.NumChildren() > 2 {
		return self.Child(2)
	} else {
		return nil
	}
}

func (self *IfStmnt) String() string {
	result := "(if " + self.Condition().String() + " " + self.ThenBlock().String()
	if self.ElseBlock() != nil {
		result += " else " + self.ElseBlock().String()
	}
	return result + ")"
}

func (self *IfStmnt) Eval(env environment.Environment) interface{} {
	c := self.Condition().Eval(env)
	if reflect.TypeOf(c).Kind() == reflect.Int && c.(int) == 1 {
		return self.ThenBlock().Eval(env)
	} else {
		e := self.ElseBlock()
		if e == nil {
			return 0
		} else {
			return e.Eval(env)
		}
	}
}


