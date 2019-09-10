# Exercise 9 - The Raft Consensus Protocol

If you do not finish during the lecture period, please finish it as homework.

### Exercise 9.1 - Read the Raft specification
Read the Raft specification carefully. 
https://raft.github.io/raft.pdf

## Exercise 9.1 - Answer the following questions
- What problem do Raft solve?
- A Raft server is a distributed statemachine. Describe the possible states and the state transistions.
- What is a split vote? How does a Raft implementation deal with that problem?
- What is a strong leader, how does it affects performance?
- What are the remote function calls (RPC) a minimal raft server must implement?
- Is a cluster of two Raft servers possible?
- What is the purpose of the attribute "term" in a Raft cluster?
- What is the purpose of the attribute "votedFor" in a Raft cluster?
- Describe how data is written to the log. When gets the data visible?

## Exercise 9.3 - Implement Raft in Go (optional in your preferred language)
Implement Raft with the following additional requirements:
- A Raft cluster should be local testable with a single unit test
- A Raft cluster should be remote testable in multiple, distributed processes
- The implementation should follow the specification as close as possible
- Use gRPC for Remote Procedure Calls

Hint: A complete implementation is about 2k LoC. You need at least 2 weeks for a complete implementation.

## Exercise 9.3 - Compare your implementation with others
Compare your implementation with Etcd and Hashicorp Raft. 

https://github.com/etcd-io/etcd/tree/master/raft

https://github.com/hashicorp/raft

Answer the following questions:
- How do the implementations differ from the original specification?
- How do these implementations use Go channels and locks?


