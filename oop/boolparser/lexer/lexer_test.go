// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package lexer

import (
	"fmt"
	"testing"
)

func TestNewLexer(t *testing.T) {

	lexer := NewLexer("a & b")
	if len(lexer.tokens) != 3 {
		t.Error(fmt.Sprintf("Wrong token count! Expected: 3, Got: %v", len(lexer.tokens)))
	}

	lexer = NewLexer("a|b")
	if len(lexer.tokens) != 3 {
		t.Error(fmt.Sprintf("Wrong token count! Expected: 3, Got: %v", len(lexer.tokens)))
	}

	lexer = NewLexer("a|b&(b|c)")
	if len(lexer.tokens) != 9 {
		t.Error(fmt.Sprintf("Wrong token count! Expected: 9, Got: %v", len(lexer.tokens)))
	}

	tok := lexer.NextToken()
	for tok != "" {
		fmt.Print(tok)
		tok = lexer.NextToken()
	}
}
