# Evaluator

The evaluator is where the rubber meets the road in the REPL. Like with many other processes in this project, it makes heavy use of recursion.

After the parser was done, we ended up with an abstract syntax tree that defined a set of expression nodes and their relative relation to one another. In an expression like `5 + 2 * 3 == 11;`, we've created a lot of AST nodes of various type: integral literals and infix expressions in this particular case.

We have:

0. The program itself.
1. An infix expression of:
- Left: IntegerLiteral: 2
- Operator: *
- Right: IntegerLiteral: 3
2. Another infix expression of:
- Left: IntegerLiteral: 5
- Operator: +
- Right: InfixExpression: 2 * 3 
3. Another infix expression of: 
- Left: InfixExpression: 5 + 2 * 3 
- Operator: = 
- Right: IntegerLiteral: 11

Now the evaluator needs to go through that tree, staring at the top and determining the values: 2 * 3 => 6; 5 + 6 => 11; 11 == 11 => true, which is the final output.

The first thing that is evaluated is the program itself, where all of the statements are collected. In a for loop, each statement (each an AST) is sent through eval process, distinguished by type.

In turn, each part of statement is also evaluated, until the process can return back with a proper value like an integer, string, or boolean. In our example, case, the first infix is evaluated; upon seeing both the left and right are integers, then the `*` operator is applied to generated an integer-type object with the value 6.

So on and so forth until the tree is completed and a final result comes out: `true` in the case of our expression.

# Conditionals and other non-trivial nodes

The evaluator function that are called for handling things like `if` statements are more complicated, of course, but not by too much. Since the evaluator itself is written in Go, simply leverage the same sorts of checks and continue evaluating further down the tree as needed, recursively. 


