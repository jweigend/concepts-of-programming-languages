// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

// NodeRPC is the remote interface for Node to Node communication in the RAFT cluster.
type NodeRPC interface {

	// AppendEntries is used by the Leader to replicate logs and it is  used as heartbeat.
	// The Leader will call the AppendEntries method on all nodes in the cluster.
	// Parmeters:
	//  term     : leaders term
	//  leaderId : leadersId to redirect calls to leader
	AppendEntries(term, leaderID, prevLogIndex, prevLogTermin int, entries []string, leaderCommit int) (int, bool)

	// RequestVote is called by candidates to gather votes.
	// It returns the current term to update the candidate
	// It returns true when the candidate received vote.
	RequestVote(term, candidateID, lastLogIndex, lastLogTerm int) (int, bool)
}
