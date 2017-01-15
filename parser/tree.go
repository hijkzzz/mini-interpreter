package parser

import (
	"stone/lexer"
	"stone/ast"
)

type Tree struct {
	parser *Parser
}

func NewTree(parser *Parser) *Tree {
	return &Tree{parser}
}

func (self *Tree) Parse(lexer *lexer.Lexer, res *[]ast.ASTree) {
	*res = append(*res, self.parser.Parse(lexer))
}

func (self *Tree) Match(lexer *lexer.Lexer) bool{
	self.parser.Match(lexer)
}
