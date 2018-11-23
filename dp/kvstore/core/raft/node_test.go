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
