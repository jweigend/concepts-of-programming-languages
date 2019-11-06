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

	// λx.x is a function which returns itself (the ID)
	ID := func(x fnf) fnf { return x }

	// Functional Numbers ONE
	ONCE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(x)
		}
	}

	// Functional Numbers TWO
	TWICE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(f(x))
		}
	}

	// Function Numbers THREE
	THRICE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(f(f(x)))
		}
	}

	// Functional Numbers SUCCESSOR(N) => λwyx.y(wyx)
	SUCCESSOR := func(w fnf) fnf {
		return func(y fnf) fnf {
			return func(x fnf) fnf {
				return y(w)(y)(x)
			}
		}
	}

	Printer := func(x fnf) fnf { fmt.Print("."); return x }
	QUAD := TWICE(TWICE)
	QUAD(Printer)(ID)
	fmt.Println()

	SUCCESSOR(TWICE)(Printer)(ID)
	fmt.Println("SUCCESSOR(TWICE) = 3")

	SUCCESSOR(THRICE)(Printer)(ID)
	fmt.Println("SUCCESSOR(THRICE) = 4")

	// Boolean TRUE as function: λx.λy.x
	TRUE := func(x fnf) fnf {
		return func(y fnf) fnf {
			return x
		}
	}

	// Boolean FALSE as function: λx.λy.y
	FALSE := func(x fnf) fnf {
		return func(y fnf) fnf {
			return y
		}
	}

	// NOT as function: λb.b False True
	NOT := func(b fnf) fnf {
		return b(FALSE)(TRUE)
	}

	// AND as function: λxy.xy False
	AND := func(x fnf, y fnf) fnf {
		return x(y)(FALSE)
	}

	fmt.Printf("AND(true, true) = %p\n", AND(TRUE, TRUE))
	fmt.Printf("AND(true, false) = %p\n", AND(TRUE, FALSE))
	fmt.Printf("AND(false, true) = %p\n", AND(FALSE, TRUE))

	fmt.Printf("Id = %p\n", ID)
	fmt.Printf("True = %p\n", TRUE)
	fmt.Printf("False = %p\n", FALSE)

	// should print false
	fmt.Printf("True(False)(True) = %p\n", TRUE(FALSE)(TRUE))
	fmt.Printf("NOT(True) = %p\n", NOT(TRUE))

	// should print true
	fmt.Printf("False(False)(True) = %p\n", FALSE(FALSE)(TRUE))
	fmt.Printf("NOT(False) = %p\n", NOT(FALSE))

	// debugging functions
	f := func(x fnf) fnf { fmt.Printf("f()\n"); return x }
	g := func(y fnf) fnf { fmt.Printf("g()\n"); return y }

	// select AND call first function f(ID)
	FALSE(FALSE)(TRUE)(f)(g)(ID)
	NOT(FALSE)(f)(g)(ID)

	// select AND call second function g(ID)
	TRUE(FALSE)(TRUE)(f)(g)(ID)
	NOT(TRUE)(f)(g)(ID)

	ONCE(NOT)(TRUE)   // -> false
	TWICE(NOT)(TRUE)  // -> true
	THRICE(NOT)(TRUE) // -> false

	fmt.Printf("ONCE(NOT)(TRUE) = %p\n", ONCE(NOT)(TRUE))
	fmt.Printf("TWICE(NOT)(TRUE) = %p\n", TWICE(NOT)(TRUE))
}
