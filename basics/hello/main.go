// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package main

import "fmt"

func main() {
	fmt.Printf("Hello %s", "Programming with Go \xE2\x98\xAF\n") // \xE2\x98\xAF -> ☯
	fmt.Printf("Hello %s", "Programming with Go ☯\n")
}
