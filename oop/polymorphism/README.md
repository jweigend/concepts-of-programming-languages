# Polymorphism

## Purpose

This example demonstrates **polymorphism in Go** using interfaces and struct embedding. It shows how Go achieves polymorphic behavior without traditional inheritance.

## Key Concepts

- **Interfaces for polymorphism** - `fmt.Stringer` interface
- **Struct embedding** - Composition over inheritance
- **No type hierarchy** - Types are independent
- **Duck typing** - If it implements the interface, it's compatible

## Example

```go
type Point struct{ x, y int }
type ColorPoint struct {
    Point  // Embedding (delegation, not inheritance)
    c int
}

// Both implement fmt.Stringer
var s fmt.Stringer
s = Point{1, 2}       // Works
s = ColorPoint{...}   // Works
```

## Important Differences from Java/C++

| Go | Java/C++ |
|----|----------|
| `p = cp` ❌ | Works (subtype assignment) |
| `p = cp.Point` ✅ | Explicit field access |
| Interface polymorphism ✅ | Class hierarchy polymorphism |

## Run

```bash
go run polymorphism.go
```

## Test

This example has no tests. It is a runnable demonstration.
