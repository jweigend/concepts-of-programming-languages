// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

package main

import "fmt"

func main() {
	var a, b = 1, 2
	fmt.Printf("Initial : a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("After a,b = b,a : a=%d, b=%d\n", a, b)
	swap0 := func(x, y int) (int, int) {
		return y, x
	}
	a, b = swap0(a, b)
	fmt.Printf("After a,b = swap0(a,b) : a=%d, b=%d\n", a, b)
	swap1(a, b)
	fmt.Printf("After swap1(a,b) : a=%d, b=%d\n", a, b)
	swap2(&a, &b)
	fmt.Printf("After swap2(&a,&b) : a=%d, b=%d\n", a, b)
	pa, pb := &a, &b
	swap3(&pa, &pb)
	fmt.Printf("After swap3(&pa, &pb): a=%d, b=%d, pa=%p, pb, %p\n", a, b, pa, pb)
}

// END0 OMIT

func swap1(x, y int) {
	x, y = y, x
}

// END1 OMIT

func swap2(x *int, y *int) {
	*x, *y = *y, *x
}

// END2 OMIT

func swap3(x **int, y **int) {
	*x, *y = *y, *x
}

// END3 OMIT
