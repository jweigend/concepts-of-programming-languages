// Package index makes a book index out of pages.
// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package index

import "fmt"

// Page contains an array of words.
type Page []string

// Book is an array of pages.
type Book []Page

// Index contains a list of pages for each word in a book.
type Index map[string][]int

// MakeIndex generates an index structure
func MakeIndex(book Book) Index {
	idx := make(Index)
	for i, page := range book {
		for _, word := range page {
			pages := idx[word]
			idx[word] = append(pages, i)
		}
	}
	return idx
}

// Stringer support
func (idx Index) String() string {
	result := ""
	for k, v := range idx {
		result += fmt.Sprintf("\n\tWord: %v : Pages: %v", k, v)
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
