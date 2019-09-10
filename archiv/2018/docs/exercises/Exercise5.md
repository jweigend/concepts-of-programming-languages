# Exercise 5 - Parser Combinators in Go

## A parser for boolean expressions

Implement a parser for the following grammar:
```
  Variable := [a-zA-Z_][a-zA-Z0-9_]*
  Atom := Variable
        | "(" ^ Expression ^ ")"
  Not  := "!"* ^ Atom
  And  := Not ^ ("&" ^ And)?
  Or   := And ^ ("|" ^  Or)?
  Expression := Or
```
Ignore whitespaces!

Use `go get` to perform the following setup steps:
- Checkout the Parser Combinators: github.com/QAhell/Parser-Gombinators/parse
- Checkout the abstract syntax tree: github.com/jweigend/concepts-of-programming-languages/oop/ast
- Checkout the exercise template: github.com/jweigend/concepts-of-programming-languages/oop/parser/main.go and main_test.go

When your setup is complete, solve the exercise:
- Run `go test` on the command line and see that all tests fail.
- Implement all the methods in main.go until `go test` succeeds.
- Use the main method to perform experiments.

## Bonus exercise: Write a parser for JSON

- Go to http://json.org/ and read the grammar.
- Implement the grammar using parser combinators.
- Write tests that cover all aspects of JSON.

