// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package main

import (
	"fmt"
)

func main() {

	// Lambda Calculus in Golang --> See Video Graham Hutton
	// https://www.youtube.com/watch?v=eis11j_iGMs
	// Lambda Calculus with JS
	// TRUE = x => y => x;
	// FALSE = x => y => y;
	// NOT = f => x => y => f(y)(x);

	// Some sample functions - strings are only used for debugging purposes

	type fnf func(fnf) fnf // recursive type definition

	f := func(x fnf) fnf { fmt.Printf("f()=%v\n", x); return x }

	g := func(y fnf) fnf { fmt.Printf("g()=%v\n", y); return y }

	id := func(z fnf) fnf { return z } // print

	// TRUE as function: λx.λy.x
	True := func(x fnf) fnf {
		return func(y fnf) fnf {
			return x
		}
	}
	// FALSE as function: λx.λy.y
	False := func(x fnf) fnf {
		return func(y fnf) fnf {
			return y
		}
	}
	// FALSE as function: λx.λy.y
	Not := func(b fnf) fnf {
		return b(False)(True)
	}

	fmt.Printf("Id = %p\n", id)
	fmt.Printf("True = %p\n", True)
	fmt.Printf("False = %p\n", False)

	// false
	fmt.Printf("True(False)(True) = %p\n", True(False)(True))
	fmt.Printf("Not(True) = %p\n", Not(True))

	// true
	fmt.Printf("False(False)(True) = %p\n", False(False)(True))
	fmt.Printf("Not(False) = %p\n", Not(False))

	// select f
	False(False)(True)(f)(g)(id)
	Not(False)(f)(g)(id)

	// select g
	True(False)(True)(f)(g)(id)
	Not(True)(f)(g)(id)
}
