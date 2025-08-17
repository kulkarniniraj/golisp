# GoLisp

A lisp interpreter written in Go, based on the [Building Your Own Lisp](https://www.buildyourownlisp.com/) book.

## File Structure

- evaluator.go: Evaluator for the lisp interpreter
- main.go: Main function for the lisp interpreter
- parser.go: Parser for the lisp interpreter
- lexer.go: Lexer for the lisp interpreter
- input_reader.go: Input reader for the lisp interpreter

## Features

- Uses [go-prompt](https://github.com/c-bata/go-prompt) for input reader
- Uses [logrus](https://github.com/sirupsen/logrus) for logging

## TODO

- [x] basic lexer
- [x] basic parser
- [x] basic evaluator
- [x] arithmetic operations
- [x] list operations, quote, list, head, tail, join
- [x] eval function
- [x] def for variable and function definition
- [x] lambda function, including variable arguments
- [x] conditional operations, if
- [ ] string operations
- [ ] standard library




