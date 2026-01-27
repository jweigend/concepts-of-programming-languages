# Flags

## Purpose

This example demonstrates how to use Go's built-in **`flag`** package to parse command-line arguments. This is the standard way to handle CLI options in Go programs.

## Key Concepts

- **`flag.String()`** - Define a string flag with a default value and description
- **`flag.Parse()`** - Parse the command-line arguments
- **`flag.PrintDefaults()`** - Print usage information for all flags

## Example Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-ip` | `192.168.1.1` | Override the default IP address |
| `-port` | `8080` | Override the default listen port |

## Run

```bash
# Show default values
go run main.go

# Override values
go run main.go -ip=10.0.0.1 -port=9090
```

## Test

This example has no tests. It is a runnable demonstration.
