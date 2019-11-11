// Copyright 2018 Johannes Weigend, copied from Johannes Siedersleben
// Licensed under the Apache License, Version 2.0

// Package blockingqueue contains a blocking LIFO container.
package blockingqueue

import "sync"

// BlockingQueue is a FIFO container with a fixed capacity.
// It  blocks a reader when it is empty and a writer when it is full.
type BlockingQueue struct {
	m        sync.Mutex
	c        sync.Cond
	data     []interface{}
	capacity int
}

// NewBlockingQueue constructs a BlockingQuee with a given capacity.
func NewBlockingQueue(capacity int) *BlockingQueue {
	q := new(BlockingQueue)
	q.c = sync.Cond{L: &q.m}
	q.capacity = capacity
	return q
}

// A1

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

// Take takes an item from the queue and blocks if the queue is empty.
func (q *BlockingQueue) Take() interface{} {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	for q.isEmpty() {
		q.c.Wait()
	}
	result := q.data[0]
	q.data = q.data[1:len(q.data)]
	q.c.Signal()
	return result
}

// A2

// isFull returns true if the capacity is reached.
func (q *BlockingQueue) isFull() bool {
	return len(q.data) == q.capacity
}

// isEmpty returns true if there are no elements in the queue.
func (q *BlockingQueue) isEmpty() bool {
	return len(q.data) == 0
}
