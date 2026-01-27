// Copyright 2018-2026 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package streams

import (
	"fmt"
	"strings"
	"testing"
)

// TestMapFilterReduce tests the Java 8 Map/Filter/Reduce Chain with generics.
// stringSlice OMIT
func TestMapFilterReduce(t *testing.T) {
	// Type-safe slice of strings (no more interface{}).
	stringSlice := []string{"a", "b", "c", "1", "D"}

	// Map/Reduce with type-safe generics
	result := Of(stringSlice).
		Map(toUpperCase).
		Filter(notDigit).
		Reduce(concat)

	if result != "A,B,C,D" {
		t.Error(fmt.Sprintf("Result should be 'A,B,C,D' but is: %v", result))
	}
	// EOF OMIT

	// Lambda (inline) - now type-safe, no type assertions needed!
	result = Of(stringSlice).
		Map(func(s string) string {
			return strings.ToUpper(s)
		}).
		Filter(func(s string) bool {
			for _, v := range s {
				if v >= '0' && v <= '9' {
					return false
				}
			}
			return true
		}).
		Reduce(func(a, b string) string {
			return a + "," + b
		})

	if result != "A,B,C,D" {
		t.Errorf("Result should be 'A,B,C,D' but is: %v", result)
	}
}

// toUpperCase converts a given string to upper case.
func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

//EOF OMIT

// notDigit returns true if the string contains no digits.
func notDigit(s string) bool {
	for _, v := range s {
		if v >= '0' && v <= '9' {
			return false
		}
	}
	return true
}

// concat produces a string by concatenating two strings with comma.
func concat(a, b string) string {
	return a + "," + b
}

// WC OMIT
//
// ========================
// Classic wordcount sample
// ========================

func TestWordCount(t *testing.T) {
	words := []string{"a", "a", "b", "b", "D", "a"}

	// Using ReduceTo for type-transforming reduction (string -> map[string]int)
	wordCounts := ReduceTo(Of(words), make(map[string]int), func(counts map[string]int, word string) map[string]int {
		counts[word]++
		return counts
	})

	// Print results
	for k, v := range wordCounts {
		fmt.Printf("%v:%v, ", k, v)
	}
	fmt.Println()

	// Verify counts
	if wordCounts["a"] != 3 {
		t.Errorf("Expected 'a' count to be 3, got %d", wordCounts["a"])
	}
	if wordCounts["b"] != 2 {
		t.Errorf("Expected 'b' count to be 2, got %d", wordCounts["b"])
	}
	if wordCounts["D"] != 1 {
		t.Errorf("Expected 'D' count to be 1, got %d", wordCounts["D"])
	}
}

// ENDWC OMIT

// TestMapTo demonstrates type-changing transformations.
func TestMapTo(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	// Transform int stream to string stream
	stringStream := MapTo(Of(numbers), func(n int) string {
		return fmt.Sprintf("num:%d", n)
	})

	result := stringStream.Collect()
	if len(result) != 5 {
		t.Errorf("Expected 5 elements, got %d", len(result))
	}
	if result[0] != "num:1" {
		t.Errorf("Expected 'num:1', got '%s'", result[0])
	}
}

// TestFlatMap demonstrates flattening transformations.
func TestFlatMap(t *testing.T) {
	words := []string{"hello", "world"}

	// Split each word into characters
	chars := FlatMap(Of(words), func(word string) []rune {
		return []rune(word)
	})

	result := chars.Collect()
	if len(result) != 10 { // "hello" (5) + "world" (5)
		t.Errorf("Expected 10 characters, got %d", len(result))
	}
}

// TestIterator demonstrates the generic iterator.
func TestIterator(t *testing.T) {
	numbers := []int{1, 2, 3}
	stream := Of(numbers)
	it := stream.Iterator()

	sum := 0
	for it.HasNext() {
		sum += it.Next()
	}

	if sum != 6 {
		t.Errorf("Expected sum to be 6, got %d", sum)
	}
}

// TestCollect demonstrates collecting stream to slice.
func TestCollect(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	result := Of(input).
		Filter(func(n int) bool { return n%2 == 0 }).
		Collect()

	if len(result) != 2 || result[0] != 2 || result[1] != 4 {
		t.Errorf("Expected [2, 4], got %v", result)
	}
}
