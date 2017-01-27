package ast

import "stone/environment"

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

func (self *Dot) Eval(env environment.Environment, args... interface{}) interface{} {
	member := self.Name()
	if ci, ok := args[0].(*ClassInfo); ok {
		if member == "new" {
			e := environment.NewNestedEnv()
			e.SetOuter(ci.Environment())
			obj := NewStoneObject(e)
			e.SetNew("this", obj)
			self.initObject(ci, e)
			return obj
		} else {
			panic(member + " not a class func")
		}
	} else if so, ok := args[0].(*StoneObject); ok {
		return so.Read(member)
	} else {
		panic("unkown ." + self.Name())
	}
}

func (self *Dot) initObject(ci *ClassInfo, env environment.Environment) {
	if ci.SuperClass() != nil {
		self.initObject(ci.SuperClass(), env)
	}
	ci.Body().Eval(env)
}
