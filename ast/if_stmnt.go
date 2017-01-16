package ast

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
	return "(if " + self.Condition().String() + " " + self.ThenBlock().String() +
		" else " + self.ElseBlock().String() + ")"
}


