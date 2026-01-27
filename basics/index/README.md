# Index

## Purpose

This example demonstrates how to build a **book index** using Go's built-in data structures. It maps words to the pages where they appear.

## Key Concepts

- **Slices** - `Page` is a slice of strings, `Book` is a slice of pages
- **Maps** - `Index` maps words (strings) to page numbers (int slices)
- **Range loops** - Iterating over slices and maps
- **Custom types** - Type aliases for better readability

## Data Structures

```go
type Page []string           // A page contains words
type Book []Page             // A book contains pages
type Index map[string][]int  // Index maps word -> page numbers
```

## Run

```bash
go test -v
```

## Test

The test in `index_test.go` creates a sample book with multiple pages and verifies that `MakeIndex()` correctly builds the word-to-page mapping.

```bash
go test ./...
```
