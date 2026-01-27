# Types

## Purpose

This directory demonstrates Go's **type system**, including basic types, type aliases, structs, pointers, and functions as first-class citizens.

## Subdirectories

| Directory | Description |
|-----------|-------------|
| `list/` | Linked list implementation |
| `slices/` | Slice operations and tests |
| `strings/` | String manipulation functions |

## Key Concepts in `simpletypes.go`

- **Type aliases** - `type Integer = int`
- **Struct types** - Custom composite types
- **Pointers** - `*int`, `**int` (double pointers)
- **Functions as values** - Passing functions as parameters
- **iota** - Auto-incrementing constants

## Run

```bash
go run simpletypes.go
```

## Test

```bash
# Test slices
go test ./slices/...

# Test strings
go test ./strings/...
```
