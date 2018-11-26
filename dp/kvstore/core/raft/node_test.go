// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

import (
	"testing"
	"time"
)

func TestHeartbeat(t *testing.T) {
	// single node cluster
	n1 := NewNode(0)
	nodes := []NodeRPC{n1}
	n1.cluster = NewCluster(nodes)

	// startHeartbeat is only allowed in leader state
	n1.statemachine.Next(CANDIDATE)
	n1.statemachine.Next(LEADER)
	n1.startHeartbeat()

	// wait one second --> check console output
	time.Sleep(1000 * time.Millisecond)
}

func TestElection(t *testing.T) {

	n1 := NewNode(0)
	n2 := NewNode(1)
	n3 := NewNode(2)

	nodes := []NodeRPC{n1, n2, n3}
	cluster := NewCluster(nodes)

	n1.Start(cluster)
	n2.Start(cluster)
	n3.Start(cluster)

	time.Sleep(3000 * time.Millisecond)

	n1.Stop()
	n2.Stop()
	n3.Stop()
}
