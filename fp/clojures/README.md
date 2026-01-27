# Closures

## Purpose

This example demonstrates **closures** in Go - functions that capture and retain access to variables from their enclosing scope.

## Key Concepts

- **Closure** - A function that "closes over" variables from its environment
- **State preservation** - Captured variables persist between calls
- **Function factories** - Functions that return functions

## Example

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

nextInt := intSeq()
fmt.Println(nextInt())  // 1
fmt.Println(nextInt())  // 2
fmt.Println(nextInt())  // 3
```

## How It Works

The inner function captures `i` from `intSeq()`. Each call to `nextInt()` accesses and modifies the same `i` variable.

## Run

```bash
go run main.go
```

## Test

This example has no tests. It is a runnable demonstration.
