package parser

import (
	"stone/lexer"
	"stone/token"
	"stone/ast"
)

/*
	Grammer Definition

	// basic
	primary : ( "[" elements ]" | "(" expr ")" | NUMBER | IDENTIFIER { postfix } | FUNC | STRING )
	factor  : "-" primary | primary
	expr 	: factor { OP factor }
	block   : "{" [ statement ] { ( ";" | EOL ) [ statement ] } "}"
	simple  : expr
	statement : "if" expr block [ "else" block ]
		| "while" expr block
		| simple
	program : [ def | defclass | statement ] ( ";" | EOL )

	// func
	parm 	: IDENTIFIER
	params  : param { "," param }
	param_list : "(" [ param ] ")"
	def 	: "def" IDENTIFIER param_list block
	args    : expr { "," expr }
	postfix : "(" [ args ] ")" | "." IDENTIFIER | "[" expr "]"

	// class
	member  : def | simple
	class_body : "{" [ member ] { ( ";" | EOL ) [ member ] "}"
	defclass : "class" IDENTIFIER [ "extends" IDENTIFIER ] class_body

	// array
	elements: expr { "," expr }
*/

type Precedence struct {
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
		"{" : true,
		"(" : true,
		")" : true,
		"+" : true,
		"-" : true,
		"*" : true,
		"/" : true,
		"%" : true,
		"==" : true,
		"=" : true,
		">" : true,
		"<" : true,
		"if" : true,
		"else" : true,
		"while" : true,
		"def" : true,
		"func": true,
		"[" : true,
		"]" : true,
		"class" : true,
		"." : true,
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
	return &Parser{lexer, reserved, operators}
}

func (self *Parser) Parse() ast.ASTree {
	return self.program()
}

// primary type
func (self *Parser) primary() ast.ASTree {
	var a ast.ASTree

	t := self.lexer.Read()
	if t.IsIdentifier() && t.GetText() == "(" {
		a = self.expr()
		self.readMatchToken(")")
	} else if t.IsIdentifier() && t.GetText() == "func" {
		a = ast.NewFunc([]ast.ASTree{self.paramList(), self.block()})
	} else if t.IsIdentifier() && t.GetText() == "[" {
		a = self.elements()
		self.readMatchToken("]")
	} else if self.isName(t) {
		a = ast.NewName(t)
		if self.testPostfix() {
			list := []ast.ASTree{a, self.postfix()}
			for self.testPostfix() {
				list = append(list, self.postfix())
			}
			a = ast.NewPrimary(list)
		}
	} else if t.IsString() {
		a = ast.NewStringLiteral(t)
	} else if t.IsNumber() {
		a = ast.NewNumberLiteral(t)
	} else {
		panic("line " + self.lexer.GetLineNumber() + " -- error" + t.GetText())
	}

	return a
}

func (self *Parser) testPrimary() bool {
	t := self.lexer.Peek(0)
	if t.IsNumber() || t.IsString() {
		return true
	} else if t.IsIdentifier() && (t.GetText() == "func" || t.GetText() == "[" || t.GetText() == "(") {
		return true
	} else if self.isName(t) {
		return true
	} else {
		return false
	}

}

func (self *Parser) factor() ast.ASTree{

	if self.isToken("-") {
		self.lexer.Read()
		return ast.NewNegativeExpr([]ast.ASTree{self.primary()})
	} else {
		return self.primary()
	}
}

func (self *Parser) testFactor() bool {
	return self.isToken("-") || self.testPrimary()
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

func (self *Parser) testExpr() bool {
	return self.testFactor()
}

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
	self.readMatchToken("{")
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
		list = append(list, ast.NewNullStmnt())
	}
	self.readMatchToken("}")
	return ast.NewBlockStmnt(list)
}

func (self *Parser) testBlock() bool {
	return self.isToken("{")
}

func (self *Parser) simple() ast.ASTree{
	return self.expr()
}

func (self *Parser) testSimple() bool {
	return self.testExpr()
}

func (self *Parser) statement() ast.ASTree{
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
		panic("line " + self.lexer.GetLineNumber() + " -- not a statement ")
	}
}

func (self *Parser) testStatement() bool {
	return self.isToken("if") ||
		self.isToken("while") ||
		self.testSimple()
}

func (self *Parser) program() ast.ASTree {
	var a ast.ASTree
	if self.testStatement() {
		a = self.statement()
	} else if self.isToken("def") {
		a = self.def()
	} else if self.isToken("class") {
		a = self.defclass()
	} else {
		a = ast.NewNullStmnt()
	}

	if self.isToken(";") || self.isToken(token.EOL) {
		self.lexer.Read()
	} else {
		panic("line " + self.lexer.GetLineNumber() + " -- lack of end symbol")
	}
	return a
}

func (self *Parser) testProgram() bool {
	return self.isToken("def") || self.isToken("class") || self.testStatement()
}

func (self *Parser) params() ast.ASTree {
	list := make([]ast.ASTree, 0)
	if self.testParams() {
		list = append(list, ast.NewName(self.lexer.Read()))
	} else {
		panic("line " + self.lexer.GetLineNumber() + " -- not param " + self.lexer.Peek(0).GetText())
	}

	for self.isToken(",") {
		self.lexer.Read()
		if self.testParams() {
			list = append(list, ast.NewName(self.lexer.Read()))
		} else {
			panic("line " + self.lexer.GetLineNumber() + " -- not param " + self.lexer.Peek(0).GetText())
		}
	}

	return ast.NewParameterList(list)
}

func (self *Parser) testParams() bool {
	return self.isName(self.lexer.Peek(0))
}

func (self *Parser) paramList() ast.ASTree {
	self.readMatchToken("(")
	var a ast.ASTree = ast.NewParameterList(nil)
	if self.testParams() {
		a = self.params()
	}
	self.readMatchToken(")")
	return a
}

func (self *Parser) testParamList() bool {
	return self.isToken("(")
}

func (self *Parser) def() ast.ASTree {
	self.readMatchToken("def")
	list := make([]ast.ASTree, 3)
	if self.isName(self.lexer.Peek(0)) {
		list[0] = ast.NewName(self.lexer.Read())
		list[1] = self.paramList()
		list[2] = self.block()
	} else {
		panic("line " + self.lexer.GetLineNumber() + " -- not identifer " + self.lexer.Peek(0).GetText())
	}

	return ast.NewDefStmnt(list)
}

func (self *Parser) testDef() bool {
	return self.isToken("def")
}

func (self *Parser) args() ast.ASTree {
	list := make([]ast.ASTree, 1)
	list[0] = self.expr()
	for self.isToken(",") {
		self.lexer.Read()
		list = append(list, self.expr())
	}
	return ast.NewArguments(list)
}

func (self *Parser) testArgs() bool {
	return self.testExpr()
}

// postfix : "(" [ args ] ")" | "." IDENTIFIER | "[" expr "]"
func (self *Parser) postfix() ast.ASTree {
	if self.isToken("(") {
		self.lexer.Read()
		var a ast.ASTree = ast.NewArguments(nil)
		if self.testArgs() {
			a = self.args()
		}
		self.readMatchToken(")")
		return a
	} else if self.isToken(".") {
		self.lexer.Read()
		var a ast.ASTree
		if self.isName(self.lexer.Peek(0)) {
			a = ast.NewDot([]ast.ASTree{ast.NewName(self.lexer.Read())})
		} else {
			panic("line " + self.lexer.GetLineNumber() + " -- not obj member " +
				self.lexer.Peek(0).GetText())
		}
		return a
	} else if self.isToken("[") {
		self.lexer.Read()
		var a ast.ASTree = self.expr()
		self.readMatchToken("]")
		return ast.NewArrayRef([]ast.ASTree{a})
	} else {
		panic("line " + self.lexer.GetLineNumber() + " -- not postfix " + self.lexer.Peek(0).GetText())
	}
}

func (self *Parser) testPostfix() bool {
	return self.isToken("(") || self.isToken(".") || self.isToken("[")
}

func (self *Parser) member() ast.ASTree {
	if self.testDef() {
		return self.def()
	} else if self.testSimple() {
		return self.simple()
	} else {
		panic("line" + self.lexer.GetLineNumber() + " -- not class member " + self.lexer.Peek(0).GetText())
	}
}

func (self *Parser) testMember() bool {
	return self.testDef() || self.testSimple()
}

func (self *Parser) classBody() ast.ASTree {
	self.readMatchToken("{")
	list := make([]ast.ASTree, 0)
	if self.testMember() {
		list = append(list, self.member())
	}
	for self.isToken(";") || self.isToken(token.EOL) {
		self.lexer.Read()
		if self.testMember() {
			list = append(list, self.member())
		}
	}
	if len(list) == 0 {
		list = append(list, ast.NewNullStmnt())
	}
	self.readMatchToken("}")

	return ast.NewClassBody(list)
}

func (self *Parser) testClassBody() bool {
	return self.isToken("{")
}

func (self *Parser) defclass() ast.ASTree {
	self.readMatchToken("class")
	list := make([]ast.ASTree, 2, 3)
	if self.isName(self.lexer.Peek(0)) {
		list[0] = ast.NewName(self.lexer.Read())
	} else {
		panic("line" + self.lexer.GetLineNumber() + " -- class name error")
	}

	if self.isToken("extends") {
		self.lexer.Read()
		if self.isName(self.lexer.Peek(0)) {
			list = append(list, ast.NewName(self.lexer.Read()))
		} else {
			panic("line" + self.lexer.GetLineNumber() + " -- extends error")
		}
	}

	list[1] = self.classBody()
	return ast.NewClassStmnt(list)
}

func (self *Parser) testDefclass() bool {
	return self.isToken("class")
}

func (self *Parser) elements() ast.ASTree {
	list := make([]ast.ASTree, 1)
	list[0] = self.expr()
	for self.isToken(",") {
		self.lexer.Read()
		list = append(list, self.expr())
	}
	return ast.NewArrayLiteral(list)
}

func (self *Parser) testElements() bool {
	return self.testExpr()
}

// helper function

func (self *Parser) isToken(name string) bool{
	t := self.lexer.Peek(0)
	return t.IsIdentifier() && name == t.GetText()
}

func (self *Parser) readMatchToken(name string) {
	t := self.lexer.Read()
	if !(t.IsIdentifier() && name == t.GetText()) {
		panic("parser error at line " + self.lexer.GetLineNumber())
	}
}

func (self *Parser) isName(t token.Token) bool {
	if t.IsIdentifier() {
		_, ok := self.reserved[t.GetText()]
		if ok {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
