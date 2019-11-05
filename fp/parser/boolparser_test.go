/**
  (C) Copyright 2018 Armin Heller

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package parser

import (
	"testing"

	. "github.com/jweigend/concepts-of-programming-languages/oop/boolparser/ast"
)

func TestMakeOr(t *testing.T) {
	var result = makeOr(Pair{Val{"a"}, Val{"b"}})
	var expected Node = Or{Val{"a"}, Val{"b"}}
	if result != expected {
		t.Errorf(
			"makeOr (Pair { Val { \"a\" }, Val { \"b\" }}) failed! Expected %v"+
				" but got wrong result %v !", expected, result)
	}
	result = makeOr(Pair{Val{"a"}, Nothing{}})
	expected = Val{"a"}
	if result != expected {
		t.Errorf(
			"makeOr (Pair { Val { \"a\" }, Nothing{} }) failed! Expected %v "+
				" but got wrong result %v !", expected, result)
	}
}

func TestMakeAnd(t *testing.T) {
	var result = makeAnd(Pair{Val{"a"}, Val{"b"}})
	var expected Node = And{Val{"a"}, Val{"b"}}
	if result != expected {
		t.Errorf(
			"makeAnd (Pair { Val { \"a\" }, Val { \"b\" }}) failed! Expected %v"+
				" but got wrong result %v !", expected, result)
	}
	result = makeAnd(Pair{Val{"a"}, Nothing{}})
	expected = Val{"a"}
	if result != expected {
		t.Errorf(
			"makeAnd (Pair { Val { \"a\" }, Nothing{} }) failed! Expected %v "+
				" but got wrong result %v !", expected, result)
	}
}

func TestMakeNot(t *testing.T) {
	var result = makeNot(0, Val{"a"})
	var expected Node = Val{"a"}
	if result != expected {
		t.Errorf("makeNot (0, Val { \"a\" }) failed! Expected %v "+
			"but got wrong result %v !", expected, result)
	}
	expected = Not{Not{Not{Val{"a"}}}}
	result = makeNot(3, Val{"a"})
	if result != expected {
		t.Errorf("makeNot (3, Val { \"a\" }) failed! Expected %v "+
			"but got wrong result %v !", expected, result)
	}
}

func TestParseVariable(t *testing.T) {
	var text = "xyz"
	var expected Node = Val{"xyz"}
	var result = parseVariable(StringToInput(text))
	if result.Result != expected {
		t.Errorf("parseVariable on input \"%v\" failed! Expected %v "+
			"but got wrong result %v !", text, expected, result.Result)
	}
	if result.RemainingInput != nil {
		var inp = result.RemainingInput.(RuneArrayInput)
		var rest = inp.Text[inp.CurrentPosition:]
		t.Errorf("parseVariable didn't eat all the input. Leftover: \"%v\"",
			string(rest))
	}
}

func TestParseExclamationMarks(t *testing.T) {
	var text = "!!!x"
	var expected int = 3
	var result = parseExclamationMarks(StringToInput(text))
	if result.Result != expected {
		t.Errorf("parseExclamationMarks on input \"%v\" failed! Expected %d "+
			"but got wrong result %d !", text, expected, result.Result)
	}
	if result.RemainingInput != nil {
		var inp = result.RemainingInput.(RuneArrayInput)
		var rest = inp.Text[inp.CurrentPosition:]
		if "x" != string(rest) {
			t.Errorf("parseExclamationMarks ate the wrong amout of input! "+
				"Leftover: \"%v\"", string(rest))
		}
	} else {
		t.Errorf("parseExclamationMarks mustn't eat all the input of \"%v\" "+
			"but it did!", text)
	}

}

func testExp(t *testing.T, text string, expected Node) {
	var result = parseExpression(StringToInput(text))
	if result.Result != expected {
		t.Errorf("parseExpression on input \"%v\" failed! Expected %v "+
			"but got wrong result %v !", text, expected, result.Result)
	}
	if result.RemainingInput != nil {
		var inp = result.RemainingInput.(RuneArrayInput)
		var rest = inp.Text[inp.CurrentPosition:]
		t.Errorf("parseExpression didn't eat all the input. "+
			"Leftover: \"%v\"", string(rest))
	}
}

func TestParseExpression(t *testing.T) {
	testExp(t, "!a", Not{Val{"a"}})
	testExp(t, "a&b", And{Val{"a"}, Val{"b"}})
	testExp(t, "a|b", Or{Val{"a"}, Val{"b"}})
	testExp(t, " a &  b", And{Val{"a"}, Val{"b"}})
	testExp(t, "a   ", Val{"a"})
	testExp(t, "a&b&c", And{Val{"a"}, And{Val{"b"}, Val{"c"}}})
	testExp(t, "a&(b&c)", And{Val{"a"}, And{Val{"b"}, Val{"c"}}})
	testExp(t, "(a&b)&c", And{And{Val{"a"}, Val{"b"}}, Val{"c"}})
	testExp(t, "a|b|c", Or{Val{"a"}, Or{Val{"b"}, Val{"c"}}})
	testExp(t, "a|(b|c)", Or{Val{"a"}, Or{Val{"b"}, Val{"c"}}})
	testExp(t, "(a|b)|c", Or{Or{Val{"a"}, Val{"b"}}, Val{"c"}})
	testExp(t, "!a & b|c&!(d|e)",
		Or{And{Not{Val{"a"}}, Val{"b"}},
			And{Val{"c"}, Not{Or{Val{"d"}, Val{"e"}}}}})

}
