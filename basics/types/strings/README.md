# Strings

## Purpose

This example demonstrates **string manipulation** in Go, including the important distinction between bytes and runes for UTF-8 handling.

## Key Concepts

- **Strings are immutable** - Operations create new strings
- **Runes** - Unicode code points for proper UTF-8 handling
- **String reversal** - Working with multi-byte characters

## Key Function

```go
func Reverse(s string) string
```

Reverses a string properly handling UTF-8 multi-byte characters.

## Run

```bash
go test -v
```

## Test

The test file `strings_test.go` verifies string operations including reversal of ASCII and Unicode strings.

```bash
go test ./...
```
