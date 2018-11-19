// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package resourcemanager

import (
	"errors"
	"fmt"
	"sync"
)

// ResourceManager detects deadlocks by finding cycles in a Ressource Allocation Graph.
// blocks, when a resource is in use. The user is responsible to implement a resource release strategy, when a deadlock is recognized.
// Sample
// Acquire ("P1", "R1" ) : P1 <- R1 (ok) - Process 1 got Resource 1
// Acquire ("P1", "R3" ) : P1 <- R3 (ok) - Process 1 got Resource 3
// Acquire ("P2", "R3" ) : P2 <- R2 (ok) - Process 2 got Resource 2
// Acquire ("P2", "R1" ) : P2 -> R1 (wait) - Process 2 cant get Resource 1 (in use by Process 1) : wait
// Acquire ("P1", "R2" ) returns false : P1 -> R2 (deadlock) - acquire will recognize the deadlock and raturns false
type ResourceManager struct {
	graph *ResourceGraph
	m     sync.Mutex
	c     sync.Cond
}

// NewResourceManager creates a Resource Manager.
func NewResourceManager() *ResourceManager {
	manager := new(ResourceManager)
	manager.c = sync.Cond{L: &manager.m}
	manager.graph = NewResourceGraph()
	return manager
}

// Acquire tries to acquire a resource. The method blocks as long as the given resource is in use.
// The implementation recognizes a deadlock situation. In this case the method returns false.
// @param processName The name of the process.
// @param resourceName The name of the resource.
// @return  true if the Resource could be acquired for the given process.
//          false if a deadlock is detected.
//
func (r *ResourceManager) Acquire(processName string, resourceName string) bool {
	r.m.Lock()
	defer r.m.Unlock()

	if resourceName == "" || processName == "" {
		panic(errors.New("processname or resourcename can not be empty"))
	}

	r.graph.AddLink(processName, resourceName) // add Px -> Ry

	for r.resourceIsInUse(resourceName, processName) {
		if r.graph.DetectCycle(processName, resourceName) {
			r.graph.RemoveLink(processName, resourceName)
			return false // Deadlock detected
		}
		r.c.Wait()
	}

	r.graph.RemoveLink(processName, resourceName) // remove Px -> Ry
	r.graph.AddLink(resourceName, processName)    // add Ry -> Px

	return true // no deadlock
}

//
// Release the resource for a given process by removing it from the process-resource-graph.
//
func (r *ResourceManager) Release(processName string, resourceName string) {
	r.m.Lock()
	defer r.m.Unlock()

	r.graph.RemoveLink(resourceName, processName)
	r.c.Signal()
}

/**
 * A resource is in use when a process owns the resource :
 * R1 -> [Px]
 */
func (r *ResourceManager) resourceIsInUse(resourceName string, processName string) bool {
	process := r.graph.Get(resourceName)
	if process == nil {
		return false
	}
	return len(process) == 1
}

func (r *ResourceManager) String() string {
	return fmt.Sprintf("%v", r.graph.String())
}
