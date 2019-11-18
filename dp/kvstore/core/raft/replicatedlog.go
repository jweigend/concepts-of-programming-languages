// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

// ReplicatedLog is the transactional log for RAFT.
type ReplicatedLog struct {
}

// NewReplicatedLog constructs a ReplicatedLog.
func NewReplicatedLog() *ReplicatedLog {
	return new(ReplicatedLog)
}
