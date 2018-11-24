package raft

// Cluster knows the RPC interface of all members.
// The cluster works with multiple nodes for testing or with RPC proxies for remote access.
type Cluster struct {
	allNodes []NodeRPC
}

// NewCluster constructs a new cluster with all Node RPC interfaces.
// For local test the NodeRPC is the Node itself. For distributed operation the NodeRPC is a proxy.
func NewCluster(allNodes []NodeRPC) *Cluster {
	return &Cluster{allNodes}
}

// GetFollowers returns the RPC interfaces of all followers for a given leader
func (c *Cluster) GetFollowers(leaderID int) []NodeRPC {
	// swap leader with last element
	c.allNodes[leaderID], c.allNodes[len(c.allNodes)-1] = c.allNodes[len(c.allNodes)-1], c.allNodes[leaderID]
	return c.allNodes[:len(c.allNodes)-1]
}
