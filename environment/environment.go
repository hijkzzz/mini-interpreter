package environment

type Environment interface {
	Set(name string, value interface{})
	Get(name string) interface{}
}
