// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

package index

import (
	"fmt"
	"testing"
)

func TestIndex(t *testing.T) {

	// prepare book
	p1 := MakePage([]string{"A", "A", "B", "C"})
	p2 := MakePage([]string{"A", "C", "D"})
	p3 := MakePage([]string{"A", "B", "D"})
	book := MakeBook([]Page{p1, p2, p3})

	// calculate index
	idx := MakeIndex(book)

	// stringer support => not automated
	fmt.Printf("Index: %v", idx)
}
