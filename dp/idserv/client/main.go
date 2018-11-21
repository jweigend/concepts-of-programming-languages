// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package main

import (
	"log"

	"github.com/jweigend/concepts-of-programming-languages/dp/idserv"
	"github.com/jweigend/concepts-of-programming-languages/dp/idserv/remote/proxy"
	"github.com/qaware/programmieren-mit-go/04_distributed-programming/idserv/server/impl"
)

// GenerateIds calls n-times NewUUID() in a loop and returns the result as slice.
func GenerateIds(count int, service idserv.IDService) []string {
	result := make([]string, count)
	for i := 0; i < count; i++ {
		result[i] = service.NewUUID("c1")
	}
	return result
}

func main() {
	var service idserv.IDService

	// Local
	service = impl.NewIDServiceImpl()
	result := GenerateIds(10, service)

	log.Printf("Got Id: %v", result)

	// Remote
	service = proxy.NewProxy()
	result = GenerateIds(10, service)

	log.Printf("Got Id: %v", result)
}
