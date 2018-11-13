// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package fan contains examples for fanout und fanin functions using channels.
package utils

import (
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {

	// FanIn reads from an array of input channels and send the results to a output channel .
	inCh := []chan int{make(chan int), make(chan int)}
	outCh := make(chan int)

	// FanIn from an array of channels
	FanIn(inCh, outCh)

	// read from one channel and print results to stdout.
	AsyncReadAndPrintFromCh(outCh)

	inCh[0] <- 2
	inCh[1] <- 1

	time.Sleep(100 * time.Millisecond)

	close(inCh[0])
	close(inCh[1])

	time.Sleep(100 * time.Millisecond)
}
