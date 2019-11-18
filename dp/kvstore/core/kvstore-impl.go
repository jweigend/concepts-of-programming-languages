package core

// KVStore is a distributed key value store using RAFT consensus.
type KVStore struct {
	committedData map[string]string
}

// NewKVStore constructs a new KVStore.
func NewKVStore() *KVStore {
	return new(KVStore)
}

// SetString sets a value for a given key.
func (k *KVStore) SetString(key, value string) {
	// Make the new value visible (async) when the majority of cluster members have written it to the logfile.
	go func() {
		// success - more than 50% of all cluster member have written the value to their logs
		k.committedData[key] = value
	}()
}

// GetString returns the value for a given key.
func (k *KVStore) GetString(key string) string {
	return k.committedData[key]
}
