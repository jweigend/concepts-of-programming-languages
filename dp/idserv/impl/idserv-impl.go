// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package impl contains the business logic of the IDService
package impl

import (
	"fmt"
	"sync/atomic"
)

// IDServiceImpl type
type IDServiceImpl struct {
}

// The last given Id.
var lastID int64

// NewIDServiceImpl creates a new instance
func NewIDServiceImpl() *IDServiceImpl {
	return new(IDServiceImpl)
}

// NewUUID implements the IDService interface.
func (ids *IDServiceImpl) NewUUID(clientID string) string {
	result := atomic.AddInt64(&lastID, 1)
	return fmt.Sprintf("%v:%v", clientID, result)
}
