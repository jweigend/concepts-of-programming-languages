# Swap

## Purpose

This example demonstrates **different ways to swap variable values** in Go, illustrating the difference between value semantics and pointer semantics.

## Key Concepts

- **Multiple return values** - Go's idiomatic way to swap
- **Value semantics** - Parameters are copied (swap doesn't work)
- **Pointer semantics** - Pass by reference using pointers
- **Double pointers** - Swapping pointer values themselves

## Swap Implementations

| Function | Works? | Why |
|----------|--------|-----|
| `a, b = b, a` | ✅ | Go's built-in tuple assignment |
| `swap0(a, b)` | ✅ | Returns swapped values |
| `swap1(x, y int)` | ❌ | Values are copied, original unchanged |
| `swap2(x, y *int)` | ✅ | Swaps values via pointers |
| `swap3(x, y **int)` | ✅ | Swaps the pointer addresses |

## Run

```bash
go run main.go
```

## Test

This example has no tests. It is a runnable demonstration.
