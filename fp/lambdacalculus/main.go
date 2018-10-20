// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package main

import (
	"fmt"
)

func main() {

	// Lambda Calculus in Golang --> See Video Graham Hutton
	// https://www.youtube.com/watch?v=eis11j_iGMs

	// Some sample functions - strings are only used for debugging purposes
	f := func(x int) string { return fmt.Sprintf("f(%v)=%v", x, x+10) }
	g := func(y int) string { return fmt.Sprintf("g(%v)=%v", y, y+20) }

	// TRUE as function: λx.λy.x
	TRUE := func(x, y func(int) string) func(int) string { return x }

	// FALSE as function: λx.λy.y
	FALSE := func(x, y func(int) string) func(int) string { return y }

	fmt.Println(TRUE(f, g)(1))
	fmt.Println(FALSE(f, g)(1))

	// NOT TRUE
	fmt.Println(TRUE(FALSE(f, g), TRUE(f, g))(2))

	// NOT FALSE
	fmt.Println(TRUE(TRUE(f, g), FALSE(f, g))(2))

	// NOT as function λb.b
	NOT := func(b func(f, g func(int) string) func(int) string) func(x, y func(int) string) func(int) string {
		return func(f, g func(int) string) func(int) string {
			return b(FALSE(f, g), TRUE(f, g))
		}
	}

	fmt.Println(NOT(TRUE)(f, g)(3))
	fmt.Println(NOT(FALSE)(f, g)(3))

}
