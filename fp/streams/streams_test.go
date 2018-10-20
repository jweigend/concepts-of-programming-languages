// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package streams

import (
	"fmt"
	"strings"
	"testing"
)

// TestMapFilterReduce tests the Java 8 Map/Filter/Reduce Chain.
func TestMapFilterReduce(t *testing.T) {
	// array of generic interfaces.
	stringSlice := []Any{"a", "b", "c", "1", "D"}

	// Map/Reduce
	result := ToStream(stringSlice).
		Map(toUpperCase).
		Filter(notDigit).
		Reduce(concat).(string)

	if result != "A,B,C,D" {
		t.Error(fmt.Sprintf("Result should be 'A,B,C,D' but is: %v", result))
	}
	// lambda (inline)
	result = ToStream(stringSlice).
		Map(func(o Any) Any {
			return strings.ToUpper(o.(string))
		}).
		Filter(func(o Any) bool {
			s := o.(string)
			result := true
			for _, v := range s {
				if v >= '0' && v <= '9' {
					result = false
					break
				}
			}
			return result
		}).
		Reduce(func(a Any, b Any) Any {
			return a.(string) + "," + b.(string)
		}).(string)
}

// toUpperCase converts a given string to upper case.
func toUpperCase(o Any) Any {
	return strings.ToUpper(o.(string))
}

//EOF OMIT

// notDigit loops over a string value and checks if the string contains a digit.
func notDigit(o Any) bool {
	s := o.(string)
	result := true
	for _, v := range s {
		if v >= '0' && v <= '9' {
			result = false
			break
		}
	}
	return result
}

// concat produces a string by concating two strings with ,.
func concat(a Any, b Any) Any {
	return a.(string) + "," + b.(string)
}

// WC OMIT
//
// ========================
// Classic wordcount sample
// ========================
//
func ExampleStream() {
	strings := []Any{"a", "a", "b", "b", "D", "a"}

	// Map/Reduce
	result := ToStream(strings).
		Map(func(o Any) Any {
			result := []Pair{Pair{o, 1}}
			return result
		}).
		Reduce(sumInts).([]Pair)

	for _, e := range result {
		fmt.Printf("%v:%v, ", e.k, e.v)
	}
}

// ENDWC OMIT

// sumValues reduces the pair arrays by adding the count for the same key.
func sumInts(a Any, b Any) Any {
	pa := a.([]Pair)
	pb := b.([]Pair)
	for i, e := range pa {
		for _, u := range pb {
			if e.k == u.k {
				pa[i].v = e.v.(int) + u.v.(int)
				return pa
			}
		}
	}
	result := append(pa, pb...)
	return result
}
