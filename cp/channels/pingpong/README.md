# Ping Pong

## Purpose

This example demonstrates **goroutine communication** through channels with a simple ping-pong game between two players.

## Key Concepts

- **Unbuffered channels** - Synchronous communication
- **Goroutines** - Concurrent execution
- **Message passing** - Sharing data by communicating

## How It Works

1. Two goroutines (`ping` and `pong`) take turns
2. They pass a `Ball` struct back and forth via a channel
3. Each hit increments a counter
4. Main goroutine controls game duration

```go
table := make(chan *Ball)
go player("ping", table)
go player("pong", table)
table <- new(Ball)  // Start game
```

## Run

```bash
go run pingpong.go
```

## Output

```
ping 1
pong 2
ping 3
pong 4
...
```

## Test

This example has no tests. It is a runnable demonstration.
