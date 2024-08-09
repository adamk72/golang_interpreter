# Notes

I purchased [_Writing An Interpreter In Go_][1] by Thorsten Ball with the intent of re-learning the basics of coding. I'd spent so many years in the abstracted realm of UI and so many more years doing program management, I thought this would be a useful mind-expanding exercise.

At the time of this writing, I hadn't written a lexer or parser in over 20 years and never written a REPL, compiler, or virtual machine. Also, I'd never written a lick of Go code either (not that it was much of a barrier compared to writing something like this in pure C).

What I share here are my notes about the project in an attempt to get things clear in my head. They probably won't make sense unless you have access to the books (there are actually two projects; one for an interpreter and the other for a compiler). Part of this exercise is also to help me solidify my understanding of Go, so you will see comments about Go sprinkled throughout.

## Definitions 

- Lexing, lexical analysis: The process of representing source code as tokens that 1) reflect the parts of a the source code and 2) are easier to use and manipulate in future processes. Also: "tokenizer" or "scanner."
- Tokens: The output of lexing. Giving names to the various meaningful parts of the source code.

## Generalities about Go

It is very common in this project to leverage the struct methods in Go. Essentially, we're creating a series of structs that we then apply methods to. Thus, we start with a Lexer struct that is assigned memory space corresponding to the input and the Lexer methods are then applied to the contents of that input field. Other fields in a given struct may be used to store state information about the current instantiation of the struct.

For all practical purposes, in Go, a struct with methods acts like a basic class object. The methods apply side effects to the struct to generate results to be later used by other functions and other structs in the application. I think this is the most important take away for this project in Go: you have to keep in mind that data is being stored outside of a given function and in the especially recursive nature of this project, data is constantly being changed.

## Interpreter

The lexer and the parser are fairly trivial chunks of code. The main hurdles are just keeping track of character and or token positions. I'm sure these would become more complicated with more nuanced languages. 






## Compiler

[1]: https://interpreterbook.com/
