# What is a REPL?

A REPL is very common way to essentially have command line-type access to a language. Standing for "Read, Eval, Print, Loop", one can type in a line of code into a terminal interface which will then:

1. Read the line of code.
2. Evaluate if the line of code is valid, and if so, what the result should be.
3. Print the result as output.
4. Look back around to receive another line of input.


# REPL in Go

The REPL for this project is simply a package that calls Go's I/O functions to take and and output the necessary content. 

Each line is passed first to a lexer, which is then added to parser that then passed the resulting program to the complier. 

The compiler and VM play the biggest parts of the REPL. Assuming no errors are found by the compiler, it generates bytecode that a new virtual machine than can run; this is the evaluation portion of the REPL. If all went well, the VM will return the last evaluated element. 

Now, in terms of the Interpreter project, much of what was noted with the compiler and VM are not needed; there an evaluator object that acts as a proxy for determining the output. The point is that the REPL is an evolutionary part of the project, being made more complex as the type of code evaluation and requirements change.
