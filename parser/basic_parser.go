package parser

import (
	"stone/lexer"
	"stone/token"
	"stone/ast"
	"strconv"
)

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

type Precedence struct{
	value int
	leftAssoc bool
}

func NewPrecedence(v int, a bool) *Precedence {
	return &Precedence{v, a}
}

type Parser struct {
	lexer *lexer.Lexer
	reserved map[string]bool
	operators map[string]*Precedence
}

func NewParser(lexer *lexer.Lexer) *Parser {
	reserved := map[string]bool {
		";" : true,
		"}" : true,
		token.EOL : true,
	}

	operators := map[string]*Precedence {
		"<" : NewPrecedence(1, true),
		">" : NewPrecedence(1, true),
		"+" : NewPrecedence(2, true),
		"-" : NewPrecedence(2, true),
		"*" : NewPrecedence(3, true),
		"/" : NewPrecedence(3, true),
		"^" : NewPrecedence(4, false),
	}
	return &Parser{lexer, reserved, operators}
}

func (self *Parser) primary() ast.ASTree{
	list := []ast.ASTree{}

	t := self.lexer.Read()
	if t.IsIdentifier() && t.GetText() == "(" {
		list[0] = self.expr()
		self.readToken(")")
	} else if t.IsIdentifier() {
		list[0] = ast.NewName(t)
	} else if t.IsString() {
		list[0] = ast.NewStringLiteral(t)
	} else if t.IsNumber() {
		list[0] = ast.NewNumberLiteral(t)
	}

	return ast.NewPrimaryExpr(list)
}

func (self *Parser) factor() ast.ASTree{

	if self.isToken("-") {
		self.lexer.Read()
		return ast.NewNegativeExpr([]ast.ASTree{self.primary()})
	} else {
		return self.primary()
	}
}

// 算法优先分析法
// 用于处理运算符优先级
func (self *Parser) expr() ast.ASTree{
	right := self.factor()
	next := self.nextOperator()
	for next != nil {
		right = self.doShift(right, next.value)
		next = self.nextOperator()
	}
	return right
}

// 如果下一个算符优先级高
// 则先归约右表达式
func (self *Parser) doShift(left ast.ASTree, prec int) ast.ASTree{
	op := ast.NewOP(self.lexer.Read())
	right := self.factor()
	next := self.nextOperator()
	for next != nil && self.rightIsExpr(prec, next) {
		right = self.doShift(right, next.value)
		next = self.nextOperator()
	}
	return ast.NewBinaryExpr([]ast.ASTree{left, op, right})
}

func (self *Parser) nextOperator() *Precedence {
	t := self.lexer.Peek(0)
	if t.IsIdentifier() {
		return self.operators[t.GetText()]
	} else {
		return nil
	}
}

func (self *Parser) rightIsExpr(prec int, nextPrec *Precedence) bool {
	if nextPrec.leftAssoc {
		return prec < nextPrec.value
	} else {
		return prec <= nextPrec.value
	}
}

func (self *Parser) block() ast.ASTree{
	self.readToken("{")
}

func (self *Parser) simple() ast.ASTree{
	return self.expr()
}

func (self *Parser) statement() ast.ASTree{

}

func (self *Parser) program() ast.ASTree{

}

func (self *Parser) readToken(name string) {
	t := self.lexer.Read()
	if !(t.IsIdentifier() && name == t.GetText()) {
		panic("parser error at line " + strconv.Itoa(t.GetLineNumber()))
	}
}

func (self *Parser) isToken(name string) bool{
	t := self.lexer.Peek(0)
	return t.IsIdentifier() && name == t.GetText()
}
