// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package ast contains an abstract syntax tree for boolean expressions.
package ast

import "fmt"

//
// ----------  AST  ------------
//

// Node is the interface to eval an abstract syntax tree (AST)
type Node interface {

	// Eval evaluates the AST. The variables of the expression are set to true or false in the vars map.
	// Missing vars (there are no key in the map) are evaluated to false.
	Eval(vars map[string]bool) bool
}

// Or is the logical OR Operator in an AST
type Or struct {
	lhs Node
	rhs Node
}

// Eval implements the Node interface
func (o Or) Eval(vars map[string]bool) bool {
	return o.lhs.Eval(vars) || o.rhs.Eval(vars)
}

func (o Or) String() string {
	return fmt.Sprintf("|(%v,%v)", o.lhs, o.rhs)
}

// And is the logical AND Operator in an AST
type And struct {
	lhs Node
	rhs Node
}

// Eval implements the Node interface
func (a And) Eval(vars map[string]bool) bool {
	return a.lhs.Eval(vars) && a.rhs.Eval(vars)
}

func (a And) String() string {
	return fmt.Sprintf("&(%v,%v)", a.lhs, a.rhs)
}

// Not is the NOT operator in the AST
type Not struct {
	ex Node
}

// Eval implements the Node interface
func (n Not) Eval(vars map[string]bool) bool {
	return !n.ex.Eval(vars)
}

func (n Not) String() string {
	return fmt.Sprintf("!(%v)", n.ex)
}

// Val is a boolean variable in an AST
type Val struct {
	name string
}

// Eval implements the Node interface
func (v Val) Eval(vars map[string]bool) bool {
	return vars[v.name] // missing vars will be evaluated to false!
}

func (v Val) String() string {
	return fmt.Sprintf("'%v'", v.name)
}
