// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package slices

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	array := [2]string{"A", "B"}
	slice := array[:] // slice the array
	fmt.Printf("Slice: %v, Capacity: %v, Length: %v\n", slice, cap(slice), len(slice))

	// append() does not work for arrays, only for slices.
	slice = append(slice, "C")
	fmt.Printf("Slice: %v, Capacity: %v, Length: %v\n", slice, cap(slice), len(slice))
	slice = append(slice, "D")
	slice = append(slice, "E")
	fmt.Printf("Slice: %v, Capacity: %v, Length: %v\n", slice, cap(slice), len(slice))
}

func TestSliceCopy(t *testing.T) {
	array := [2]string{"A", "B"}
	slice := array[:1] // slice the array
	fmt.Printf("Array: %p\n", &array)
	fmt.Printf("Slice: %p\n", slice)
	slice1 := slice
	fmt.Printf("Slice 1: %p\n", slice1)

	f := func(s []string) {
		fmt.Printf("Slice 1: %p\n", s)
	}
	f(slice1)

	slice1 = append(slice1, "C")
	fmt.Printf("Slice 1: %p\n", slice1)

	array[0] = "X"
	fmt.Printf("Slice: %s, %p\n", slice, slice)
	fmt.Printf("Array: %s, %p\n", array, &array)

}
