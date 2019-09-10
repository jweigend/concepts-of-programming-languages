// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package idserv contains the IDService API.
package idserv

// IDService can be used to produce network wide unique ids.
type IDService interface {

	// NewUUID generates an UUID with a given client prefix.
	NewUUID(clientID string) string
}
