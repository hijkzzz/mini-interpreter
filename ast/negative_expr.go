package ast

import "fmt"

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
	return fmt.Sprintf("-%v", self.Operand())
}
