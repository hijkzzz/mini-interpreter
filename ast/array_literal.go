package ast

type ArrayLiteral struct {
	astList
}

func NewArrayLiteral(list []ASTree) *ArrayLiteral {
	return &ArrayLiteral{astList{list}}
}

func (self *ArrayLiteral) Size() int {
	return self.NumChildren()
}
