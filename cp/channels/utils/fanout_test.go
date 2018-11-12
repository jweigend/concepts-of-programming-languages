// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package fan contains examples for fanout und fanin functions using channels.
package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestFanOut(t *testing.T) {

	inCh := make(chan int)
	task := func(v int, res chan int) {
		fmt.Printf("Doing something with v:%v\n", v)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Finished v: %v\n", v)
		res <- v * v // return v*v
	}

	outChan := FanOut(inCh, task)

	// write data to in channel
	for i := 0; i < 10; i++ {
		inCh <- i
	}

	AsyncReadAndPrintFromCh(outChan)

	time.Sleep(100 * time.Millisecond)

	close(inCh)
}
