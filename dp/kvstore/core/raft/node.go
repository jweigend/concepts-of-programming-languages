package raft

import (
	"math/rand"
	"time"
)

// Node is a RAFT node.
type Node struct {
	uuid          int
	config        *Configuration
	statemachine  *Statemachine
	replicatedLog *ReplicatedLog
	timeOutTicker *time.Ticker
	currentTerm   int
	votedFor      int
}

// NewNode constructs a node.
func NewNode(uuid int, config *Configuration) *Node {
	node := new(Node)
	node.uuid = uuid
	node.config = config
	node.currentTerm = 0
	node.votedFor = 0
	node.statemachine = NewStatemachine()
	node.replicatedLog = NewReplicatedLog()
	return node
}

func (n *Node) Start() {
	// todo move time interval to configuration
	timeOutTicker := time.NewTicker(time.Duration(100+rand.Intn(1000)) * time.Millisecond)
	go func() {
		for range timeOutTicker.C {
			n.Timeout()
		}
	}()
}

func (n *Node) Stop() {
	n.timeOutTicker.Stop()
}

func (n *Node) Timeout() {
	if n.statemachine.current == FOLLOWER {
		n.StartElection()
	}
}

func (n *Node) StartElection() {
	n.statemachine.Next(CANDIDATE)

}
