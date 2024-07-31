package evaluator

import (
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}