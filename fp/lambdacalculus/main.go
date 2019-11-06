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

	// Numbers 
	ONCE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(x)
		}
	}
	
	TWICE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(f(x))
		}
	}

	THRICE := func(f fnf) fnf {
		return func(x fnf) fnf {
			return f(f(f(x)))
		}
	}

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

	ONCE(NOT)(TRUE) // -> false 
	TWICE(NOT)(TRUE) // -> true
	THRICE(NOT)(TRUE) // -> false

	fmt.Printf("ONCE(NOT)(TRUE) = %p\n", ONCE(NOT)(TRUE))
	fmt.Printf("TWICE(NOT)(TRUE) = %p\n", TWICE(NOT)(TRUE))

	// // SUCC = λwyx.y(wyx)
	// SUCC := func (w, y, x fnf) fnf {
	// 	return y(w(y(x)))
	// }

	// fmt.Printf("SUCC(ONCE, NOT, TRUE) = %p\n", SUCC(ONCE, NOT, TRUE))
	// fmt.Printf("SUCC(TWICE, NOT, TRUE) = %p\n", SUCC(TWICE, NOT, TRUE))

}
