# Boolean Parser (OOP)

## Purpose

This example implements a **boolean expression parser** using object-oriented design with an **Abstract Syntax Tree (AST)**.

## Key Concepts

- **Lexer/Parser/AST pattern** - Classic compiler design
- **Visitor pattern** - For AST traversal
- **Interfaces** - Abstract node types
- **Recursive evaluation** - Tree-based expression evaluation

## Components

| Directory | Responsibility |
|-----------|----------------|
| `lexer/` | Tokenizes input string |
| `parser/` | Builds AST from tokens |
| `ast/` | AST node types and evaluation |

## Grammar

```
Expression := Or
Or         := And ("|" And)*
And        := Not ("&" Not)*
Not        := "!" Atom | Atom
Atom       := "true" | "false" | "(" Expression ")"
```

## AST Nodes

- `Or` - Logical OR of children
- `And` - Logical AND of children
- `Not` - Logical negation
- `Val` - Boolean literal

## Run

```bash
go test -v ./...
```

## Test

Tests in each subdirectory verify lexer tokenization, parser construction, and AST evaluation.

```bash
go test ./ast/... ./lexer/... ./parser/...
```
