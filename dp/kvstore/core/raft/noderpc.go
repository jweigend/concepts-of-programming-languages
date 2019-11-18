// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

// NodeRPC is the remote interface for Node to Node communication in the RAFT cluster.
type NodeRPC interface {

	// AppendEntries is used by the Leader to replicate logs and it is  used as heartbeat.
	// The Leader will call the AppendEntries method on all nodes in the cluster.
	// Arguments
	// - term         : leaders term
	// - leaderId     : leadersId to redirect calls to leader
	// - prevLogIndex : Index of log entry immediately processing
	// - prevLogTerm  : Term of the prevLogIndex entry
	// - entries      : log entries to store (empty for heartbeat)
	// - leaderCommit : leaders commit index
	// Results
	// - term         : current termin (for leader, update itself)
	// - success      : true if follower contained entry matching prevLogIndex and prevLogTerm
	AppendEntries(term, leaderID, prevLogIndex, prevLogTerm int, entries []string,
		leaderCommit int) (int, bool)

	// RequestVote is called by candidates to gather votes.
	// It returns the current term to update the candidate
	// It returns true when the candidate received vote.
	// Arguments
	// - term         : candidates term
	// - candidateID  : candidate requesting vote
	// - lastLogIndex : index of candidates last log entry
	// - lastLogTerm  : term of candidates last log entry
	// Results
	// - term         : the current term, for candidate to update itself
	// - voteGranted  : true means candidate received vote
	RequestVote(term, candidateID, lastLogIndex, lastLogTerm int) (int, bool)
}
