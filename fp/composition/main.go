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
	type any interface{}
	type function func(any) any

	compose := func(g, f function) function {
		return func(x any) any {
			return g(f(x))
		}
	}
	square := func(x any) any { return x.(int) * x.(int) }
	f2 := func(x any) any { return f(x.(int)) }

	fmt.Printf("%v\n", compose(square, f2)(2))              // --> 16
	fmt.Printf("%v\n", compose(compose(square, f2), f2)(2)) // --> 256
}
