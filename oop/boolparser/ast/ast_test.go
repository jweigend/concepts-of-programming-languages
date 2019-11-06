// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package ast

import (
	"testing"
)

func TestAST(t *testing.T) {

	// A AND B OR C
	ast := Or{And{Val{"A"}, Val{"B"}}, Val{"C"}}

	truthTable := [][]bool{
		{false, false, false, false},
		{false, false, true, true},
		{false, true, true, true},
		{false, true, false, false},
		{true, false, false, false},
		{true, true, false, true},
		{true, false, true, true},
		{true, true, true, true},
	}

	for i, tt := range truthTable {
		
		vars := map[string]bool{"A": tt[0], "B": tt[1], "C": tt[2]}
		result := ast.Eval(vars)
		
		expected := truthTable[i][3]
		if result != expected {
			t.Errorf("Expected %v but got %v. (Expression := %v, Vars := %v)", expected, result, ast, vars)
		}
	}
}
