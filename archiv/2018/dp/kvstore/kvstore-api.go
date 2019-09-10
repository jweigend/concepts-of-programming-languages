// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package kvstore contains the KeyValueStore API.
package kvstore

// KVStore is the business interface to a distributed key value store.
type KVStore interface {

	// Sets a value for a given key.
	SetString(key, value string)

	// GetString returns the value for a given key.
	GetString(key string) string
}
