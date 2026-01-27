# Dining Philosophers (Locks)

## Purpose

This example implements the **Dining Philosophers Problem** using **locks and a resource manager**. Compare this with the channel-based implementation in `cp/channels/philosophers/`.

## The Problem

- 5 philosophers sit at a round table
- Each needs 2 forks to eat (shared with neighbors)
- Challenge: Prevent deadlock and starvation

## Key Concepts

- **Resource Manager** - Detects and prevents deadlocks
- **Mutex synchronization** - Traditional locking approach
- **Resource Allocation Graph** - Cycle detection for deadlocks

## Components

| Component | Role |
|-----------|------|
| `Philosopher` | Goroutine that thinks and eats |
| `Table` | Holds references to forks |
| `ResourceManager` | Manages resource allocation |

## Run

```bash
go test -v
```

## Test

The test in `philosophers_test.go` runs the simulation and verifies deadlock-free execution.

```bash
go test ./...
```
