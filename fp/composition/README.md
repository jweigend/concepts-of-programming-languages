# Function Composition

## Purpose

This example demonstrates **function composition** in Go - combining simple functions to build more complex ones, a fundamental concept in functional programming.

## Key Concepts

- **Composition operator (g∘f)** - Apply f, then apply g to the result
- **Higher-order functions** - Functions that take/return functions
- **Generic composition** - Using `interface{}` for any type

## Mathematical Definition

$(g \circ f)(x) = g(f(x))$

## Example

```go
f := func(x int) int { return x * x }     // square
g := func(x int) int { return x + 1 }     // increment
gf := func(x int) int { return g(f(x)) }  // (g∘f)

gf(2)  // = g(f(2)) = g(4) = 5
```

## Generic Compose Function

```go
compose := func(g, f function) function {
    return func(x any) any {
        return g(f(x))
    }
}
```

## Run

```bash
go run main.go
```

## Test

This example has no tests. It is a runnable demonstration.
