package ast

type ArrayRef struct {
	astList
}

func NewArrayRef(list []ASTree) *ArrayRef {
	return &ArrayRef{astList{list}}
}

func (self *ArrayRef) Index() ASTree {
	return self.Child(0)
}

func (self *ArrayRef) String() string {
	return "[" + self.Index().String() + "]"
}
