package vm

import "testing"
func TestIndexExpressions(t *testing.T) { tests := []vmTestCase{
	{"[1, 2, 3][1]", 2},
	{"[1, 2, 3][0 + 2]", 3},
	{"[[1, 1, 1]][0][0]", 1},
	{"[][0]", Null},
	{"[1, 2, 3][99]", Null},
	{"[1][-1]", Null},
	{"{1: 1, 2: 2}[1]", 1},
	{"{1: 1, 2: 2}[2]", 2},
	{"{1: 1}[0]", Null},
	{"{}[0]", Null},
}
runVmTests(t, tests)
}