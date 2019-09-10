# Exercise 7 - Concurrent Programming in Go

If you do not finish during the lecture period, please finish it as homework.

## Exercise 7.1 - Resource Allocation Graph

Build a type ResourceGraph. Use the following as starting point.
```go
// ResourceGraph is a simple RAG ResourceAllocationGraph.
// A graph is a collection of edges. Each edge is of the form
// edge := source -> [target1, target2, ... targetN]
type ResourceGraph struct {
	edges map[string][]string
}
```
The Graph should have methods for:
- Adding links (source -> target) to the graph
- Removing links from the graph
- Detecting cycles in the graph

Write a unit test which tests these methods

## Exercise 7.2 - Resource Manager

Write a resource manager to block clients if a resource is in use. The manager should avoid deadlocks by 
looking for cycles in the RAG. 

```go
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
```
Write the Acquire() and Release() methods.

## Exercise 7.3 - Dining Philosophers with an the Resource Manager

Use the master solution for the Dining Philosophers problem und remove the channels. Try the ResourceManager to control the resource allocation.
```go
        // take forks
        gotForks := false
        for !gotForks {
            gotForks = manager.Acquire("P" + p.id, "F" + p.id)
            if gotForks {
                gotForks = manager.Acquire(ph, f2)
                if !gotForks { // deadlock detected
                    manager.Release("P" + p.id, "F" + ((p.id + 1) % COUNT)
                }
            } else {
                log.Println("Deadlock detected -> try again")
            }
        }
        
        // put forks
        manager.Release("P" + id, "F" + ((id + 1) % COUNT))   
        manager.Release("P" + id, "F" + id)
```

```go
      // put forks
        manager.Release("P" + id, "F" + ((id + 1) % COUNT))   
        manager.Release("P" + id, "F" + id)
```