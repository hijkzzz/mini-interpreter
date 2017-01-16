package ast

type BinaryExpr struct {
	astList
}

func NewBinaryExpr(list []ASTree) *BinaryExpr {
	return &BinaryExpr{astList{list}}
}

func (self *BinaryExpr) Left() ASTree{
	return self.Child(0)
}

func (self *BinaryExpr) Operator() string {
	return self.Child(1).(ASTLeaf).Token().GetText()
}

func (self *BinaryExpr) Right() ASTree {
	return self.Child(2)
}
