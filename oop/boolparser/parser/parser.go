// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package parser contains an parser/interpreter for boolean expressions.
package parser

import (
	"fmt"

	"github.com/jweigend/concepts-of-programming-languages/oop/boolparser/ast"
	"github.com/jweigend/concepts-of-programming-languages/oop/boolparser/lexer"
)

//  ---------------------------------------------------------
// The expression should have the following EBNF form:
//  EBNF
//  ---------------------------------------------------------
// 	<expression> ::= <term> { <or> <term> }
// 	<term> ::= <factor> { <and> <factor> }
// 	<factor> ::= <var> | <not> <factor> | (<expression>)
// 	<or>  ::= '|'
// 	<and> ::= '&'
// 	<not> ::= '!'
//  <var> ::= '[a-zA-Z0-9]*'
//  ---------------------------------------------------------

// Parser is a recursive decent parser for boolean expressions.
type Parser struct {
	rootNode ast.Node
	token    string // ll(1)
	lexer    *lexer.Lexer
}

// NewParser constructs a recursive descent parser and compiles the input of the lexer.
func NewParser(lexer *lexer.Lexer) *Parser {
	b := Parser{lexer: lexer}
	b.parse()
	return &b
}

// Eval evaluates the AST tree against the given var map.
func (p *Parser) Eval(vars map[string]bool) bool {
	return p.rootNode.Eval(vars)
}

// String implements interface Stringer.
func (p *Parser) String() string {
	return fmt.Sprintf("%v", p.rootNode)
}

//
// ----------  PARSING  ------------
//

// parse expr and build the AST.
func (p *Parser) parse() {
	p.expression()
}

// see BNF expression.
func (p *Parser) expression() {
	p.term()
	for p.token == "|" {
		lhs := p.rootNode
		p.term()
		rhs := p.rootNode
		p.rootNode = &ast.Or{LHS: lhs, RHS: rhs}
	}
}

// see BNF term.
func (p *Parser) term() {
	p.factor()
	for p.token == "&" {
		lhs := p.rootNode
		p.factor()
		rhs := p.rootNode
		p.rootNode = &ast.And{LHS: lhs, RHS: rhs}
	}
}

// see BNF factor.
func (p *Parser) factor() {
	p.token = p.lexer.NextToken()
	if p.token == "" {
		return // end
	} else if p.token == "!" {
		p.factor()
		p.rootNode = &ast.Not{Ex: p.rootNode}
	} else if p.token == "(" {
		p.expression()
		p.token = p.lexer.NextToken()
	} else if isVar(p.token) {
		p.rootNode = ast.Val{Name: p.token}
		p.token = p.lexer.NextToken()
	} else {
		panic(fmt.Sprintf("Unknown symbol %v", p.token))
	}
}

// isVar checks if a token is a identifier which starts with an ASCII Letter.
func isVar(token string) bool {
	if len(token) == 0 {
		panic("Empty token!")
	}
	return token[0] >= byte('0') && token[0] <= byte('z')
}
