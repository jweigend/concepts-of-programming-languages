# Mail

## Purpose

This example demonstrates **interface-based design** and **dependency injection** in Go through a mail sending abstraction.

## Key Concepts

- **Interface abstraction** - `Sender` interface hides implementation
- **Dependency injection** - Client receives sender via constructor
- **Multiple implementations** - SMTP, mock, etc.
- **Separation of concerns** - Client doesn't know transport details

## Components

| Directory | Responsibility |
|-----------|----------------|
| `mail.go` | Core interfaces and types |
| `client/` | Mail client using Sender interface |
| `smtp/` | SMTP implementation |
| `util/` | Helper utilities |

## Interface

```go
type Sender interface {
    SendMail(address Address, message string)
}
```

## Run

```bash
go test -v ./client/...
```

## Test

The test in `client/` uses a mock sender to verify the client logic without actual network calls.

```bash
go test ./client/...
```
