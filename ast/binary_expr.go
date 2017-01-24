package ast

import (
	"reflect"
	"stone/environment"
	"strconv"
)

/*
	抽象语法树——双目表达式
 */

type BinaryExpr struct {
	astList
}

func NewBinaryExpr(list []ASTree) *BinaryExpr {
	return &BinaryExpr{astList{list}}
}

func (self *BinaryExpr) Left() ASTree{
	return self.Child(0)
}

func (self *BinaryExpr) Operator() string {
	op, ok := self.Child(1).(*OP)
	if ok {
		return op.token.GetText()
	} else {
		panic("bad operator")
	}
}

func (self *BinaryExpr) Right() ASTree {
	return self.Child(2)
}

func (self *BinaryExpr) Eval(env environment.Environment, args... interface{}) interface{} {
	op := self.Operator()
	if op == "="  {
		right := self.Right().Eval(env)
		return self.computeAssign(env, right)
	} else {
		left := self.Left().Eval(env)
		right := self.Right().Eval(env)
		return self.computeOp(left, op, right)
	}
}

func (self *BinaryExpr) computeAssign(env environment.Environment, right interface{}) interface{} {
	// = 左边必须是变量名
	l, ok := self.Left().(*Name)
	if ok {
		env.Set(l.Name(), right)
		return right
	} else {
		panic("bad assignment")
	}
}

func (self *BinaryExpr) computeOp(left interface{}, op string, right interface{}) interface{} {
	leftKind := reflect.TypeOf(left).Kind()
	rightKind := reflect.TypeOf(right).Kind()

	if leftKind == reflect.Int && rightKind == reflect.Int {
		return self.computeNumber(left.(int), op, right.(int))
	} else {
		if op == "+" {
			if leftKind == reflect.String && rightKind == reflect.String {
				return left.(string) + right.(string)
			} else if leftKind == reflect.String && rightKind == reflect.Int {
				return left.(string) + strconv.Itoa(right.(int))
			} else if leftKind == reflect.Int && rightKind == reflect.String {
				return strconv.Itoa(left.(int)) + right.(string)
			} else {
				panic("bad +")
			}
		} else if op == "==" {
			if left == right  {
				return 1
			} else {
				return 0
			}
		} else {
			panic("bad type")
		}
	}
}

func (self *BinaryExpr) computeNumber(left int, op string, right int) int {
	switch op {
	case "+": return left + right
	case "-": return left - right
	case "*": return left * right
	case "/": return left / right
	case "%": return left % right
	case "==": if left == right {
		return 1
	} else {
		return 0
	}
	case ">": if left > right {
		return 1
	} else {
		return 0
	}
	case "<": if left < right {
		return 1
	} else {
		return 0
	}
	default: panic("bad operator")
	}
}
