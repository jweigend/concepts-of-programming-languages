// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

package palindrome

import "testing"

//START OMIT
func TestPalindrome(t *testing.T) {
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
	if !IsPalindrome2("☯otto☯") == true {
		t.Error("isPalindrome('☯otto☯' should be true. But is false.")
	}
	if !IsPalindrome2("") == true {
		t.Error("isPalindrome(Empty string should be a palindrome. But is not. Method returns false.")
	}
	if !IsPalindrome2("o") == true {
		t.Error("isPalindrome('o' should be true. But is false.")
	}
	if !IsPalindrome2("oto") == true {
		t.Error("isPalindrome('oto' should be true. But is false.")
	}
	if !IsPalindrome2("ottos") == false {
		t.Error("isPalindrome('ottos' should be false. But is true.")
	}
}
