package raft

type ReplicatedLog struct {
}

func NewReplicatedLog() *ReplicatedLog {
	return new(ReplicatedLog)
}
