// Copyright 2018-2026 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package stack

import (
	"testing"

	"github.com/jweigend/concepts-of-programming-languages/oop/rational"
)

func TestStack(t *testing.T) {
	// String stack using generics
	s := NewStack[string]()

	s.Push("1")
	s.Push("2")
	s.Push("3")

	val, err := s.Pop()
	if err != nil || val != "3" {
		t.Error("Pop() did not return 3")
	}

	val, err = s.Pop()
	if err != nil || val != "2" {
		t.Error("Pop() did not return 2")
	}

	val, err = s.Pop()
	if err != nil || val != "1" {
		t.Error("Pop() did not return 1")
	}

	// Test error handling for empty stack (idiomatic Go)
	_, err = s.Pop()
	if err == nil {
		t.Error("Pop() on empty stack should return error")
	}
}

func TestStackWithRational(t *testing.T) {
	// Rational stack using generics
	s := NewStack[rational.Rational]()

	r1 := rational.NewRational(1, 2)
	r2 := rational.NewRational(2, 4)

	s.Push(r1)
	s.Push(r2)

	val, err := s.Pop()
	if err != nil || val != r2 {
		t.Error("Pop() did not return r2")
	}
}

// TestNoTypeCasting demonstrates that generics eliminate the need for type assertions
func TestNoTypeCasting(t *testing.T) {
	// With generics, no type casting needed!
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)

	sum := 0
	for i := 0; i < s.Size(); i++ {
		sum += s.Get(i) // No .(int) needed - already typed!
	}

	if sum != 3 {
		t.Errorf("Sum result should be 3 but is %v", sum)
	}
}

// END OMIT
