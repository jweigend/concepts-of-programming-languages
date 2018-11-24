package raft

import (
	"testing"
	"time"
)

func TestHeartbeat(t *testing.T) {
	n := NewNode(0)
	n.statemachine.Next(CANDIDATE)
	n.statemachine.Next(LEADER)
	// startHeartbeat is only allowed in leader state
	n.startHeartbeat()
	// wait 5 millisecond --> check console output
	time.Sleep(5000 * time.Millisecond)
}

func TestElection(t *testing.T) {

	n1 := NewNode(0)
	n2 := NewNode(1)
	n3 := NewNode(2)
	nodes := []NodeRPC{n1, n2, n3}

	c := NewCluster(nodes)

	n1.Start(c)
	n2.Start(c)
	n3.Start(c)

	time.Sleep(15000 * time.Millisecond)

}
