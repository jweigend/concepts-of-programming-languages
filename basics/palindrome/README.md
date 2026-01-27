# Palindrome

## Purpose

This example demonstrates **three different implementations** of a palindrome checker, highlighting the difference between byte-based and rune-based string handling in Go.

## Key Concepts

- **ASCII vs UTF-8** - Single-byte vs multi-byte character handling
- **Runes** - Go's representation of Unicode code points
- **String reversal** - Comparing a string with its reverse

## Implementations

| Function | Approach | UTF-8 Support |
|----------|----------|---------------|
| `IsPalindrome()` | Byte comparison | ASCII only |
| `IsPalindrome2()` | Rune comparison | Full UTF-8 |
| `IsPalindrome3()` | String reversal | Full UTF-8 |

## Example

```go
IsPalindrome("anna")  // true
IsPalindrome2("αννα") // true (Greek letters)
```

## Run

```bash
go test -v
```

## Test

The test in `palindrome_test.go` verifies all three implementations with various ASCII and Unicode test cases.

```bash
go test ./...
```
