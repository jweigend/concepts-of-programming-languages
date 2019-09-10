// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0
package strings

import (
	"fmt"
	"testing"
)

func TestStringLength(t *testing.T) {

	// Unicode code length
	length := len("A")
	fmt.Printf("Length: %d\n", length)
	length = len("±")
	fmt.Printf("Length: %d\n", length)
	length = len("☯")
	fmt.Printf("Length: %d\n", length)
}

func TestReverse(t *testing.T) {

	s1 := "Hello, world"
	s2 := "Hello, 世界"
	s3 := "The quick bròwn 狐 jumped over the lazy 犬"

	if Reverse(Reverse(s1)) != s1 {
		t.Errorf("Reversed string %s ist not equal to %s", Reverse(Reverse(s1)), s1)
	}
	if Reverse(Reverse(s2)) != s2 {
		t.Errorf("Reversed string %s ist not equal to %s", Reverse(Reverse(s2)), s2)
	}
	if Reverse(Reverse(s3)) != s3 {
		t.Errorf("Reversed string %s ist not equal to %s", Reverse(Reverse(s3)), s3)
	}
}
