package vm

import "testing"

func TestCallingFunctionsWithoutArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
						 let fivePlusTen = fn() { 5 + 10; };
						 fivePlusTen();
						 `,
			expected: 15,
		},
		{
			input: `
							 let one = fn() { 1; };
							 let two = fn() { 2; };
							 one() + two()
							 `,
			expected: 3,
		},
		{
			input: `
							 let a = fn() { 1 };
							 let b = fn() { a() + 1 };
							 let c = fn() { b() + 1 };
							 c();
							 `,
			expected: 3,
		},
	}
	runVmTests(t, tests)
}

func TestFunctionsWithReturnStatement(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
		let earlyExit = fn() { return 99; 100; };
		earlyExit();
		`,
			expected: 99,
		},
		{
			input: `
		let earlyExit = fn() { return 99; return 100; };
		earlyExit();
		`,
			expected: 99,
		},
	}
	runVmTests(t, tests)
}

func TestFunctionsWithoutReturnValue(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
						 let noReturn = fn() { };
						 noReturn();
						 `,
			expected: Null,
		},
		{
			input: `
						 let noReturn = fn() { };
						 let noReturnTwo = fn() { noReturn(); };
						 noReturn();
						 noReturnTwo();
						 `,

			expected: Null,
		},
	}
	runVmTests(t, tests)
}

func TestFirstClassFunctions(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
						 let returnsOne = fn() { 1; };
						 let returnsOneReturner = fn() { returnsOne; };
						 returnsOneReturner()();
						 `,
			expected: 1,
		},
	}
	runVmTests(t, tests)
}
