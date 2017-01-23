package ast

import (
	"stone/environment"
)

/*
	抽象语法树节点
	astree, ast_list, ast_leaf 为抽象类
 */

type ASTree interface {
	Child(i int) ASTree
	Children() []ASTree
	NumChildren() int
	Location() string
	String() string

	Eval(env environment.Environment, args... interface{}) interface{}
}
