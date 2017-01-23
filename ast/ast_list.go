package ast

import (
	"bytes"
	"stone/environment"
)

/*
	抽象语法树——非叶节点基类
 */
type ASTList interface {
	ASTree
}

type astList struct {
	children []ASTree
}

func (self *astList) Child(i int) ASTree {
	return self.children[i]
}

func (self *astList) NumChildren() int{
	return len(self.children)
}

func (self *astList) Children() []ASTree {
	return self.children
}

func (self *astList) Location() string {
	for _, t := range self.children {
		if t.Location() != "" {
			return t.Location()
		}
	}
	return ""
}

func (self *astList) String() string {
	var buf bytes.Buffer
	buf.WriteString("(")
	sep := ""
	for _, t := range self.children {
		buf.WriteString(sep)
		buf.WriteString(t.String())
		sep = " "
	}
	buf.WriteString(")")
	return buf.String()
}

func (self *astList) Eval(env environment.Environment, args... interface{}) interface{} {
	panic("cannot eval " + self.String())
}
