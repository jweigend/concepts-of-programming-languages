// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package resourcemanger

import "errors"

// ResourceManager detects deadlocks by finding cycles in a Ressource Allocation Graph.
// blocks, when a resource is in use. The user is responsible to implement a resource release strategy, when a deadlock is recognized.
// Sample
// Aquire ("P1", "R1" ) : P1 <- R1 (ok) - Process 1 got Resource 1
// Aquire ("P1", "R3" ) : P1 <- R3 (ok) - Process 1 got Resource 3
// Aquire ("P2", "R3" ) : P2 <- R2 (ok) - Process 2 got Resource 2
// Aquire ("P2", "R1" ) : P2 -> R1 (wait) - Process 2 cant get Resouce 1 (in use by Process 1) : wait
// aquire ("P1", "R2" ) returns false : P1 -> R2 (deadlock) - aquire will recognize the deadlock and raturns false
type ResourceManager struct {
	graph ResourceGraph
}

// NewResourceManager creates a Resource Manager.
func NewResourceManager() *ResourceManager {
	manager := new(ResourceManager)
	(*manager).graph = *NewResourceGraph()
	return manager
}

// Aquire tries to aquire a resource. The method blocks as long as the given resource is in use.
// The implementation recognizes a deadlock situation. In this case the method returns false.
// @param processName The name of the process.
// @param resourceName The name of the resource.
// @return  true if the Resource could be aquired for the given process.
//          false if a deadlock is detected.
//
func (r *ResourceManager) Aquire(processName string, resourceName string) bool {

	if resourceName == "" || processName == "" {
		panic(errors.New("processname or resourcename was empty"))
	}

	r.graph.AddLink(processName, resourceName) // add Px -> Ry

	for r.resourceIsInUse(resourceName, processName) {
		if r.graph.DetectCycle(processName, resourceName) {
			r.graph.RemoveLink(processName, resourceName)
			return false // Deadlock detected
		}
		// wait() // todo wait on channel
	}

	r.graph.RemoveLink(processName, resourceName) // remove Px -> Ry
	r.graph.AddLink(resourceName, processName)    // add Ry -> Px

	return true // no deadlock
}

//
// Release the resource for a given process by removing it from the process-resource-graph.
//
func (r *ResourceManager) release(processName string, resourceName string) {
	r.graph.RemoveLink(resourceName, processName)
	//notifyAll()
}

/**
 * A resource is in use when a process owns the resource :
 * R1 -> [Px]
 */
func (*ResourceManager) resourceIsInUse(resourceName string, processName string) bool {
	//List<String> process = graph.get(resourceName);
	//if (process == null) return false;
	//return process.size() == 1;
	return true
}
