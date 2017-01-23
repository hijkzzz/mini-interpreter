package ast

import (
	"fmt"
	"stone/environment"
)

type NegativeExpr struct {
	astList
}

func NewNegativeExpr(list []ASTree) *NegativeExpr {
	return &NegativeExpr{astList{list}}
}

func (self *NegativeExpr) Operand() ASTree {
	return self.Child(0)
}

func (self *NegativeExpr) String() string {
	return fmt.Sprintf("(-%v)", self.Operand())
}

func (self *NegativeExpr) Eval(env environment.Environment, args... interface{}) interface{} {
	v := self.Operand().Eval(env)
	switch v.(type) {
	case int:
		return -v.(int)
	}
	panic("bad type for -")
}
