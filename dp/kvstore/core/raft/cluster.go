package raft

// Cluster knows the RPC interface of all members.
// The cluster works with multiple nodes for testing or with RPC proxies for remote access.
type Cluster struct {
	members []NodeRPC
}
