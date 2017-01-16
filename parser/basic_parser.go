package parser

import "stone/lexer"

/*
	Grammer Definition

	primary : "(" expr ")" | NUMBER | ID | STRING
	factor  : "-" primary | primary
	expr 	: factor { OP factor }
	block   : "{" [ statement ] { ( ";" | EOL ) [ statement ] } "}"
	simple  : expr
	statement : "if" expr block [ "else" block ]
		| "while" expr block
		| simple
	program : [ statement ] ( ";" | EOL )
 */

type Parser struct {
	lexer *lexer.Lexer
	reserved map[string]bool
}

func (self *Parser) primary() {

}

func (self *Parser) factor() {

}

func (self *Parser) expr() {

}

func (self *Parser) block() {

}

func (self *Parser) simple() {

}

func (self *Parser) statemtnt() {

}

func (self *Parser) program() {

}

