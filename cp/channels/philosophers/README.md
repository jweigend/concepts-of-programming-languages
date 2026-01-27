# Dining Philosophers (Channels)

## Purpose

This example implements the classic **Dining Philosophers Problem** using Go channels for synchronization. This demonstrates deadlock-free concurrent resource acquisition.

## The Problem

- 5 philosophers sit at a round table
- Each needs 2 forks to eat (shared with neighbors)
- Challenge: Prevent deadlock and starvation

## Key Concepts

- **Channels for coordination** - Philosophers request forks via channels
- **Table as coordinator** - Central arbiter for fork allocation
- **Deadlock prevention** - Atomic acquisition of both forks

## Components

| Component | Role |
|-----------|------|
| `Philosopher` | Goroutine that thinks and eats |
| `Table` | Manages fork allocation via channels |

## Run

```bash
go test -v
```

## Test

The test in `philosophers_test.go` runs the simulation and verifies that all philosophers can eat without deadlock.

```bash
go test ./...
```
