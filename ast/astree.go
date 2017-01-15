package ast

type ASTree interface {
	Child(i int) ASTree
	Children() []ASTree
	NumChildren() int
	Location() string
}
