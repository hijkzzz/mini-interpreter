package ast

type PrimaryExpr struct {
	astList
}

func NewPrimaryExpr(list []ASTree) *PrimaryExpr {
	return &PrimaryExpr{astList{list}}
}
