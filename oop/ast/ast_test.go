// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package ast

import (
	"fmt"
	"testing"
)

func TestAST(t *testing.T) {
	ast := Or{And{Val{"A"}, Val{"B"}}, Val{"C"}}

	vars := map[string]bool{"A": true, "B": false, "C": true}
	result := ast.Eval(vars)
	fmt.Printf("expression := %v\n", ast)
	fmt.Printf("vars := %v\n", vars)
	fmt.Printf("result := %v\n", result)

	vars = map[string]bool{"A": true, "B": true, "C": false}
	result = ast.Eval(vars)
	fmt.Printf("expression := %v\n", ast)
	fmt.Printf("vars := %v\n", vars)
	fmt.Printf("result := %v\n", result)

}
