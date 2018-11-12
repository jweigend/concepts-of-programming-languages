![Logo](../../doc/article/img/Logo.png "Logo")
### *Concurrency with Go*

- Blocking Queue
- Resource Manager

## Using the go sync package to write a Java like Blocking Queue.
**Go Channels are great. No Question. But sometimes you want to have more control about when to block and what to do if some condition is reached.**
The Go sync package offers the same possibilities as java.util.concurrent does.


```go
import "sync"

// BlockingQueue is a FIFO container with a fixed capacity.
// It  blocks a reader when it is empty and a writer when it is full.
type BlockingQueue struct {
	m        sync.Mutex
	c        sync.Cond
	data     []interface{}
	capacity int
}
```

```go
// NewBlockingQueue constructs a BlockingQuee with a given capacity.
func NewBlockingQueue(capacity int) *BlockingQueue {
	q := new(BlockingQueue)
	q.c = sync.Cond{L: &q.m}
	q.capacity = capacity
	return q
}
```

```go
// Put puts an item in the queue and blocks it the queue is full.
func (q *BlockingQueue) Put(item interface{}) {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	for q.isFull() {
		q.c.Wait()
	}
	q.data = append(q.data, item)
	q.c.Signal()
}
```


```go
// Take takes an item from the queue and blocks if the queue is empty.
func (q *BlockingQueue) Take() interface{} {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	for q.isEmpty() {
		q.c.Wait()
	}
	result := q.data[0]
	q.data = q.data[1 : len(q.data)]	
	q.c.Signal()	
	return result
}
```

```go
// isFull returns true if the capacity is reached.
func (q *BlockingQueue) isFull() bool {
	return len(q.data) == q.capacity
}

// isEmpty returns true if there are no elements in the queue.
func (q *BlockingQueue) isEmpty() bool {
	return len(q.data) == 0
}
```


```go
func TestBlockingQueue(t *testing.T) {

	bq := NewBlockingQueue(1)
	done := make(chan bool)

	// slow writer
	go func() {
		bq.Put("A")
		time.Sleep(100 * time.Millisecond)
		bq.Put("B")
		time.Sleep(100 * time.Millisecond)
		bq.Put("C")
	}()

	// reader will be blocked
	go func() {
		item := bq.Take()
		fmt.Printf("Got %v\n", item)
		item = bq.Take()
		fmt.Printf("Got %v\n", item)
		item = bq.Take()
		fmt.Printf("Got %v\n", item)
		done <- true
	}()

	// block while done
	<-done
}
```