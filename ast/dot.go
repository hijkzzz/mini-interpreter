package ast

type Dot struct {
	astList
}

func NewDot(list []ASTree) *Dot {
	return &Dot{astList{list}}
}

func (self *Dot) Name() string {
	return self.Child(0).(*Name).token.GetText()
}

func (self *Dot) String() string {
	return "." + self.Name()
}
