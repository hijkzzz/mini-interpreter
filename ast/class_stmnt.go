package ast

import "stone/environment"

type ClassStmnt struct {
	astList
}

func NewClassStmnt(list []ASTree) *ClassStmnt {
	return &ClassStmnt{astList{list}}
}

func (self *ClassStmnt) Name() string {
	return self.Child(0).(*Name).token.GetText()
}

func (self *ClassStmnt) SuperClass() string {
	if self.NumChildren() < 3 {
		return ""
	} else {
		return self.Child(2).(*Name).token.GetText()
	}
}

func (self *ClassStmnt) Body() *ClassBody {
	return self.Child(1).(*ClassBody)
}

func (self *ClassStmnt) String() string {
	return "(class " + self.Name() + " " + self.SuperClass() + " " + self.Body().String() + ")"
}

func (self *ClassStmnt) Eval(env environment.Environment, args... interface{}) interface{} {
	ci := NewClassInfo(self, env)
	env.Set(self.Name(), ci)
	return self.Name()
}
