// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Println("process id: ", pid)
}
