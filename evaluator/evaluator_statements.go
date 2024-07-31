package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range statements {
		result = Eval(stmt)
	}

	return result
}