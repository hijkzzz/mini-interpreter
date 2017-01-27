package ast

import (
	"stone/environment"
	"reflect"
)

type Arguments struct {
	astList
}

func NewArguments(list []ASTree) *Arguments {
	return &Arguments{astList{list}}
}

func (self *Arguments) Size() int {
	return self.NumChildren()
}

func (self *Arguments) Eval(callerEnv environment.Environment, args... interface{}) interface{} {
	value := reflect.ValueOf(args[0])
	if value.Type().ConvertibleTo(reflect.TypeOf((*Function)(nil))) {
		fnc := args[0].(*Function)
		params := fnc.Parameters()
		if params.Size() != self.Size() {
			panic("bad number of arguments")
		}

		newEnv := fnc.MakeEnv()
		for num, a := range self.Children() {
			params.Eval(newEnv, num, a.Eval(callerEnv))
		}
		return fnc.Body().Eval(newEnv)
	} else if value.Type().ConvertibleTo(reflect.TypeOf((*environment.NativeFunction)(nil))) {
		fnc := args[0].(*environment.NativeFunction)
		nparams := fnc.NumParammeters()
		if nparams != self.Size() {
			panic("bad number of arguments")
		}
		params := make([]reflect.Value, nparams)
		for i := 0; i < len(params); i++ {
			params[i] = reflect.ValueOf(self.Child(i).Eval(callerEnv))
		}
		return fnc.Invoke(params)
	} else {
		panic("bad function")
	}
}
