package ast

type BlockStmnt struct {
	astList
}

func NewBlockStmnt(list []ASTree) *BlockStmnt{
	return &BlockStmnt{astList{list}}
}
