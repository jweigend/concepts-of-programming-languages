package raft

import (
	"math/rand"
	"time"
)

// Node is a RAFT node.
type Node struct {
	id            int
	config        *Configuration
	statemachine  *Statemachine
	replicatedLog *ReplicatedLog
	timeOutTicker *time.Ticker
	currentTerm   int
	votedFor      int
}

// NewNode constructs a node.
func NewNode(id int, config *Configuration) *Node {
	node := new(Node)
	node.id = id
	node.config = config
	node.currentTerm = 0
	node.votedFor = 0
	node.statemachine = NewStatemachine()
	node.replicatedLog = NewReplicatedLog()
	return node
}

func (n *Node) Start() {
	// todo move time interval to configuration
	n.timeOutTicker = time.NewTicker(time.Duration(100+rand.Intn(1000)) * time.Millisecond)
	go func() {
		for range n.timeOutTicker.C {
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
	} else if n.statemachine.current == LEADER {

	}
}

func (n *Node) StartElection() {
	n.statemachine.Next(CANDIDATE)

}

// AppendEntries implementation is used as heardbeat and log replication.
func (n *Node) AppendEntries(term, leaderID, prevLogIndex, prevLogTermin int, entries []string, leaderCommit int) (currentTerm int, success bool) {
	if n.statemachine.Current() == LEADER {
		return n.currentTerm, false
	} else if term < n.currentTerm {
		return n.currentTerm, false
	}

	// heartbeat
	if len(entries) == 0 {
		//ResetElectionTimer()
	}

	// todo

	return n.currentTerm, true
}

// RequestVote is called by candidates to gather votes.
// It returns the current term to update the candidate
// It returns true when the candidate received vote.
func (n *Node) RequestVote(term, candidateID, lastLogIndex, lastLogTerm int) (int, bool) {
	// todo
	return 0, false
}
