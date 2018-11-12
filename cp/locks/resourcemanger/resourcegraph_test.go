// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package resourcemanger

import (
	"testing"
)

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
