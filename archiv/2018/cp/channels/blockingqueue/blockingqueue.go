// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package blockingqueue contains a blocking LIFO container.
package blockingqueue

// BlockingQueue is a FIFO container with a fixed capacity.
// It  blocks a reader when it is empty and a writer when it is full.
type BlockingQueue struct {
	channel chan interface{}
}

// NewBlockingQueue constructs a BlockingQueue with a given capacity.
func NewBlockingQueue(capacity int) *BlockingQueue {
	q := BlockingQueue{make(chan interface{}, capacity)}
	return &q
}

// Put puts an item in the queue and blocks it the queue is full.
func (q *BlockingQueue) Put(item interface{}) {
	q.channel <- item
}

// Take takes an item from the queue and blocks if the queue is empty.
func (q *BlockingQueue) Take() interface{} {
	return <-q.channel
}

// EOF
