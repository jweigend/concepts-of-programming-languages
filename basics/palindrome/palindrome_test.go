// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

package palindrome

import "testing"

//START OMIT
// palindrome_test.go
func TestPalindrome(t *testing.T) {
	if !IsPalindrome("") {
		t.Error("isPalindrome('' should be true. But is false.")
	}
	if !IsPalindrome("o") {
		t.Error("isPalindrome('o' should be true. But is false.")
	}
	if !IsPalindrome("oto") {
		t.Error("isPalindrome('oto' should be true. But is false.")
	}
	if IsPalindrome("ottos") {
		t.Error("isPalindrome('ottos' should be false. But is true.")
	}
	//END OMIT
}

func TestPalindrome2(t *testing.T) {
	testPalindromeUTF8(t, IsPalindrome2)
}

func TestPalindrome3(t *testing.T) {
	testPalindromeUTF8(t, IsPalindrome3)
}

func testPalindromeUTF8(t *testing.T, isPalindrome func(word string) bool) {
	if !isPalindrome("☯otto☯") {
		t.Error("isPalindrome('☯otto☯' should be true. But is false.")
	}
	if !isPalindrome("") {
		t.Error("isPalindrome(Empty string should be a palindrome. But is not. Method returns false.")
	}
	if !isPalindrome("o") {
		t.Error("isPalindrome('o' should be true. But is false.")
	}
	if !isPalindrome("oto") {
		t.Error("isPalindrome('oto' should be true. But is false.")
	}
	if isPalindrome("ottos") {
		t.Error("isPalindrome('ottos' should be false. But is true.")
	}
}
