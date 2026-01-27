// Copyright 2018-2026 Johannes Weigend
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package stack contains LIFO functions.
package stack

import "errors"

// Stack is a generic LIFO container using Go 1.18+ generics.
// Type parameter T can be any type.
type Stack[T any] struct {
	data []T
}

// NewStack constructs an empty stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push pushes a value on the stack.
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop pops a value from the stack. Returns error if empty (idiomatic Go).
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.data) == 0 {
		return zero, errors.New("cannot pop: empty stack")
	}
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return result, nil
}

// Size returns the count of elements in the Stack
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// Get returns the n-th element in the Stack
func (s *Stack[T]) Get(idx int) T {
	return s.data[idx]
}

// END OMIT
