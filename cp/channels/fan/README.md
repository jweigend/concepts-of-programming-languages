# Fan-In / Fan-Out

## Purpose

This example demonstrates the **Fan-In** and **Fan-Out** concurrency patterns using Go channels.

## Key Concepts

- **Fan-Out** - Distribute work from one channel to multiple workers
- **Fan-In** - Collect results from multiple channels into one
- **Channel pipelines** - Composing concurrent operations

## Patterns

### Fan-In
Multiple input channels → Single output channel

```go
FanIn(inputChannels []chan int, output chan int)
```

### Fan-Out
Single input channel → Multiple worker goroutines

## Use Cases

- Parallel processing of independent tasks
- Load balancing across workers
- Aggregating results from multiple sources

## Run

```bash
go test -v
```

## Test

The tests in `fanin_test.go` and `fanout_test.go` verify correct data distribution and collection across multiple channels.

```bash
go test ./...
```
