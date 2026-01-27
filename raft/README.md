# Raft

## Purpose

This example implements the **Raft consensus algorithm** using the etcd/raft library. Raft is used for distributed consensus in systems like etcd and Consul.

## Key Concepts

- **Leader election** - One node becomes leader
- **Log replication** - Leader replicates entries to followers
- **Consensus** - Majority agreement on state
- **Fault tolerance** - Survives minority failures

## How Raft Works

1. **Leader Election**: Nodes vote to elect a leader
2. **Log Replication**: Leader accepts writes, replicates to followers
3. **Commit**: Entry committed when majority acknowledges
4. **Consistency**: All nodes eventually have same log

## Components

| Component | Role |
|-----------|------|
| `node` | Raft state machine instance |
| `raftNetwork` | Simulated network transport |
| `Peer` | Node identifier |

## Example Flow

```go
// Create 3-node cluster
peers := []raft.Peer{{ID: 1}, {ID: 2}, {ID: 3}}

// Start nodes
nodes := startNodes(peers)

// Wait for leader election
waitLeader(nodes)

// Propose data
nodes[0].Propose(ctx, []byte("data"))
```

## Run

```bash
go run raft.go
```

## Test

⚠️ Tests are excluded from `make all` due to longer execution time.

```bash
go test -v  # Run explicitly if needed
```
