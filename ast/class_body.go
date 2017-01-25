package ast

type ClassBody struct {
	astList
}

func NewClassBody(list []ASTree) *ClassBody {
	return &ClassBody{astList{list}}
}
