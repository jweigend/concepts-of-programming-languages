package raft

// ReplicatedLog is the transactional log for RAFT.
type ReplicatedLog struct {
}

// NewReplicatedLog constructs a ReplicatedLog.
func NewReplicatedLog() *ReplicatedLog {
	return new(ReplicatedLog)
}
