# Systems Programming

## Purpose

This directory contains examples of **low-level systems programming** in Go, including Linux container implementation and system calls.

## ⚠️ Linux Only

These examples use Linux-specific system calls and will not compile on macOS or Windows.

## Files

| File | Description |
|------|-------------|
| `container.go` | Simple Docker-like container implementation |
| `empty.go` | Placeholder for non-Linux builds |
| `samples/` | Additional examples |

## Container Example

Based on Liz Rice's container implementation. Demonstrates:

- **Linux namespaces** - PID, UTS, Mount isolation
- **System calls** - `clone()`, `unshare()`
- **Cgroups** - Resource limiting
- **chroot** - Filesystem isolation

## Key Concepts

- **`syscall` package** - Direct kernel interaction
- **`/proc` filesystem** - Process information
- **Namespace flags** - `CLONE_NEWPID`, `CLONE_NEWUTS`, `CLONE_NEWNS`

## Run

```bash
# Must run as root on Linux
sudo go run container.go run /bin/sh
```

## Test

This example has no tests. It is a runnable demonstration (Linux only).
