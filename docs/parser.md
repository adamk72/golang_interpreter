# How the Parser works

The parser is a struct that holds information about the token list. It is passed a lexer that has been given an input string and where the lexer pulls apart the source code character by character, the parser pulls the lexer results apart token by token.

The parser has a list of functions that it can apply to tokens based on their token type (determined by the lexer). 

The heart of the parser is a `for` loop called `ParseProgram` that goes through the tokenized source code and parses each statement. Statements come in three flavors (for the Monkey language):

1. Let statements, for assignment.
2. Return statements, for giving a final value in an expression.
3. Expressions statements, for evaluating values of an expression.

## Statements

Statements are a collection of tokens organized in anticipated ways (i.e., they have a syntax). As the parser finds valid statements they are added to an _Abstract Syntax Tree_ (AST) which is called the "program" inside of the project.

For instance, if `let` is found, then the next token needs to be of the `IDENT` token type; if not, it return `nil`.

So this is valid:
```
let x = 5;
```
whereas this is not:
```
let = foo;
```
since the equal sign is not of type `IDENT`.

# Recursive Descent Parser

For this parser, the author chose to write a recursive descent parser, also known as a "Pratt parser" based on the work of one Vaughan Pratt. 

# How the AST comes into play

The "program" that is generated is literally just a list of statement nodes (which are simply Go structs). 

How the AST is used by the parser is by simplifying the process of turning lexer token types into proper nodes for evaluation in the syntax tree.

So the process is lexer -> tokens; parser(tokens -> ast nodes) -> abstract syntax tree.

Going back to the 'let' example, the parser ges the `let` token and creates an ast `letStatement` node out of it. 

Then, if the next token is a valid `IDENT` type token, it adds that as a node the "program" list. The next token should be a type of `ASSIGN` and if so, then what follows must then be an expression.

This shows the recursive nature of the program. Start with a high level statement, then break it down a bit to confirm what type of statement it is (let, return, or expression), and then when another expression is encountered, start a new parsing processes again until a proper statement is generated and passed back up as a formal result.

# Parsing expressions

Everything that isn't a let or return statement is an expression statement, meaning it will express some value at some point in the parsing process. 

## Expression Recursion

Everything starts by finding a "left side" function (from the `prefixPraseFns` list that is registered with the parser). If it exists, then the parser keeps moving through each token until a certain criterion is met:

1. The next token is _not_ a semicolon and additionally, the "precedence" of the current token is less than the next token.
2. If for the next token, here's _not_ an function from the `infixParseFns` list to act on.

We'll start with the second criterion first.

### Infixes

If next token is an "infix", that is an operator such as '+' or '-' and the like, then there is more work to be done. The next token is chosen and passed to the function. 

Usually, these functions do exactly what you'd expect. Take the first (left) token and the next (right) token and further add then to the tree (The act of adding, for example, will happen later). Of course, the right token could lead to being another expression, in which case the recursion continues.

#### Left parens and brackets as infixes
What's interesting is that the left parens and brackets (`(` and `[`) can act as an "infix" of source, such as for call expressions and index expressions. 

Normally, the left parens and brackets are "prefixes," being used to indicate the start of grouped expressions (e.g., parameter lists) and array literals.

This demonstrates how one needs to look at language holistically when creating the logic behind the parser.

## Precedence & parseExpression

Precedence is the "magic" of the parser and is based on the work of the previously mentioned Vaughan Pratt and his 1973 paper, _Top Down Operator Precedence_ of which Ball goes into great detail in the _Writing an Interpreter_ book. I will just skim the idea here, based on the code.

We're all familiar with mathematical precedence, where `5 + 2 * 3 = 11` based on the fact that `2 * 3 = 6` occurs first, then the `5` is added to result in `11`. Parentheses are used to be explicit; `5 + (2 * 6) = 11` whereas `(5 + 2) * 3 = 21` because now the precedence on the plus operator has changed.

Something similar is happening in the parser. Every expression parse occurs in the context of a precedence value. This happens with `parseExpression` and the very first time it's called, it gets passed a `LOWEST` precedence value.

Recall that the parser has a list of infix function (only three in total); the chosen function runs on the next token which may have a higher precedence the previous (for example,the product operator is higher than the sum operator).

Each of these function recursively call `parseExpression` until such a time that the precedence condition is triggered (based on a map, where math precedences are as normal, then followed by prefixes, calls, and finally array indexes as the highest).

As the recursion unwinds, each function is added to the AST, building up a tree of different types of expressions structures.

Let's adjust our previous math expression to be a proper Monkey expression: `5 + 2 * 3 == 11;`; parsing it would go something like this:

1. Enter the initial `for` loop. We're actually more interested in the _next_ token, not the current token at any give time. The first token will be resolved last, in this case the `5`. Numbers are considered "prefixes" for the sake of the process and an integer literal is created for it. We continue into the loop (recursion level: 0).
2. Inside the loop check the the next token, a `+`; it has an infix function that is called and adds the `+` token to the tree and then calls `parseExpression` on the next token, the `2`. When it does that, the precedence level of the `+` (4)  is passed to `parseExpression` (recursion level: 1). 
3. `2` is an integer literal, so goes through same process as the `5`. It gets assigned to an integer literal and the next token is called `*` (recursion level: 2).
4. `*` has a level of 5 which gets passed to the `parseExpression` for the `3`  (recursion level: 3). 
5. The following `==` has a p-level of only 2, and passes the 11 on to the next level  (recursion level: 4).
6. Finally, the next token is semicolon and now the p-level is 4. When the check for the semicolon returns, it's gets the default, `LOWEST` of 0. The loop condition fails and the recursion starts to unwind.
7. As each `parseExpression` is exited, the AST node that was initially created is "completed." In the case this expression, the `11` is set as the right hand part of the `==` operator, whose left side happens to be an expression, `2 * 3`.
8. Recall that we still have to unwind inside of a for loop, so there is the possibility that we will continue on the with more prefix additions to the AST nodes. 

Of course, the whole thing about parsing over that expressions and putting into the AST is to evaluate it. In the Monkey language, the expression `5 + 2 * 3 == 11;` will evaluate to `true` where as `5 + 2 * 3;` would evaluate to `11`. 