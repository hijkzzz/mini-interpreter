package parser

import "stone/lexer"

type Element interface {
	Parse (lexer *lexer.Lexer, res []ASTree)
	Match (lexer *lexer.Lexer) bool
}
