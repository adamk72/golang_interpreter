package vm

import (
	"monkey/compiler"
	"testing"
)

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
		{
			input: `
			let returnsOneReturner = fn() {
			let returnsOne = fn() { 1; };
			returnsOne;
			};
			returnsOneReturner()();
			`,
			expected: 1,
		},
	}
	runVmTests(t, tests)
}
func TestCallingFunctionsWithBindings(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let one = fn() { let one = 1; one };
			one();
			`,
			expected: 1,
		},
		{
			input: `
			let oneAndTwo = fn() { let one = 1; let two = 2; one + two; };
			oneAndTwo();
			`,
			expected: 3,
		},
		{
			input: `
			let oneAndTwo = fn() { let one = 1; let two = 2; one + two; };
			let threeAndFour = fn() { let three = 3; let four = 4; three + four; };
			oneAndTwo() + threeAndFour();
			`,
			expected: 10,
		},
		{
			input: `
			let firstFoobar = fn() { let foobar = 50; foobar; };
			let secondFoobar = fn() { let foobar = 100; foobar; };
			firstFoobar() + secondFoobar();
			`,
			expected: 150,
		},
		{
			input: `
			let globalSeed = 50;
			let minusOne = fn() {
			let num = 1;
			globalSeed - num;
			}
			let minusTwo = fn() {
			let num = 2;
			globalSeed - num;
			}
			minusOne() + minusTwo();
			`,
			expected: 97,
		},
	}
	runVmTests(t, tests)
}
func TestCallingFunctionsWithArgumentsAndBindings(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
			let identity = fn(a) { a; };
			identity(4);
			`,
			expected: 4,
		},
		{
			input: `
			let sum = fn(a, b) { a + b; };
			sum(1, 2);
			`,
			expected: 3,
		},
		{
			input: `
			let sum = fn(a, b) {
			let c = a + b;
			c; };
			sum(1, 2);
			`,
			expected: 3,
		}, {
			input: `
			let sum = fn(a, b) {
			let c = a + b;
			c; };
			sum(1, 2) + sum(3, 4);`,
			expected: 10,
		}, {
			input: `
			let sum = fn(a, b) {
			let c = a + b;
			c; };
			let outer = fn() {
			sum(1, 2) + sum(3, 4);
			};
			outer();
			`,
			expected: 10,
		},
		{
			input: `
			let globalNum = 10;
			let sum = fn(a, b) {
			let c = a + b;
			c + globalNum;
			};
			let outer = fn() {
			sum(1, 2) + sum(3, 4) + globalNum;
			};
			outer() + globalNum;
			`,
			expected: 50,
		},
	}
	runVmTests(t, tests)
}

func TestCallingFunctionsWithWrongArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			input:    `fn() { 1; }(1);`,
			expected: `wrong number of arguments: want=0, got=1`,
		},
		{
			input:    `fn(a) { a; }();`,
			expected: `wrong number of arguments: want=1, got=0`,
		},
		{
			input:    `fn(a, b) { a + b; }(1);`,
			expected: `wrong number of arguments: want=2, got=1`,
		},
	}

	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err = vm.Run()
		if err == nil {
			t.Fatalf("expected VM error but resulted in none.")
		}

		if err.Error() != tt.expected {
			t.Fatalf("wrong VM error: want=%q, got=%q", tt.expected, err)
		}
	}
}
