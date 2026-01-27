# Exceptions

## ⚠️ WARNING: This is an Anti-Pattern in Go!

This example demonstrates how try-catch-finally **could** be implemented in Go using `panic` and `recover`. However, **you should NOT use this pattern in production Go code**.

## Why This is an Anti-Pattern

Go deliberately chose **explicit error handling** over exceptions:

```go
// ✅ Idiomatic Go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("failed: %w", err)
}
defer cleanup()

// ❌ Anti-Pattern (this example)
TryCatch(
    func() { Throw("error") },
    func(ex Exception) { /* catch */ },
    func() { /* finally */ },
)
```

## Educational Purpose

This example is included to show:
1. How `panic`/`recover` works internally
2. Why Go chose explicit error handling over exceptions
3. How other languages implement exception handling

## Key Concepts

- **`panic`** - Stops normal execution (like `throw`)
- **`recover`** - Catches a panic in a deferred function
- **`defer`** - Ensures cleanup runs even after panic

## Run

```bash
go run exceptions.go
```

## When to Use panic/recover

- **panic**: Only for truly unrecoverable errors (programming bugs)
- **recover**: Only at package boundaries to prevent crashes
