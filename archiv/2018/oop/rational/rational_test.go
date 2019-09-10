// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

package rational

import (
	"fmt"
	"testing"
)

func TestRational(t *testing.T) {

	r1 := NewRational(1, 2)
	r2 := NewRational(2, 4)

	// test equal
	if r1 != r2 {
		t.Error("1/2 should be equal to 2/4 but is not.")
	}

	// test multiply
	r3 := r1.Multiply(r2)
	if r3 != NewRational(1, 4) {
		t.Error(fmt.Sprintf("1/2 * 1/2 should be 1/4 but ist %v", r3))
	}

	// test add
	r4 := r3.Add(r3)
	if r4 != NewRational(1, 2) {
		t.Error(fmt.Sprintf("1/4 + 1/4 should be 1/2 but ist %v", r4))
	}
}

// END OMIT
