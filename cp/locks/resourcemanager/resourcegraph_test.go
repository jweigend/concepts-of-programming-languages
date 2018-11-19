// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package resourcemanager

import (
	"fmt"
	"testing"
	"time"
)

// TestResourceGraph tests the resource graph.
func TestResourceGraph(t *testing.T) {
	graph := NewResourceGraph()

	graph.AddLink("a", "b")
	graph.AddLink("a", "c")
	graph.AddLink("b", "c")

	// check no cycle
	if graph.DetectCycle("c", "a") {
		t.Error("c->a is not a cycle!")
	}

	// add a cycle
	graph.AddLink("c", "a")
	if !graph.DetectCycle("c", "a") {
		t.Error("c->a is a cycle!")
	}

	graph.RemoveLink("c", "a")
	if graph.DetectCycle("c", "a") {
		t.Error("c->a is not a cycle!")
	}
}

func TestSimpleResourceManager(t *testing.T) {
	manager := NewResourceManager()
	manager.Acquire("P1", "R1")
	manager.Acquire("P2", "R2")
	go manager.Acquire("P1", "R2") // should block
	fmt.Printf("%v", manager)
	manager.Release("P2", "R2")
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("%v", manager)
}

// TestResourceManager tests the manager which blocks if a resource is in use and no deadlock is detected.
// Acquire ("P1", "R1" ) : P1 <- R1 (ok) - Process 1 got Resource 1
// Acquire ("P1", "R3" ) : P1 <- R3 (ok) - Process 1 got Resource 3
// Acquire ("P2", "R2" ) : P2 <- R2 (ok) - Process 2 got Resource 2
// Acquire ("P2", "R1" ) : P2 -> R1 (wait) - Process 2 cant get Resource 1 (in use by Process 1) : wait
// Acquire ("P1", "R2" ) returns false : P1 -> R2 (deadlock) - acquire will recognize the deadlock and raturns false
func TestResourceManager(t *testing.T) {

	manager := NewResourceManager()
	p1 := make(chan bool)
	p2 := make(chan bool)

	go func() {
		manager.Acquire("P1", "R1")
		manager.Acquire("P1", "R3")
		p1 <- true
	}()
	// block p1 until finished
	<-p1

	go func() {
		manager.Acquire("P2", "R2")
		p2 <- false
		manager.Acquire("P2", "R1") // waits
		p2 <- true
	}()

	if manager.Acquire("P1", "R1") {
		t.Error("P1/R2 should detect a deadlock and return false")
	}

	if <-p2 {
		t.Error("P2 should block when requesting R1")
	}

	manager.Release("P1", "R1")

	if !<-p2 {
		t.Error("P2 should not block anymore when requesting R1")
	}
}
