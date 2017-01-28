package ast

import (
	"stone/environment"
	"reflect"
)

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

func (self *ArrayRef) Eval(env environment.Environment, args... interface{}) interface{} {
	if value, ok := args[0].([]interface{}); ok {
		index := self.Index().Eval(env)
		if reflect.TypeOf(index).Kind() == reflect.Int && index.(int) >= 0 && index.(int) < len(value){
			return value[index.(int)]
		} else {
			panic("bad index")
		}
	} else {
		panic("bad array access")
	}
}