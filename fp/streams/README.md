# Streams

## Purpose

This example implements a **Java 8 Streams-like API** for functional data processing in Go using **Go 1.18+ Generics**.

## Key Concepts

- **Type-safe generics** - No more `interface{}` and type assertions
- **Map/Filter/Reduce** - Core functional operations
- **Method chaining** - Fluent API style
- **Type-transforming operations** - `MapTo()`, `FlatMap()`, `ReduceTo()` for cross-type transformations
- **Iterator pattern** - Underlying iteration abstraction

## API

### Basic Usage (same-type operations)

```go
// Type-safe stream - no type assertions needed!
result := Of([]string{"a", "b", "c", "1", "D"}).
    Map(strings.ToUpper).
    Filter(notDigit).
    Reduce(concat)  // result is string, not interface{}
```

### Type-Transforming Operations

```go
// Transform int stream to string stream
stringStream := MapTo(Of(numbers), func(n int) string {
    return fmt.Sprintf("num:%d", n)
})

// Reduce to different type
wordCounts := ReduceTo(Of(words), make(map[string]int), 
    func(counts map[string]int, word string) map[string]int {
        counts[word]++
        return counts
    })

// Flatten nested structures
chars := FlatMap(Of(words), func(word string) []rune {
    return []rune(word)
})
```

## Operations

| Operation | Type | Description |
|-----------|------|-------------|
| `Of(slice)` | Constructor | Create stream from slice |
| `Map(f)` | Intermediate | Transform each element (same type) |
| `Filter(p)` | Intermediate | Keep elements matching predicate |
| `Reduce(a)` | Terminal | Combine all elements into one |
| `Collect()` | Terminal | Return elements as slice |
| `MapTo(s, f)` | Transform | Map to different type (standalone function) |
| `FlatMap(s, f)` | Transform | Map and flatten (standalone function) |
| `ReduceTo(s, init, f)` | Transform | Reduce to different type (standalone function) |

## Types

```go
type Predicate[T any] func(T) bool
type Mapper[T any] func(T) T
type TransformMapper[T, R any] func(T) R
type Accumulator[T any] func(T, T) T
type Pair[K, V any] struct { Key K; Value V }
```

## Why Standalone Functions for Type Transformation?

Go generics don't support method-level type parameters. Therefore, operations that change the element type (like `MapTo`) must be standalone functions rather than methods on `Stream[T]`.

## Run

```bash
go test -v
```

## Test

The test in `streams_test.go` demonstrates:
- Basic Map/Filter/Reduce with type safety
- WordCount using `ReduceTo`
- Type transformation with `MapTo`
- Flattening with `FlatMap`
- Iterator usage
- Collecting to slice

```bash
go test ./...
```
