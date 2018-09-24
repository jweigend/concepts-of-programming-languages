// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package main

import (
	"flag"
	"fmt"
)

func main() {
	ip := flag.String("ip", "192.168.1.1", "Overrides the default IP address.")
	port := flag.String("port", "8080", "Overrides the default listen port.")

	if len(flag.Args()) == 0 {
		fmt.Printf("Program Usage:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Println("\nDefault value for IP: " + *ip)
	fmt.Println("\nDefault value for port: " + *port)
}
