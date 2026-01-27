# Sorting

## Purpose

This example demonstrates how to implement **custom sorting** in Go using the `sort.Interface`. It shows how to sort a slice of structs by different fields.

## Key Concepts

- **`sort.Interface`** - Requires `Len()`, `Less()`, and `Swap()` methods
- **Type aliases for sorting** - Create different sorters for different fields
- **Struct tags** - JSON tags for serialization

## Data Structure

```go
type Planet struct {
    Name       string
    Aphelion   float64  // Farthest distance from Sun (million km)
    Perihelion float64  // Closest distance from Sun (million km)
    Axis       int64    // Semi-major axis (km)
    Radius     float64  // Planet radius
}
```

## Sorters

| Sorter | Sorts by |
|--------|----------|
| `AxisSorter` | Semi-major axis (distance from Sun) |
| `NameSorter` | Planet name (alphabetically) |

## Run

```bash
go run main.go
```

## Test

This example has no tests. It is a runnable demonstration.
