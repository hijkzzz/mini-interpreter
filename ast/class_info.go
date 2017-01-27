package ast

import "stone/environment"

type ClassInfo struct {
	def *ClassStmnt
	env environment.Environment
	super *ClassInfo
}

func NewClassInfo(def *ClassStmnt, env environment.Environment) *ClassInfo {
	obj := env.Get(def.SuperClass())
	if obj == nil {
		return &ClassInfo{def, env, nil}
	} else if super, ok := obj.(*ClassInfo); ok {
		return &ClassInfo{def, env, super}
	} else {
		panic("super not a class")
	}
}

func (self *ClassInfo) Name() string {
	return self.def.Name()
}

func (self *ClassInfo) SuperClass() *ClassInfo {
	return self.super
}

func (self *ClassInfo) Body() *ClassBody {
	return self.def.Body()
}

func (self *ClassInfo) Environment() environment.Environment {
	return self.env
}

func (self *ClassInfo) String() string {
	return "<class " + self.Name() + ">"
}
