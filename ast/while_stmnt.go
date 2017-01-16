package ast


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
	return "(while " + self.Condition().String() + " " + self.Body().String() + ")"
}
