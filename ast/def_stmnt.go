package ast

import "stone/environment"

type DefStmnt struct {
	astList
}

func NewDefStmnt(list []ASTree) *DefStmnt {
	return &DefStmnt{astList{list}}
}

func (self *DefStmnt) Name() string {
	return self.Child(0).(ASTLeaf).Token().GetText()
}

func (self *DefStmnt) Parameters() ParameterList {
	return self.Child(1).(ParameterList)
}

func (self *DefStmnt) Body() BlockStmnt {
	return self.Child(2).(BlockStmnt)
}

func (self *DefStmnt) String() string {
	return "(def " + self.Name() + " " + self.Parameters().String() + self.Body().String() + ")"
}

func (self *DefStmnt) Eval(env environment.Environment) interface{} {

}
