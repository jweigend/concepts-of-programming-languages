// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

import (
	"testing"
	"time"
)

func TestBigCluster(t *testing.T) {

	n1 := NewNode(0)
	n2 := NewNode(1)
	n3 := NewNode(2)
	n4 := NewNode(3)
	n5 := NewNode(4)

	nodes := []*Node{n1, n2, n3, n4, n5}

	cluster := NewCluster(nodes)

	cluster.StartAll()

	time.Sleep(10000 * time.Millisecond)

	ok, err := cluster.Check()
	if !ok {
		t.Error(err)
	}

	cluster.StopAll()
	time.Sleep(1000 * time.Millisecond) // wait for grace shutdown
}

func TestHeartbeat(t *testing.T) {
	// single node cluster
	n1 := NewNode(0)
	nodes := []*Node{n1}
	n1.cluster = NewCluster(nodes)

	// startHeartbeat is only allowed in leader state
	n1.statemachine.Next(CANDIDATE)
	n1.statemachine.Next(LEADER)
	n1.heartbeatTimer.resetC <- true

	// wait two second --> check console output
	time.Sleep(2000 * time.Millisecond)
	n1.Stop()
}
