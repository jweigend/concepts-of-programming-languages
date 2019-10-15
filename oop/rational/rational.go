// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package rational implements rational numbers.
package rational

import "fmt"

// Rational represents a rational number numerator/denominator.
type Rational struct {
	numerator   int
	denominator int
}

// END1 OMIT

// NewRational constructor function
func NewRational(numerator int, denominator int) Rational {
	if denominator == 0 {
		panic("division by zero")
	}
	r := Rational{}
	divisor := gcd(numerator, denominator)
	r.numerator = numerator / divisor
	r.denominator = denominator / divisor
	return r
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
