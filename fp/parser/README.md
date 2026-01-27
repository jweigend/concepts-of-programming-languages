# Parser (Functional)

## Purpose

This example implements a **parser combinator library** for parsing boolean expressions using functional programming techniques.

## Key Concepts

- **Parser combinators** - Small parsers combined into larger ones
- **Recursive descent parsing** - Top-down parsing approach
- **Higher-order functions** - Parsers that return parsers
- **Immutable input** - Functional approach to parsing

## Grammar

```
Expression := Or Spaces*
Or         := And ("|" Or)?
And        := Not ("&" And)?
Not        := "!" Not | Atom
Atom       := "true" | "false" | "(" Expression ")"
```

## Combinators

| Combinator | Purpose |
|------------|---------|
| `AndThen` | Sequence two parsers |
| `OrElse` | Try first parser, fall back to second |
| `Optional` | Zero or one occurrence |
| `Convert` | Transform parse result |

## Run

```bash
go test -v
```

## Test

The test in `boolparser_test.go` verifies parsing of various boolean expressions like `"true & false"`, `"!true | false"`, etc.

```bash
go test ./...
```
