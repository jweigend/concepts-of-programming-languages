// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

package palindrome

import "testing"

//START OMIT
func TestPalindrome1(t *testing.T) {
	// Unicode Character do not work with a simple string implementation
	if !IsPalindrome("") == true {
		t.Error("isPalindrome('' should be true. But is false.")
	}
	if !IsPalindrome("o") == true {
		t.Error("isPalindrome('o' should be true. But is false.")
	}
	if !IsPalindrome("oto") == true {
		t.Error("isPalindrome('oto' should be true. But is false.")
	}
	if !IsPalindrome("ottos") == false {
		t.Error("isPalindrome('ottos' should be false. But is true.")
	}
	//END OMIT
	if IsPalindrome("☯otto☯") == true {
		t.Error("isPalindrome('☯otto☯' does not work with the simple string implementation!")
	}
}

func TestPalindrome2(t *testing.T) {
	testPalindromeUTF8(t, IsPalindrome2)
}

func TestPalindrome3(t *testing.T) {
	testPalindromeUTF8(t, IsPalindrome3)
}

func testPalindromeUTF8(t *testing.T, isPalindrome func(word string) bool) {
	if !isPalindrome("☯otto☯") == true {
		t.Error("isPalindrome('☯otto☯' should be true. But is false.")
	}
	if !isPalindrome("") == true {
		t.Error("isPalindrome(Empty string should be a palindrome. But is not. Method returns false.")
	}
	if !isPalindrome("o") == true {
		t.Error("isPalindrome('o' should be true. But is false.")
	}
	if !isPalindrome("oto") == true {
		t.Error("isPalindrome('oto' should be true. But is false.")
	}
	if !isPalindrome("ottos") == false {
		t.Error("isPalindrome('ottos' should be false. But is true.")
	}
}
