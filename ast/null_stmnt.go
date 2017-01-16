package ast

type NullStmnt struct {
	astList
}

func NewNullStmnt (list []ASTree) *NullStmnt {
	return &NullStmnt{astList{list}}
}
