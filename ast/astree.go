package ast

import (
	"stone/environment"
)

/*
	抽象语法树接口
 */

type ASTree interface {
	Child(i int) ASTree
	Children() []ASTree
	NumChildren() int
	Location() string
	String() string

	Eval(env environment.Environment, args... interface{}) interface{}
}
