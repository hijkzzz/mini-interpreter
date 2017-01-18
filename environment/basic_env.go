package environment

type BasicEnv struct {
	values map[string]interface{}
}

func NewBasicEnv() *BasicEnv {
	return &BasicEnv{make(map[string]interface{})}
}

func (self *BasicEnv) Get(name string) interface{} {
	return self.values[name]
}

func (self *BasicEnv) Set(name string, value interface{}) {
	self.values[name] = value
}
