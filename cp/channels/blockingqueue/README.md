# Blocking Queue (Channels)

## Purpose

This example demonstrates a **thread-safe blocking queue** implemented using Go **channels**. This is the idiomatic Go approach for concurrent data structures.

## Key Concepts

- **Buffered channels** - Fixed capacity queue
- **Channel blocking** - Automatic blocking when full/empty
- **Goroutine-safe** - No explicit locks needed

## API

```go
q := NewBlockingQueue(capacity)
q.Put(item)        // Blocks if queue is full
item := q.Take()   // Blocks if queue is empty
```

## How It Works

The queue uses a buffered channel internally. The channel's built-in blocking semantics handle all synchronization automatically.

## Run

```bash
go test -v
```

## Test

The test in `blockingqueue_test.go` verifies the blocking behavior with multiple producer and consumer goroutines.

```bash
go test ./...
```
