// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package resourcemanger

// ResourceGraph is a simple RAG ResourceAllocationGraph.
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
	destinations := (*r).edges[dest]
	destinations = append(destinations, source)
	(*r).edges[dest] = destinations
}

// RemoveLink removes a link from the graph.
func (r *ResourceGraph) RemoveLink(source, dest string) {
	destinations := (*r).edges[dest]
	destinations = removeItem(destinations, source)
	(*r).edges[dest] = destinations
}

// DetectCycle reports true if there is a cycle between source and dest
func (r *ResourceGraph) DetectCycle(source string, dest string) bool {
	return r.detectCycle1(source, source, dest)
}

// Internal helper does the work.
func (r *ResourceGraph) detectCycle1(first string, source string, dest string) bool {

	if first == dest {
		return true
	}
	result := false
	destinations := (*r).edges[dest]
	if destinations == nil {
		return result
	}

	for _, element := range destinations {
		result = result || r.detectCycle1(first, dest, element)
	}

	return result
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
