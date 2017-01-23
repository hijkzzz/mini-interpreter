package ast

import (
	"stone/environment"
)

type BlockStmnt struct {
	astList
}

func NewBlockStmnt(list []ASTree) *BlockStmnt{
	return &BlockStmnt{astList{list}}
}

func (self *BlockStmnt) Eval(env environment.Environment, args... interface{}) interface{} {
	var result interface{}

	// 返回最后一条语句执行结果
	for _, value := range self.children {
		result = value.Eval(env)
	}
	return result
}
