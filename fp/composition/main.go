// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package main

import (
	"fmt"
)

func main() {

	f := func(x int) int { return x * x }

	g := func(x int) int { return x + 1 }

	// Functional Composition: (g◦f)(x)
	gf := func(x int) int { return g(f(x)) }

	fmt.Printf("%v\n", gf(2)) // --> 5

	// Generic Composition
	compose := func(g, f func(int) int) func(int) int {
		return func(x int) int {
			return g(f(x))
		}
	}

	square := func(x int) int { return x * x }
	fmt.Printf("%v\n", compose(square, f)(2)) // --> 16
}
