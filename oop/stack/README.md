# Stack

## Purpose

This example implements a **generic LIFO (Last-In-First-Out) stack** using Go 1.18+ generics for type safety.

## Key Concepts

- **Go Generics** - Type parameter `[T any]` for compile-time type safety
- **Error handling** - Returns `(T, error)` instead of panic (idiomatic Go)
- **LIFO data structure** - Last pushed item is first popped
- **Slice-based implementation** - Dynamic sizing

## API

```go
s := NewStack[string]()      // Type-safe string stack
s.Push("value")              // Add to top
value, err := s.Pop()        // Remove and return top (with error handling)
size := s.Size()             // Number of elements
item := s.Get(idx)           // Access by index
```

## Modern Go (1.18+) vs Legacy

```go
// Modern (type-safe)
s := NewStack[int]()
s.Push(42)
val, err := s.Pop()  // val is int, no casting needed

// Legacy (interface{})
s := NewStack()
s.Push(42)
val := s.Pop().(int)  // Runtime casting, can panic
```

## Run

```bash
go test -v
```

## Test

The test in `stack_test.go` verifies push, pop, size, and empty stack behavior.

```bash
go test ./...
```
