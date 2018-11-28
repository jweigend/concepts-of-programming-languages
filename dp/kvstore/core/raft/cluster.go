// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

import "errors"

// Cluster knows the RPC interface of all members.
// The cluster works with multiple nodes for testing or with RPC proxies for remote access in a distributed scenario.
type Cluster struct {
	allNodes []*Node
}

// NewCluster constructs a new cluster with all Node RPC interfaces.
// For local test the NodeRPC is the Node itself. For distributed operation the NodeRPC is a proxy.
func NewCluster(allNodes []*Node) *Cluster {
	return &Cluster{allNodes}
}

// GetRemoteNodes return the NodeRPC If of all nodes.
func (c *Cluster) GetRemoteFollowers(leaderID int) []NodeRPC {
	buf := make([]NodeRPC, len(c.allNodes)-1)
	j := 0
	for i, n := range c.allNodes {
		if i == leaderID {
			continue
		}
		buf[j] = n
		j++
	}
	return buf
}

// StopAll stops all nodes in the cluster.
func (c *Cluster) StartAll() {
	for _, n := range c.allNodes {
		n.Start(c.GetRemoteFollowers(n.id))
	}
}

// StopAll stops all nodes in the cluster.
func (c *Cluster) StopAll() {
	for _, n := range c.allNodes {
		n.Stop()
	}
}

// StopLeader stops the leader in the cluster.
func (c *Cluster) StopLeader() *Node {
	for _, n := range c.allNodes {
		if n.isLeader() {
			n.Stop()
			return n
		}
	}
	return nil // no leader found
}

// Check checks if a cluster is in a valid state.
// There should be exact one leader.
func (c *Cluster) Check() (bool, error) {
	leaderCount := 0
	followerCount := 0
	candidateCount := 0
	for _, n := range c.allNodes {
		if n.isLeader() {
			leaderCount++
		} else if n.isCandidate() {
			candidateCount++
		} else if n.isFollower() {
			followerCount++
		}
	}
	if leaderCount > 1 {
		return false, errors.New("there are multiple leaders in the cluster")
	} else if leaderCount == 0 {
		return false, errors.New("there is no leader in the cluster")
	} else if followerCount < len(c.allNodes)/2 {
		return false, errors.New("there are not enough followers -> split brain")
	}
	return true, nil
}
