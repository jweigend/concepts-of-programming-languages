## *Concurrency with Go - Resource Manager*

ResourceManager implements a deadlock avoidance strategy by finding cycles in a Ressource Allocation Graph.
It blocks requests, when a resource is in use. The user is responsible to implement a retry strategy when a deadlock is recognized.

The Resource Manager has two operations:

1. Acquire(processName, resourceName) bool
2. Release(processName, resourceName) 

If a resource is already in use, the Acquire() method blocks until the resource is free. 
The Acquire() method returns true if the resource could be claimed. In case of a deadlock the method returns false. 
It is up to the caller to retry the operation, if the resource is in use and the claim will produce a deadlock operation. 

### Sample

```
Acquire ("P1", "R1" ) : P1 <- R1 (ok) - Process 1 got Resource 1
Acquire ("P1", "R3" ) : P1 <- R3 (ok) - Process 1 got Resource 3
Acquire ("P2", "R2" ) : P2 <- R2 (ok) - Process 2 got Resource 2
Acquire ("P2", "R1" ) : P2 -> R1 (wait) - Process 2 cant get Resource 1 (in use by Process 1) : wait
Acquire ("P1", "R2" ) returns false : P1 -> R2 (deadlock) - acquire will recognize the deadlock and returns false
```

### Dining Philosophers
```go
        // take forks
        for !gotForks {
            gotForks := manager.Acquire("P" + id, "F" + id))
            if gotForks {
                gotForks = manager.Acquire("P" + id, "F" + (id + 1) % COUNT)
                if !gotForks { // deadlock detected
                    m.Release("P" + id, "F" + id))
                }
            } else {
                // deadlock detected -> try again
            }
        }
        
        // eat
        
        // put forks
        manager.Release("P" + id, "F" + ((id + 1) % COUNT))   
        manager.Release("P" + id, "F" + id)
```
