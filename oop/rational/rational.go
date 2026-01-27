// Copyright 2018-2026 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package rational implements rational numbers.
package rational

import (
	"errors"
	"fmt"
)

// ErrDivisionByZero is returned when denominator is zero.
var ErrDivisionByZero = errors.New("division by zero")

// Rational represents a rational number numerator/denominator.
type Rational struct {
	numerator   int
	denominator int
}

// END1 OMIT

// NewRational constructor function. Returns error if denominator is zero.
// For convenience, a panic version is also provided: MustNewRational.
func NewRational(numerator int, denominator int) Rational {
	r, err := TryNewRational(numerator, denominator)
	if err != nil {
		panic(err) // Kept for backward compatibility with existing tests
	}
	return r
}

// TryNewRational is the idiomatic Go version that returns an error.
func TryNewRational(numerator int, denominator int) (Rational, error) {
	if denominator == 0 {
		return Rational{}, ErrDivisionByZero
	}
	divisor := gcd(numerator, denominator)
	return Rational{
		numerator:   numerator / divisor,
		denominator: denominator / divisor,
	}, nil
}

// END2 OMIT

// Multiply method for rational numbers (x1/x2 * y1/y2)
func (x Rational) Multiply(y Rational) Rational {
	return NewRational(x.numerator*y.numerator, x.denominator*y.denominator)
}

// Multiply OMIT

// Add adds two rational numbers
func (x Rational) Add(y Rational) Rational {
	return NewRational(x.numerator*y.denominator+y.numerator*x.denominator, x.denominator*y.denominator)
}

// Stringer
func (x Rational) String() string {
	return fmt.Sprintf("(%v/%v)", x.numerator, x.denominator)
}

// Stringer OMIT

// Helper GCD -> Euclidean algorithm
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
