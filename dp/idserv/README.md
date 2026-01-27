# ID Service

## Purpose

This example implements a **distributed UUID service** with a clean separation between API, core implementation, and remote access (proxy/stub pattern).

## Key Concepts

- **Interface-based design** - `IDService` interface
- **Proxy pattern** - Remote access abstraction
- **Stub pattern** - Server-side request handling
- **Network transparency** - Same interface locally and remotely

## Architecture

```
client/ ──→ proxy/ ──→ [network] ──→ stub/ ──→ core/
         (client)                  (server)
```

## Components

| Directory | Responsibility |
|-----------|----------------|
| `idserv-api.go` | Interface definition |
| `core/` | Local implementation |
| `client/` | Example client code |
| `remote/proxy/` | Client-side network proxy |
| `remote/stub/` | Server-side request handler |
| `remote/server/` | Network server |

## Interface

```go
type IDService interface {
    NewUUID(clientID string) string
}
```

## Run

```bash
make  # See Makefile for targets
```

## Test

This example has no automated tests.
