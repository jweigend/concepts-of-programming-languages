// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

import (
	"log"
)

func (n *Node) log(msg string) {
	log.Printf("[%v] [%v] [%v] : %v", n.id, n.statemachine.Current(), n.currentTerm, msg)
}

func (n *Node) isLeader() bool {
	return n.statemachine.Current() == LEADER
}

func (n *Node) isFollower() bool {
	return n.statemachine.Current() == FOLLOWER
}

func (n *Node) isCandidate() bool {
	return n.statemachine.Current() == CANDIDATE
}

func (n *Node) isNotLeader() bool {
	return n.isFollower() || n.isCandidate()
}
