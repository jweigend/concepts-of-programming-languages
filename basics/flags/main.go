// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package main

import (
	"flag"
	"fmt"
)

// Simple test for the Go flag API.
func main() {
	// construct a string flag with a default ip address and a description.
	ip := flag.String("ip", "192.168.1.1", "Overrides the default IP address.")
	port := flag.String("port", "8080", "Overrides the default listen port.")

	// flag.Args() parses the arg of our program.
	if len(flag.Args()) == 0 {
		fmt.Printf("Program Usage:\n")
		// PrintDefaults() prints a description and the default values to stdout.
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Println("\nDefault value for IP: " + *ip)
	fmt.Println("\nDefault value for port: " + *port)
}

// END OMIT
