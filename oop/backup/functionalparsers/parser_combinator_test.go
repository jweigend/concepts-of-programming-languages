// (c) Armin Heller 2018

package functionalparsers

import (
	"container/list"
	"fmt"
	"testing"
)

func TestExpectString(t *testing.T) {
	var input = inputOfString("ABCD")
	var result = expectString("AB")(input)
	if result.result != "AB" {
		t.Error(fmt.Sprintf("expectString fails, result=%s!", result.result))
	}
	result = expectString("CD")(input)
	if result.result != nil {
		t.Error(fmt.Sprintf("expectString fails, result=%s!", result.result))
	}
}

func TestCombinators(t *testing.T) {
	var p = expectString("0").orElse(expectString("1").andThen(expectString("2")).apply(func(x interface{}) interface{} {
		return x.(ResultPair).first.(string) + x.(ResultPair).second.(string)
	})).anyTimes()
	var i = inputOfString("0120012")
	var res = p(i)
	var expected = []string{"0", "12", "0", "0", "12"}
	if res.result == nil {
		t.Error("Expected the result of the parse to be non-nil!")
	}
	if res.result.(*list.List).Len() != len(expected) {
		t.Error(fmt.Sprintf("Expected %d elements in the result but got %d.", len(expected), res.result.(*list.List).Len()))
	}
	var idx = 0
	var l = res.result.(*list.List)
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value != expected[idx] {
			t.Error(fmt.Sprintf("Element at index %d doesn't match!", idx))
		}
		idx++
	}
}

func TestBoolParser(t *testing.T) {

	q := expression(inputOfString("(a & (b | c & b)) & d")).result.(*and)
	if q.String() != "&(&('a',|('b',&('c','b'))),'d')" {
		t.Error(fmt.Sprintf("Wrong string representation: %v", q.String()))
	}

	res := expression(inputOfString("a & b & !c"))
	p, ok := res.result.(node)

	// set vars
	vars := map[string]bool{
		"a": true,
		"b": true,
		"c": false,
	}
	if p.Eval(vars) != true || !ok || res.restOfInput.getNextInput() != nil {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// set vars
	vars = map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}
	if p.Eval(vars) != false {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// set vars
	vars = map[string]bool{
		"a": false,
		"b": false,
		"c": false,
	}
	if p.Eval(vars) != false {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	res = expression(inputOfString("(a & (b | c & b)) & d"))
	p, ok = res.result.(*and)

	// set vars
	vars = map[string]bool{
		"a": true,
		"b": true,
		"c": false,
		"d": true,
	}
	if p.Eval(vars) != true || !ok || res.restOfInput.getNextInput() != nil {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// test string support
	if p.(*and).String() != "&(&('a',|('b',&('c','b'))),'d')" {
		t.Error(fmt.Sprintf("Wrong string representation: %v", p.(*and).String()))
	}
}
