// Copyright 2018-2026 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package streams contains a minimal Java8 Streams like API for slices using Go 1.18+ generics.
// Usage:
//
//	result := Of(slice).
//		Map(toUpperCase).
//		Filter(notDigit).
//		Reduce(concat)
//
//	- Of() converts a slice of type T to a Stream[T].
//	- Map() applies a function to all elements (same type).
//	- Filter() filters elements based on a predicate.
//	- Reduce() combines all elements and returns a single element.
//
// For type-changing transformations, use MapTo[T, R]() function.
// This is a very first draft with no lazy or parallel support.
package streams

// Predicate function returns true if an element should be kept.
type Predicate[T any] func(T) bool

// Mapper function maps a value to another value of the same type.
type Mapper[T any] func(T) T

// TransformMapper maps a value of type T to a value of type R.
type TransformMapper[T, R any] func(T) R

// Accumulator function combines two elements into one.
type Accumulator[T any] func(T, T) T

// Pair of two values with generic types.
type Pair[K, V any] struct {
	Key   K
	Value V
}

// Iterator interface for iterating over stream elements.
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

// Stream interface is implemented for container types.
type Stream[T any] interface {
	Iterator() Iterator[T]
	Map(m Mapper[T]) Stream[T]
	Filter(p Predicate[T]) Stream[T]
	Reduce(a Accumulator[T]) T
	Collect() []T
}

// SliceIterator iterates over a slice.
type SliceIterator[T any] struct {
	slice      []T
	currentPos int
}

// NewSliceIterator creates a new iterator for a slice.
func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{slice: slice}
}

// HasNext returns true if there are more elements.
func (s *SliceIterator[T]) HasNext() bool {
	return len(s.slice) > s.currentPos
}

// Next returns the next element or panics if none available.
func (s *SliceIterator[T]) Next() T {
	if len(s.slice) > s.currentPos {
		cp := s.currentPos
		s.currentPos++
		return s.slice[cp]
	}
	var zero T
	panic("No such element")
	return zero
}

// Of creates a new Stream from a slice. This is the main entry point.
func Of[T any](s []T) Stream[T] {
	return NewSliceStream(s)
}

// SliceStream is a stream implementation for slices.
type SliceStream[T any] struct {
	data []T
}

// NewSliceStream returns a new stream.
func NewSliceStream[T any](data []T) *SliceStream[T] {
	return &SliceStream[T]{data: data}
}

// Map applies the Mapper on all elements.
func (s *SliceStream[T]) Map(mapper Mapper[T]) Stream[T] {
	for i, e := range s.data {
		s.data[i] = mapper(e)
	}
	return s
}

// Filter keeps only elements matching the predicate.
func (s *SliceStream[T]) Filter(predicate Predicate[T]) Stream[T] {
	var filtered []T
	for _, e := range s.data {
		if predicate(e) {
			filtered = append(filtered, e)
		}
	}
	s.data = filtered
	return s
}

// Reduce combines all elements using the accumulator.
// Returns the zero value of T if the stream is empty.
func (s *SliceStream[T]) Reduce(accumulate Accumulator[T]) T {
	var result T
	for i, e := range s.data {
		if i == 0 {
			result = e
		} else {
			result = accumulate(result, e)
		}
	}
	return result
}

// Collect returns the stream elements as a slice.
func (s *SliceStream[T]) Collect() []T {
	return s.data
}

// Iterator returns an iterator over the stream elements.
func (s *SliceStream[T]) Iterator() Iterator[T] {
	return NewSliceIterator(s.data)
}

// ============================================================================
// Type-transforming operations (standalone functions due to Go generics limitations)
// ============================================================================

// MapTo transforms a Stream[T] to a Stream[R] using a transformation function.
// This is a standalone function because Go doesn't support method-level type parameters.
func MapTo[T, R any](stream Stream[T], mapper TransformMapper[T, R]) Stream[R] {
	input := stream.Collect()
	output := make([]R, len(input))
	for i, e := range input {
		output[i] = mapper(e)
	}
	return NewSliceStream(output)
}

// FlatMap transforms each element to a slice and flattens the result.
func FlatMap[T, R any](stream Stream[T], mapper TransformMapper[T, []R]) Stream[R] {
	input := stream.Collect()
	var output []R
	for _, e := range input {
		output = append(output, mapper(e)...)
	}
	return NewSliceStream(output)
}

// ReduceTo reduces a Stream[T] to a single value of type R.
func ReduceTo[T, R any](stream Stream[T], initial R, accumulate func(R, T) R) R {
	result := initial
	it := stream.Iterator()
	for it.HasNext() {
		result = accumulate(result, it.Next())
	}
	return result
}

// ============================================================================
// Legacy API (deprecated - for backward compatibility)
// ============================================================================

// Any is deprecated: Use generics instead. Kept for backward compatibility.
type Any = any

// ToStream is deprecated: Use Of() instead. Kept for backward compatibility.
func ToStream(s []Any) Stream[Any] {
	return Of(s)
}
