// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package strings contains string utility functions.
package strings

// Reverse reverses an unicode string.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// End OMIT
