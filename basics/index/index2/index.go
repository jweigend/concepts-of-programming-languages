// Package index2 makes a book index out of pages.
// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package index2

import (
	"fmt"
	"sort"
)

// Page contains an array of words.
type Page []string

// Book is an array of pages.
type Book []Page

// Set is the set of book page numbers for a given word to avoid duplicates (see index package)
type Set map[int]bool

// Index contains a list of pages for each word in a book.
type Index map[string]Set

// MakeIndex generates an index structure
func MakeIndex(book Book) Index {
	idx := make(Index)
	for i, page := range book {
		for _, word := range page {
			if idx[word] == nil {
				idx[word] = make(Set)
			}
			idx[word][i] = true
		}
	}
	return idx
}

// Stringer support
func (idx Index) String() string {
	result := ""
	for k, v := range idx {
		// To store the keys in slice in sorted order
		var keys []int
		for k := range v {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		result += fmt.Sprintf("\n\tWord: %v : Pages: %v", k, keys)
	}
	return result + "\n"
}

// MakePage constructs a page from a string array.
func MakePage(words []string) Page {
	page := new(Page)
	*page = words
	return *page
}

// MakeBook constructs a book from a page array.
func MakeBook(pages []Page) Book {
	book := new(Book)
	*book = pages
	return *book
}
