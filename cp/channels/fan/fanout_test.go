// Copyright 2018-2026 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package fan contains fan-in and fan-out functions for channel operations.
package fan

import (
	"testing"
	"time"
)

func TestFanOut(t *testing.T) {

	inCh := make(chan int, 10)
	// write data to in channel
	for i := 0; i < 10; i++ {
		inCh <- i
	}

	task := func(v int, res chan int) {
		time.Sleep(1 * time.Millisecond) // simulate long running calculation
		res <- v * v                     // return v*v
	}

	outChan := FanOut(inCh, task)

	Print(outChan)

	time.Sleep(100 * time.Millisecond)

	close(inCh)
}
