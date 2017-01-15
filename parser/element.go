package parser

import (
	"stone/lexer"
	"stone/ast"
)

type Element interface {
	Parse (lexer *lexer.Lexer , res *[]ast.ASTree)
	Match (lexer *lexer.Lexer) bool
}
