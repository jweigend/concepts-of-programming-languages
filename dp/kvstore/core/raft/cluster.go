// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

// Cluster knows the RPC interface of all members.
// The cluster works with multiple nodes for testing or with RPC proxies for remote access in a distributed scenario.
type Cluster struct {
	allNodes []NodeRPC
}

// NewCluster constructs a new cluster with all Node RPC interfaces.
// For local test the NodeRPC is the Node itself. For distributed operation the NodeRPC is a proxy.
func NewCluster(allNodes []NodeRPC) *Cluster {
	return &Cluster{allNodes}
}

// GetFollowers returns the RPC interfaces of all followers for a given leader
func (c *Cluster) GetFollowers(leaderID int) []NodeRPC {
	buf := make([]NodeRPC, len(c.allNodes))
	copy(buf, c.allNodes)
	result := append(buf[:leaderID], buf[leaderID+1:]...)
	return result
}
