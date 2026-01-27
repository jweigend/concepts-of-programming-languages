# Lambda Calculus

## Purpose

This example implements **Lambda Calculus** concepts in Go, demonstrating how computation can be expressed purely through functions. Based on Graham Hutton's video lecture.

## Key Concepts

- **Church Numerals** - Numbers represented as functions
- **Lambda expressions** - Pure function definitions
- **SUCCESSOR function** - Building numbers from other numbers
- **Self-referential function type** - `type fnf func(fnf) fnf`

## Church Numerals

Numbers are represented by how many times they apply a function:

| Number | Definition | Meaning |
|--------|------------|---------|
| ZERO | λf.λx.x | Apply f zero times |
| ONCE | λf.λx.f(x) | Apply f once |
| TWICE | λf.λx.f(f(x)) | Apply f twice |
| THRICE | λf.λx.f(f(f(x))) | Apply f three times |

## SUCCESSOR

`SUCCESSOR(n)` returns `n+1`:

```go
SUCCESSOR(TWICE) → THRICE
```

## Run

```bash
go run main.go
```

The output uses dots to visualize the numeric value (one dot per application).

## Test

This example has no tests. It is a runnable demonstration.
