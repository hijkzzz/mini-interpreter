package environment

type Environment interface {
	Get(name string) interface{}
	Set(name string, value interface{})
	SetNew(name string, value interface{})

	Where(name string) Environment
	SetOuter(e Environment)
}
