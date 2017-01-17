package parser

import (
	"stone/lexer"
	"stone/token"
	"stone/ast"
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

type BasicParser struct {
	lexer *lexer.Lexer
	reserved map[string]bool
	operators map[string]*Precedence
}

func NewBasicParser(lexer *lexer.Lexer) *BasicParser {
	reserved := map[string]bool {
		";" : true,
		"}" : true,
		token.EOL : true,
	}

	operators := map[string]*Precedence {
		"=" : NewPrecedence(1, false),
		"==": NewPrecedence(2, true),
		"<" : NewPrecedence(2, true),
		">" : NewPrecedence(2, true),
		"+" : NewPrecedence(3, true),
		"-" : NewPrecedence(3, true),
		"*" : NewPrecedence(4, true),
		"/" : NewPrecedence(4, true),
		"%" : NewPrecedence(4, true),
	}
	return &BasicParser{lexer, reserved, operators}
}

func (self *BasicParser) primary() ast.ASTree{
	list := make([]ast.ASTree, 1)
	t := self.lexer.Read()
	if t.IsIdentifier() && t.GetText() == "(" {
		list[0] = self.expr()
		self.readToken(")")
	} else if t.IsIdentifier() && self.reserved[t.GetText()] == false {
		list[0] = ast.NewName(t)
	} else if t.IsString() {
		list[0] = ast.NewStringLiteral(t)
	} else if t.IsNumber() {
		list[0] = ast.NewNumberLiteral(t)
	} else {
		panic("parser error at line " + self.lexer.GetLineNumber())
	}

	return ast.NewPrimaryExpr(list)
}

func (self *BasicParser) testPrimary() bool {
	t := self.lexer.Peek(0)
	if t.IsIdentifier() && (self.reserved[t.GetText()] == true) {
		return false
	}
	return true
}

func (self *BasicParser) factor() ast.ASTree{

	if self.isToken("-") {
		self.lexer.Read()
		return ast.NewNegativeExpr([]ast.ASTree{self.primary()})
	} else {
		return self.primary()
	}
}

func (self *BasicParser) testFactor() bool {
	return self.isToken("-") || self.testPrimary()
}

// 算法优先分析法
// 用于处理运算符优先级
func (self *BasicParser) expr() ast.ASTree{
	right := self.factor()
	next := self.nextOperator()
	for next != nil {
		right = self.doShift(right, next.value)
		next = self.nextOperator()
	}
	return right
}

func (self * BasicParser) testExpr() bool {
	return self.testFactor()
}

func (self *BasicParser) doShift(left ast.ASTree, prec int) ast.ASTree{
	op := ast.NewOP(self.lexer.Read())
	right := self.factor()
	next := self.nextOperator()
	for next != nil && self.rightIsExpr(prec, next) {
		right = self.doShift(right, next.value)
		next = self.nextOperator()
	}
	return ast.NewBinaryExpr([]ast.ASTree{left, op, right})
}

func (self *BasicParser) nextOperator() *Precedence {
	t := self.lexer.Peek(0)
	if t.IsIdentifier() {
		return self.operators[t.GetText()]
	} else {
		return nil
	}
}

func (self *BasicParser) rightIsExpr(prec int, nextPrec *Precedence) bool {
	if nextPrec.leftAssoc {
		return prec < nextPrec.value
	} else {
		return prec <= nextPrec.value
	}
}

func (self *BasicParser) block() ast.ASTree{
	self.readToken("{")
	list := make([]ast.ASTree, 0)
	if self.testStatement() {
		list = append(list, self.statement())
	}

	t := self.lexer.Peek(0)
	for t.IsIdentifier() && (t.GetText() == ";" || t.GetText() == token.EOL){
		self.lexer.Read()
		if self.testStatement() {
			list = append(list, self.statement())
		}
		t = self.lexer.Peek(0)
	}

	if len(list) == 0 {
		list = append(list, ast.NewNullStmnt([]ast.ASTree{}))
	}
	self.readToken("}")
	return ast.NewBlockStmnt(list)
}

func (self *BasicParser) testBlock() bool {
	return self.isToken("{")
}

func (self *BasicParser) simple() ast.ASTree{
	return self.expr()
}

func (self *BasicParser) testSimple() bool {
	return self.testExpr()
}

func (self *BasicParser) statement() ast.ASTree{
	list := make([]ast.ASTree, 2)

	if self.isToken("if") {
		self.lexer.Read()
		list[0] = self.expr()
		list[1] = self.block()
		if self.isToken("else") {
			self.lexer.Read()
			list = append(list, self.block())
		}
		return ast.NewIfStmnt(list)
	} else if self.isToken("while") {
		self.lexer.Read()
		list[0] = self.expr()
		list[1] = self.block()
		return ast.NewWhileStmnt(list)
	} else if self.testSimple() {
		return self.simple()
	} else {
		panic("parser error at line " + self.lexer.GetLineNumber())
	}
}

func (self *BasicParser) testStatement() bool {
	return self.isToken("if") ||
		self.isToken("while") ||
		self.testSimple()
}

func (self *BasicParser) Program() ast.ASTree{
	var a ast.ASTree
	if self.testStatement() {
		a = self.statement()
	} else {
		a = ast.NewNullStmnt([]ast.ASTree{})
	}

	if self.isToken(";") || self.isToken(token.EOL) {
		self.lexer.Read()
	} else {
		panic("parser error at line " + self.lexer.GetLineNumber())
	}
	return a
}

func (self *BasicParser) TestProgram() bool {
	return self.testStatement()
}

func (self *BasicParser) readToken(name string) {
	t := self.lexer.Read()
	if !(t.IsIdentifier() && name == t.GetText()) {
		panic("parser error at line " + self.lexer.GetLineNumber())
	}
}

func (self *BasicParser) isToken(name string) bool{
	t := self.lexer.Peek(0)
	return t.IsIdentifier() && name == t.GetText()
}
