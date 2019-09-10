// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package util contains a simple Service Locator.
package util

// Registry is a simple Service Locator for interfaces.
type Registry struct {
	services map[string]interface{}
}

// NewRegistry constructor
func NewRegistry() *Registry {
	registry := new(Registry)
	registry.services = make(map[string]interface{})
	return registry
}

// Register registers a single interface for a unique name
func (s *Registry) Register(name string, some interface{}) {
	s.services[name] = some
}

// Get returns the registered interface for a given name
func (s *Registry) Get(name string) interface{} {
	result := s.services[name]
	return result
}
