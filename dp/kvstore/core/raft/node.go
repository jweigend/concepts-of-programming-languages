package raft

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Node is a node in a Raft consensus cluster. It is called "server" in the original Raft paper.
// Node seems to be more accurate because for testing purposes we can run multiple nodes in a single server process.
type Node struct {
	id             int
	statemachine   *Statemachine
	replicatedLog  *ReplicatedLog
	electionTimer  *time.Timer
	heartbeatTimer *time.Timer // runs only if the node is in MASTER state
	currentTerm    int
	votedFor       int
	cluster        *Cluster
}

// NewNode constructs a node.
func NewNode(id int) *Node {
	node := new(Node)
	node.id = id
	node.currentTerm = 0
	node.votedFor = 0
	node.statemachine = NewStatemachine()
	node.replicatedLog = NewReplicatedLog()
	return node
}

// Start setup timers for heartbeat and election.
func (n *Node) Start(cluster *Cluster) {
	n.cluster = cluster
	n.restartElectionTimer()
}

// Stop stops timers.
func (n *Node) Stop() {
	if n.electionTimer != nil {
		n.electionTimer.Stop()
	}
	if n.heartbeatTimer != nil {
		n.heartbeatTimer.Stop()
	}
}

// restartElectionTimer restarts random timer.
func (n *Node) restartElectionTimer() {
	if n.electionTimer != nil {
		n.electionTimer.Stop()
	}
	n.electionTimer = time.NewTimer(time.Duration(2000+rand.Intn(3000)) * time.Millisecond)
	go func() {
		<-n.electionTimer.C
		n.electionTimeout()
	}()
}

// startHeartbeat starts an heartbeat and runs forever until the timer ist stopped.
func (n *Node) startHeartbeat() {
	if n.heartbeatTimer != nil {
		n.heartbeatTimer.Stop()
	}
	n.heartbeatTimer = time.NewTimer(time.Duration(500+rand.Intn(1000)) * time.Millisecond)
	go func() {
		_, ok := <-n.heartbeatTimer.C // check this: If the time was stopped, the channel should return false (closed?)
		if ok {
			n.sendHeartbeat()
			defer n.startHeartbeat() // restart again
		}
	}()
}

// sendHeartbeat
func (n *Node) sendHeartbeat() {
	if n.statemachine.current != LEADER {
		panic("setHeatbeat should only run on LEADER")
	}
	log.Printf("[%v] SendHeartbeat on Node: %v", n.statemachine.Current(), n.id)
}

// electionTimeout happens when a node receives no heartbeat in a given time period.
func (n *Node) electionTimeout() {
	if n.statemachine.current == LEADER {
		panic("The election timeout should not happen, when a node is LEADER.")
	}
	n.startElectionProcess()
}

// startElectionProcess sends a RequestVote request to other members in the cluster.
// if successful - we get are the new leader in a new term.
func (n *Node) startElectionProcess() {
	n.statemachine.Next(CANDIDATE)
	n.currentTerm++ // new term starts now
	electionWon := n.executeElection()
	if electionWon {
		n.log(fmt.Sprintf("Election won. Starting as leader."))
		n.startLeader()
	} else {
		n.log(fmt.Sprintf("Election was not won. Restarting election timer."))
		n.restartElectionTimer() // try again, split vote or cluster down
	}
}

// executeElection executes a leader election by sending RequestVote to other nodes.
// for all other nodes in the cluster RequestVote is sent
func (n *Node) executeElection() bool {
	rpcIfs := n.cluster.GetFollowers(n.id)
	var wg sync.WaitGroup
	votes := make([]bool, len(rpcIfs))
	wg.Add(len(rpcIfs))
	for i, rpcIf := range rpcIfs {
		go func(w *sync.WaitGroup, i int, rpcIf NodeRPC) {
			term, ok := rpcIf.RequestVote(n.currentTerm, n.id, 0, 0)
			if term > n.currentTerm {
				// convert to follower
				n.statemachine.Next(FOLLOWER)
				n.restartElectionTimer()
			}
			votes[i] = ok
			w.Done()
		}(&wg, i, rpcIf)
	}
	wg.Wait() // wait until all nodes have voted

	// Count votes
	nbrOfVotes := 0
	for _, vote := range votes {
		if vote {
			nbrOfVotes++
		}
	}
	// If more than 50% respond with true - The election was won!
	return (nbrOfVotes > len(rpcIfs)/2)
}

func (n *Node) startLeader() {
	n.statemachine.Next(LEADER)
	n.electionTimer = nil
	n.startHeartbeat()
}

// NodeRPC server implementation

// AppendEntries implementation is used as heardbeat and log replication.
func (n *Node) AppendEntries(term, leaderID, prevLogIndex, prevLogTermin int, entries []string, leaderCommit int) (currentTerm int, success bool) {
	if n.statemachine.Current() == LEADER {
		return n.currentTerm, false
	} else if term < n.currentTerm {
		return n.currentTerm, false
	}

	// heartbeat received in FOLLOWER -> reset election timer!
	if len(entries) == 0 {
		n.restartElectionTimer()
	} else {
		// todo: replicate logs
		log.Printf("[%v] AppendEntries replicate logs on Node: %v", n.statemachine.Current(), n.id)

	}

	return n.currentTerm, true
}

// RequestVote is called by candidates to gather votes.
// It returns the current term to update the candidate
// It returns true when the candidate received vote.
func (n *Node) RequestVote(term, candidateID, lastLogIndex, lastLogTerm int) (int, bool) {
	if term <= n.currentTerm {
		return n.currentTerm, false
	}
	n.currentTerm = term // ok: we join the master term
	n.log(fmt.Sprintf("RequestVote received from MASTER: %v. Vote OK.", candidateID))

	return n.currentTerm, true
}

func (n *Node) log(msg string) {
	log.Printf("[%v] [%v] [%v] : %v", n.id, n.statemachine.Current(), n.currentTerm, msg)
}
