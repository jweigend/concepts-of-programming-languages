package blockingqueue

import (
	"fmt"
	"testing"
	"time"
)

// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

func TestBlockingQueue(t *testing.T) {
	bq1 := NewBlockingQueue(1)
	done := make(chan bool)
	// slow writer
	go func(bq *BlockingQueue) {
		bq.Put("A")
		time.Sleep(100 * time.Millisecond)
		bq.Put("B")
		time.Sleep(100 * time.Millisecond)
		bq.Put("C")
	}(bq1)
	// reader will be blocked
	go func(bq *BlockingQueue) {
		item := bq.Take()
		fmt.Printf("Got %v\n", item)
		item = bq.Take()
		fmt.Printf("Got %v\n", item)
		item = bq.Take()
		fmt.Printf("Got %v\n", item)
		done <- true
	}(bq1)

	<-done
}

// EOF OMIT
