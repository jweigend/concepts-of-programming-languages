// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package lexer contains an interpreter (parser/compiler) for boolean expressions with variables.
package lexer

// Lexer is a simple tokenizer for generating tokens for the boolean parser.
type Lexer struct {
	input   string
	tokens  []string
	current int
}

// NewLexer constructs a lexer from a given string.
func NewLexer(input string) *Lexer {
	lexer := new(Lexer)
	lexer.input = input
	lexer.current = 0
	lexer.tokens = splitTokens(input)
	return lexer
}

// NextToken returns the next token. A token is a non empty string. The function returns "" if there is no token available.
func (l *Lexer) NextToken() string {
	if l.current == len(l.tokens) {
		return ""
	}
	token := l.tokens[l.current]
	l.current++
	return token
}

func splitTokens(input string) []string {
	result := make([]string, 0)
	token := ""
	for i := 0; i < len(input); i++ {
		currentChar := input[i]
		if currentChar == byte(' ') {
			continue // ignore whitespace
		}
		// START OMIT
		switch currentChar {
		// check for terminal
		case byte('&'), byte('|'), byte('!'), byte('('), byte(')'):
			if token != "" {
				result = append(result, token)
				token = ""
			}
			result = append(result, string(currentChar))
			break
		// var assumed
		default:
			token += string(currentChar) // concat var chars
		}
		// END OMIT
	}

	// append last token if exists
	if token != "" {
		result = append(result, token)
	}
	return result
}
