# Versions and Dependencies

This document provides an overview of all versions and dependencies used in this project.

**Last updated:** January 26, 2026

---

## Go Version

| Component | Version |
|-----------|---------|
| Go | 1.24.0 |
| Platform | linux/amd64 |

---

## Direct Dependencies

| Package | Version | Description |
|---------|---------|-------------|
| github.com/coreos/etcd | v3.3.27+incompatible | Distributed key-value store |
| github.com/golang/protobuf | v1.5.4 | Protocol Buffers for Go |
| golang.org/x/sys | v0.40.0 | System-level packages for Go |
| google.golang.org/grpc | v1.78.0 | gRPC framework for Go |

---

## Indirect Dependencies

| Package | Version | Description |
|---------|---------|-------------|
| github.com/gogo/protobuf | v1.3.2 | Extended Protocol Buffers |
| golang.org/x/net | v0.49.0 | Network packages |
| golang.org/x/text | v0.33.0 | Text processing (Unicode, etc.) |
| google.golang.org/genproto/googleapis/rpc | v0.0.0-20260126211449-d11affda4bed | Google API proto definitions |
| google.golang.org/protobuf | v1.36.11 | Protocol Buffers (new API) |

---

## Project Structure

| Directory | Description |
|-----------|-------------|
| `basics/` | Go programming basics |
| `cp/` | Concurrent Programming (Channels, Locks) |
| `dp/` | Design Patterns |
| `fp/` | Functional Programming |
| `oop/` | Object-Oriented Programming |
| `raft/` | Raft Consensus Algorithm |
| `sp/` | Systems Programming |
| `docs/` | Documentation and presentations |

---

## Build Commands

```bash
# Build, test and install all packages
make all

# Build only
make build

# Test only
make test

# Clean up
make clean

# Test coverage
make coverage

# Start slideshow (Port 3999)
make slideshow
```

---

## Notes

- This project requires at least **Go 1.24.0** due to updated dependencies.
- The `sp/` (Systems Programming) examples are Linux-specific.
- The `raft/` package is skipped during `make all` (slow tests).
