package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func evalIfExpression(ie *ast.IfExpression) object.Object { condition := Eval(ie.Condition)
	if isTruthy(condition) {
	return Eval(ie.Consequence)
	} else if ie.Alternative != nil { return Eval(ie.Alternative)
	} else {
	return NULL
	} }
	