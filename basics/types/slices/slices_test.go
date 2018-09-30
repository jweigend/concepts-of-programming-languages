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
