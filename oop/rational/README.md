# Rational Numbers

## Purpose

This example implements **rational number arithmetic** in Go, demonstrating how to create a custom numeric type with proper mathematical operations.

## Key Concepts

- **Value types** - Rational is a struct passed by value
- **Constructor function** - `NewRational()` ensures valid state
- **Automatic simplification** - GCD reduction on creation
- **Method receivers** - Operations as methods

## API

```go
r1 := NewRational(1, 2)  // 1/2
r2 := NewRational(2, 3)  // 2/3

sum := r1.Add(r2)        // 1/2 + 2/3 = 7/6
product := r1.Multiply(r2)  // 1/2 * 2/3 = 1/3
```

## Implementation Details

- Division by zero panics in constructor
- Results are automatically simplified (e.g., 4/8 → 1/2)
- Uses Euclidean algorithm for GCD

## Run

```bash
go test -v
```

## Test

The test in `rational_test.go` verifies arithmetic operations and automatic simplification.

```bash
go test ./...
```
