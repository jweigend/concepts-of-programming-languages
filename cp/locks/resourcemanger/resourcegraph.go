// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package resourcemanger contains a blocking resource manger which avoids deadlocks
package resourcemanger

import (
	"encoding/json"
	"fmt"
)

// ResourceGraph is a simple RAG ResourceAllocationGraph.
// A graph is a collection of edges. Each edge is of the form
// edge := source -> [target1, target2, ... targetN]
type ResourceGraph struct {
	edges map[string][]string
}

// NewResourceGraph creates an empty graph.
func NewResourceGraph() *ResourceGraph {
	resourceGraph := new(ResourceGraph)
	resourceGraph.edges = make(map[string][]string)
	return resourceGraph
}

// AddLink adds a link to the graph.
func (r *ResourceGraph) AddLink(source, dest string) {
	destinations := r.edges[source]
	destinations = append(destinations, dest)
	r.edges[source] = destinations
}

// RemoveLink removes a link from the graph.
func (r *ResourceGraph) RemoveLink(source, dest string) {
	destinations := r.edges[source]
	destinations = removeItem(destinations, dest)
	if len(destinations) > 0 {
		r.edges[source] = destinations
	} else {
		delete(r.edges, source) // remove entry complete from map
	}
}

// DetectCycle reports true if there is a cycle between source and dest
func (r *ResourceGraph) DetectCycle(source string, dest string) bool {
	return r.detectCycle1(source, source, dest)
}

// Get destinations.
func (r *ResourceGraph) Get(dest string) []string {
	destinations := r.edges[dest]
	return destinations
}

// Internal helper does the work.
func (r *ResourceGraph) detectCycle1(first string, source string, dest string) bool {

	if first == dest {
		return true
	}
	result := false
	destinations := r.edges[source]
	if destinations == nil {
		return result
	}

	for _, element := range destinations {
		result = result || r.detectCycle1(first, dest, element)
	}

	return result
}

func (r *ResourceGraph) String() string {
	b, _ := json.MarshalIndent(r.edges, "", "  ")
	return fmt.Sprintf("%v", string(b))
}

// Helper removes an item from a string list.
func removeItem(list []string, item string) []string {

	for i := len(list) - 1; i >= 0; i-- {
		v := list[i]
		if v == item {
			// found
			list[i] = list[len(list)-1]
			return list[0 : len(list)-1]
		}
	}
	return list // not found
}
