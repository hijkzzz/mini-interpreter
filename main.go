package main

import (
	"os"
	"stone/lexer"
	"stone/environment"
	"stone/token"
	"stone/parser"
	"fmt"
)

func usage() {
	fmt.Println("[Usage] stone test")
}

func main() {
	defer func(){
		if err := recover(); err!=nil{
			fmt.Println(err)
		}
	}()

	if os.Args == nil || len(os.Args) != 2 {
		usage()
		return
	}

	fin, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	LEXER := lexer.NewLexer(fin)
	PARSER := parser.NewParser(LEXER)
	ENV := environment.NewNativeEnv()

	for LEXER.Peek(0) != token.EOF {
		stmnt := PARSER.Parse()
		stmnt.Eval(ENV)
	}
}
