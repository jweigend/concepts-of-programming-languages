# Exercise 4 - Building Parsers

During the course you learned about grammar, lexers and parsers.
The exercise is split up into three parts:

1) Implement a lexer to extract tokens from a boolean expression
2) Implement a parser to build an Abstract Syntax Tree and to evaluate the expression
3) Implement both the lexer and the parser using Antlr

If you do not finish during the lecture period, please try this at home.

## Exercise 4.1 - Lexer for boolean expressions

Write a lexer for boolean expressions inclusive braces.
Use the following grammar definition:

```bnf
<expression> ::= <term> { <or> <term> }
<term> ::= <factor> { <and> <factor> }
<factor> ::= <var> | <not> <factor> | (<expression>)
<or>  ::= '|'
<and> ::= '&'
<not> ::= '!'
<var> ::= '[a-zA-Z0-9]*'
```

_Note:_ A lexer only uses the lexer rules from the grammar.

The lexer has a method

```go
func splitTokens(input string) []string {
```

to split the expression into tokens and a further method

```go
func (l *Lexer) NextToken() string
```

that iterates over the tokens.

**Disclaimer:** Feel free you use your very own software design.

🤥 **Write tests! Otherwise it does not happened!** 🤥

## Exercise 4.2 - Parser

The parser is split up into two parts:

- Define and implement an Abstract Syntax Tree
- Token parsing and Abstract Syntax Tree building

_Reminder:_ Use the following grammar definition:

```bnf
<expression> ::= <term> { <or> <term> }
<term> ::= <factor> { <and> <factor> }
<factor> ::= <var> | <not> <factor> | (<expression>)
<or>  ::= '|'
<and> ::= '&'
<not> ::= '!'
<var> ::= '[a-zA-Z0-9]*'
```

### Exercise 4.2.1 - Abstract Syntax Tree (AST)

Write a program which builds an AST with nodes to evaluate logical expressions with (And, Not, Or with variables).

```text
Sample Expression: `A AND B OR C`

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

_Why named variables:_ This allows us to build the AST once and use it for multiple variable values.

Notes that might help:

- Interfaces and Polymorphism
- Nodes are different but what are the commonalities?
- Simply follow the rules

🤥 **Write tests! Otherwise it does not happened!** 🤥

Write a unit test which builds the AST and evaluate the expression with given boolean values for the variables A, B, C.

### Exercise 4.2.2 Recursive Descent Parser

Write a recursive descent parser. The parser must implement the grammar rules  (that was enough hint).

🤥 **Write tests! Otherwise it does not happened!** 🤥

### Exercise 4.3 Antlr

We now use Antlr to generate a lexer and a parser for a given grammar definition.

Follow the go Antlr quick-start: <https://github.com/antlr/antlr4/blob/master/doc/go-target.md>

You need to do the following things:

- Antlr Setup (see above)
- Define an Antlr grammar file (`boolparser.g4`)
- Generate lexer and parser source code
- Use the generated files to parse boolean expressions

Should be not to hard 🤙
