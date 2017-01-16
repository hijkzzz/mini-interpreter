package ast

import "fmt"

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
	return fmt.Sprintf("(while %v %v)", self.Condition(), self.Body())
}
