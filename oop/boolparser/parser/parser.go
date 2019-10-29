// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package parser contains an parser/interpreter for boolean expressions.
package parser

import (
	"fmt"
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
	rootNode node
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
// ----------  AST  ------------
//

// Abstract syntax tree (AST) node types: and, or, not, value.
type node interface {
	// Eval evaluates the AST. The variables of the expression are set to true or false in the vars map.
	// Missing vars (there are no key in the map) are evaluated to false.
	Eval(vars map[string]bool) bool
}

// or
type or struct {
	lhs node
	rhs node
}

func (o *or) Eval(vars map[string]bool) bool {
	return o.lhs.Eval(vars) || o.rhs.Eval(vars)
}

func (o *or) String() string {
	return fmt.Sprintf("|(%v,%v)", o.lhs, o.rhs)
}

// and
type and struct {
	lhs node
	rhs node
}

func (a *and) Eval(vars map[string]bool) bool {
	return a.lhs.Eval(vars) && a.rhs.Eval(vars)
}

func (a *and) String() string {
	return fmt.Sprintf("&(%v,%v)", a.lhs, a.rhs)
}

// not
type not struct {
	ex node
}

func (n *not) Eval(vars map[string]bool) bool {
	return !n.ex.Eval(vars)
}

func (n *not) String() string {
	return fmt.Sprintf("!(%v)", n.ex)
}

// val
type val struct {
	name string
}

func newVal(name string) *val {
	return &val{name}
}

func (v *val) Eval(vars map[string]bool) bool {
	return vars[v.name] // missing vars will be evaluated to false!
}

func (v *val) String() string {
	return fmt.Sprintf("'%v'", v.name)
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
		p.rootNode = &or{lhs, rhs}
	}
}

// see BNF term.
func (p *Parser) term() {
	p.factor()
	for p.token == "&" {
		lhs := p.rootNode
		p.factor()
		rhs := p.rootNode
		p.rootNode = &and{lhs, rhs}
	}
}

// see BNF factor.
func (p *Parser) factor() {
	p.token = p.lexer.NextToken()
	if p.token == "" {
		return // end
	} else if p.token == "!" {
		p.factor()
		p.rootNode = &not{p.rootNode}
	} else if p.token == "(" {
		p.expression()
		p.token = p.lexer.NextToken()
	} else if isVar(p.token) {
		p.rootNode = newVal(p.token)
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
