package compiler

import (
	"monkey/code"
	"testing"
)

func TestArrayLiterals(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: "[]", expectedConstants: []interface{}{}, expectedInstructions: []code.Instructions{
				code.Make(code.OpArray, 0),
				code.Make(code.OpPop),
			},
		}, {
			input: "[1, 2, 3]", expectedConstants: []interface{}{1, 2, 3}, expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpArray, 3),
				code.Make(code.OpPop),
			}},

		{
			input: "[1 + 2, 3 - 4, 5 * 6]", expectedConstants: []interface{}{1, 2, 3, 4, 5, 6}, expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpSub),
				code.Make(code.OpConstant, 4),
				code.Make(code.OpConstant, 5),
				code.Make(code.OpMul),
				code.Make(code.OpArray, 3),
				code.Make(code.OpPop),
			}},
	}

	runCompilerTests(t, tests)
}
