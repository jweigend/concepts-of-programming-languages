# Blocking Queue (Locks)

## Purpose

This example demonstrates a **thread-safe blocking queue** implemented using **mutexes and condition variables**. Compare this with the channel-based implementation in `cp/channels/blockingqueue/`.

## Key Concepts

- **`sync.Mutex`** - Mutual exclusion for critical sections
- **`sync.Cond`** - Condition variables for waiting/signaling
- **Lock-based synchronization** - Traditional approach (Java-style)

## API

```go
q := NewBlockingQueue(capacity)
q.Put(item)        // Blocks if queue is full
item := q.Take()   // Blocks if queue is empty
```

## How It Works

1. `Put()` acquires lock, waits while full, adds item, signals
2. `Take()` acquires lock, waits while empty, removes item, signals

## Comparison: Channels vs Locks

| Aspect | Channels | Locks |
|--------|----------|-------|
| Code | Simpler | More verbose |
| Flexibility | Limited | Full control |
| Go idiom | Preferred | When needed |

## Run

```bash
go test -v
```

## Test

The test in `blockingqueue_test.go` verifies blocking behavior with concurrent producers and consumers.

```bash
go test ./...
```
