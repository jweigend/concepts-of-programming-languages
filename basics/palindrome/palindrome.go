// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package palindrome implements multiple functions for palindromes.
package palindrome

// IsPalindrome implementation
func IsPalindrome(word string) bool {
	for pos := 0; pos < len(word)/2; pos++ {
		if word[pos] != word[len(word)-pos-1] {
			return false
		}
	}
	return true
}

// IsPalindrome2 is implemented by using runes.
func IsPalindrome2(word string) bool {
	var runes = []rune(word)
	for pos, ch := range runes {
		if ch != runes[len(runes)-pos-1] {
			return false
		}
	}
	return true
}
