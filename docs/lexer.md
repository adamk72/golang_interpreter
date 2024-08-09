# How the Lexer works

The lexer is a character scanner (for ASCII characters, that is) that exams each char and tries to figure out if 1) it is a single character token in its own right (such as an operator) or 2) the start of a bigger token, like an identifier or keyword.

In the project, this is done through a switch statement. Some tokens are simply created immediately (most of the math operators) where they are assigned a token type for later tracking as well as the actual token itself.

Others have checks; the `peekChar()` function, for example, looks at the next char to see if a meaningful token is created (like the `!=` token). 

For identifiers, a lookup table is employed to determine is a word is a keyword or not.

# How the Lexer is used

A `New` lexer is passed to the Parser which, through a wrapper function, uses the lexer's `NextToken` function. 

ALl the lexer does is provide a list of tokens; it's up the parser to make more sense of the source code.

# Go methods, receivers, and character positions

The lexer tracks the position of the characters it has been reading through fields of the lexer struct, `position` and `readPosition`. 

In Go, one can access the input (of type `string`) directly through index which results in Go code like this: 

```go
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
```

I highlight this to demonstrate how the lexer struct/object keeps track of things through its fields. "l" is the lexer "[receiver][1]" (), and `l.readPosition` is the mutable integer type of the current position within `l.input`. When the lexer is instantiated (through a simple, `l := &Lexer{input: input}`), the receiver method is then allowed to act on the struct members.

[1]: https://gobyexample.com/methods

