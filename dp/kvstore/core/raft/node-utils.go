// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package raft is an implementation of the RAFT consensus algorithm.
package raft

import (
	"log"
	"time"
)

func (n *Node) log(msg string) {
	log.Printf("[%v] [%v] [%v] : %v", n.id, n.statemachine.Current(), n.currentTerm, msg)
}

func (n *Node) isLeader() bool {
	return n.statemachine.Current() == LEADER
}

func (n *Node) isFollower() bool {
	return n.statemachine.Current() == FOLLOWER
}

func (n *Node) isCandidate() bool {
	return n.statemachine.Current() == CANDIDATE
}

func (n *Node) isNotLeader() bool {
	return n.isFollower() || n.isCandidate()
}

type timercontrol struct {
	stopC  chan bool
	resetC chan bool
}

func createPeriodicTimer(d func() time.Duration, timeout func()) timercontrol {
	stopC := make(chan bool)
	resetC := make(chan bool)
	go func() {
		timer := time.NewTimer(d())
		timer.Stop() // Timer must be started explicit by using the reset channel
		for {
			select {
			case <-stopC:
				// log.Println("Timer Stopped")
				timer.Stop()
				break
			case <-timer.C:
				// log.Println("Timer Timeout")
				timer.Stop()
				timer.Reset(d())
				go timeout()
				break
			case <-resetC:
				// log.Println("Timer Reset")
				timer.Stop()
				timer.Reset(d())
			}
		}
	}()
	return timercontrol{stopC, resetC}
}
