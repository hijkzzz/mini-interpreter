package environment

type NestedEnv struct {
	values map[string]interface{}
	outer Environment
}

func NewNestedEnv() *NestedEnv {
	return &NestedEnv{make(map[string]interface{}), nil}
}

func (self *NestedEnv) Get(name string) interface{} {
	v, ok := self.values[name]
	if !ok && self.outer != nil {
		return self.outer.Get(name)
	} else {
		return v
	}
}

func (self *NestedEnv) SetNew(name string, value interface{}) {
	self.values[name] = value
}

func (self *NestedEnv) Set(name string, value interface{}) {
	e := self.Where(name)
	if e == nil {
		e = self
	}
	e.SetNew(name, value)
}

func (self *NestedEnv) Where(name string) Environment {
	_, ok := self.values[name]
	if ok {
		return self
	} else if self.outer != nil {
		return self.outer.Where(name)
	} else {
		return nil
	}
}

func (self *NestedEnv) SetOuter(e Environment) {
	self.outer = e
}
