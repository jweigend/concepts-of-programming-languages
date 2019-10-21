# Exercise 4 - Building Parsers

If you do not finish during the lecture period, please finish it as homework.

## Exercise 4.1 - AST - Abstract Syntax Tree
Write a programm which builds an AST with nodes to evaluate logical expressions with (And, Not, Or with variables)

```
Sample Expression: A AND B OR C

             ----------
             |   OR   |
             ----------
            /          \
        ---------      ----------
        |  AND  |      |  Var:C |
        ---------      ----------
        /       \
  ---------   --------- 
  | Var:A |   | Var:B |
  ---------   ---------
```

The tree should be evaluated with a evaluation methods which supports named variables:

```go
eval(vars map[string]bool) bool
```

Write a unit test which builds the AST and evaluate the expression with given boolean values for the variables A, B, C.

## Exercise 4.2 - Lexer
Write a lexer for boolean expressions inclusive braces.

## Exercise 4.2 - Parser
Write a recursive descent parser and test the following expressions in a unit test:

1) a & b | c

Test all 8 combinations for a, b and c (true or false) in a unit test:

2) lether & (blue-metallic | green-metallic) | premium
