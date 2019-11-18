// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

import (
	"log"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {

	tc := createPeriodicTimer(func() time.Duration {
		return time.Duration(1000) * time.Millisecond
	}, func() { log.Println("Timeout") })

	tc.resetC <- true
	time.Sleep(2 * time.Second)
	tc.stopC <- true
	time.Sleep(1 * time.Second)
	tc.resetC <- true
	time.Sleep(500 * time.Millisecond)
	tc.resetC <- true
	time.Sleep(500 * time.Millisecond)
	tc.resetC <- true

	time.Sleep(5 * time.Second)
}
