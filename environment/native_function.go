package environment

import "reflect"

type NativeFunction struct {
	name string
	method reflect.Value
	numParams int
}

func NewNativeFunction(name string, method reflect.Value) *NativeFunction {
	if method.Kind() != reflect.Func {
		panic("not a func")
	}
	return &NativeFunction{name, method, method.Type().NumIn()}
}

func (self *NativeFunction) NumParammeters() int {
	return self.numParams
}

func (self *NativeFunction) String() string {
	return "<native " + self.name + ">"
}

func (self *NativeFunction) Invoke(in []reflect.Value) interface{} {
	reslut := self.method.Call(in)
	return reslut[0].Interface()
}
