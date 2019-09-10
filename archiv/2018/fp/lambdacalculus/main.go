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
	// NOT = b => b(FALSE)(TRUE);

	// This is the key: A Recursive function definition for all functions!!!
	type fnf func(fnf) fnf

	ID := func(x fnf) fnf { return x }

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

	// NOT as function: λb.b false true
	Not := func(b fnf) fnf {
		return b(False)(True)
	}

	fmt.Printf("Id = %p\n", ID)
	fmt.Printf("True = %p\n", True)
	fmt.Printf("False = %p\n", False)

	// should print false
	fmt.Printf("True(False)(True) = %p\n", True(False)(True))
	fmt.Printf("Not(True) = %p\n", Not(True))

	// should print true
	fmt.Printf("False(False)(True) = %p\n", False(False)(True))
	fmt.Printf("Not(False) = %p\n", Not(False))

	// debugging functions
	f := func(x fnf) fnf { fmt.Printf("f()\n"); return x }
	g := func(y fnf) fnf { fmt.Printf("g()\n"); return y }

	// select and call first function f(ID)
	False(False)(True)(f)(g)(ID)
	Not(False)(f)(g)(ID)

	// select and call second function g(ID)
	True(False)(True)(f)(g)(ID)
	Not(True)(f)(g)(ID)
}
