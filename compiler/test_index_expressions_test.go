package compiler

import (
	"monkey/code"
	"testing"
)

func TestIndexExpressions(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: "[1, 2, 3][1 + 1]", expectedConstants: []interface{}{1, 2, 3, 1, 1}, expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpArray, 3),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpConstant, 4),
				code.Make(code.OpAdd),
				code.Make(code.OpIndex),
				code.Make(code.OpPop),
			}},
		{
			input: "{1: 2}[2 - 1]", expectedConstants: []interface{}{1, 2, 2, 1}, expectedInstructions: []code.Instructions{code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpHash, 2),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpSub),
				code.Make(code.OpIndex),
				code.Make(code.OpPop),
			}},
	}

	runCompilerTests(t, tests)
}
