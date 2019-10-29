// (c) Armin Heller 2018

package functionalparsers

import "container/list"

//  ---------------------------------------------------------
// The expression should have the following EBNF form:
//  EBNF
//  ---------------------------------------------------------
// 	<expression> ::= <term> { <or> <term> }
// 	<term> ::= <factor> { <and> <factor> }
// 	<factor> ::= <constant> | <not> <factor> | (<expression>)
// 	<value> ::= var
// 	<or>  ::= '|'
// 	<and> ::= '&'
// 	<not> ::= '!'
//  <var> ::= '[a-zA-Z0-9]*'
//  ---------------------------------------------------------

// Meine Grammatik ist etwas anders aufgebaut, die Ausdrücke sind alle rechts-geklammert.
// Sowohl & als auch | sind assoziativ, deshalb macht das für die Auswertung keinen Unterschied.
// Die Strings der nodes sehen jedoch anders aus.
// <expression> ::= <term> ( <or> <expression> )?
// <term>       ::= <factor> ( <and> <term> )?
// <var>        ::= [a-zA-Z0-9]+ // Auch Ziffern am Anfang des Namens erlaubt. Auch Namen erlaubt, die nur Ziffern enthalten.
// Der Rest der Grammatik ist gleich.

func expression(i ParserInput) ParserResult {
	return MyParser(term).andThen(
		expect("|").andThen(expression).optional().second()).apply(convertPairToOrNode)(i)
}

func term(i ParserInput) ParserResult {
	return MyParser(factor).andThen(expect("&").andThen(term).optional().second()).apply(convertPairToAndNode)(i)
}

func factor(i ParserInput) ParserResult {
	return expect("!").andThen(factor).second().apply(convertToNotNode).orElse(
		ident.apply(newValue).orElse(
			expect("(").andThen(MyParser(expression).andThen(expect(")"))).apply(
				dropParenthesis)))(i)
}

// Parses identifiers
var ident = maybeSpacesInFront(expectIdentifier)

// expectIdentifier parses exactly one identifier from the input
func expectIdentifier(i ParserInput) ParserResult {
	result := MyParser(expectIdentifierCharacter).onceOrMore()(i)
	if result.result != nil {
		resultString := ""
		for e := result.result.(*list.List).Front(); e != nil; e = e.Next() {
			resultString = resultString + string(e.Value.(rune))
		}
		result.result = resultString
	}
	return result
}

// MyParser parses some input and returns a result
type MyParser func(ParserInput) ParserResult

// ParserInput is the input of a Parser
type ParserInput interface {
	// Current code point or '\0' if we're at the end of the input.
	getCurrentCodePoint() rune
	// Rest of the input or nil if we're at the end.
	getNextInput() ParserInput
}

// ParserResult is the result of a MyParser
type ParserResult struct {
	// Result of the parse or nil if parsing failed
	result interface{}
	// The rest of the input after producing the result, in case there is any
	restOfInput ParserInput
}

// RuneArrayInput is an implementation of ParserInput
type RuneArrayInput struct {
	text  []rune
	index int
}

// Converts a string into a ParserInput
func inputOfString(s string) RuneArrayInput {
	return RuneArrayInput{[]rune(s), 0}
}

func (i RuneArrayInput) getCurrentCodePoint() rune {
	if i.index >= len(i.text) {
		return '\x00'
	}
	return i.text[i.index]
}

func (i RuneArrayInput) getNextInput() ParserInput {
	if i.index >= len(i.text) {
		return nil
	}
	return RuneArrayInput{i.text, i.index + 1}
}

// expectCodePoint succeeds if the next code point is exactly the expected one
func expectCodePoint(r rune) MyParser {
	return func(i ParserInput) ParserResult {
		result := ParserResult{}
		if r == i.getCurrentCodePoint() {
			result.result = r
			result.restOfInput = i.getNextInput()
		}
		return result
	}
}

// expectCodePoints succeeds if the next several code points are exactly what's specified by rs.
func expectCodePoints(rs []rune) MyParser {
	return func(i ParserInput) ParserResult {
		input := i
		result := ParserResult{}
		for _, r := range rs {
			result = expectCodePoint(r)(input)
			if nil == result.restOfInput {
				return result
			}
			input = result.restOfInput
		}
		result.result = rs
		result.restOfInput = input
		return result
	}
}

// expectString Expects the code points of the string argument in the parser input
func expectString(s string) MyParser {
	return func(i ParserInput) ParserResult {
		result := expectCodePoints([]rune(s))(i)
		res, ok := result.result.([]rune)
		if ok {
			result.result = string(res)
		}
		return result
	}
}

// anyTimes uses zero or more times the parser p to parse
func (p MyParser) anyTimes() MyParser {
	return func(i ParserInput) ParserResult {
		result := ParserResult{}
		result.result = list.New()
		result.restOfInput = i
		for result.restOfInput != nil {
			oneResult := p(result.restOfInput)
			if oneResult.result == nil {
				return result
			}
			result.result.(*list.List).PushBack(oneResult.result)
			result.restOfInput = oneResult.restOfInput
		}
		return result
	}
}

// onceOrMore is like anyTimes except that zero times is not allowed
func (p MyParser) onceOrMore() MyParser {
	return func(i ParserInput) ParserResult {
		resultList := list.New()
		resultRestOfInput := i
		for resultRestOfInput != nil {
			oneResult := p(resultRestOfInput)
			if oneResult.result == nil {
				break
			}
			resultList.PushBack(oneResult.result)
			resultRestOfInput = oneResult.restOfInput
		}
		if resultList.Len() > 0 {
			return ParserResult{resultList, resultRestOfInput}
		}
		return ParserResult{}
	}
}

// ResultPair is just a pair of results
type ResultPair struct {
	first  interface{}
	second interface{}
}

// andThen parses p first and then q
func (p MyParser) andThen(q MyParser) MyParser {
	return func(i ParserInput) ParserResult {
		result := ParserResult{}
		firstResult := p(i)
		if firstResult.result != nil {
			secondResult := q(firstResult.restOfInput)
			if secondResult.result != nil {
				result.result = ResultPair{firstResult.result, secondResult.result}
				result.restOfInput = secondResult.restOfInput
				return result
			}
		}
		return result
	}
}

// orElse parses p if possible and if p doesn't work it parses q
func (p MyParser) orElse(q MyParser) MyParser {
	return func(i ParserInput) ParserResult {
		firstResult := p(i)
		if firstResult.result != nil {
			return firstResult
		}
		return q(i)
	}
}

// apply applies a function to the result of a parser if the parsing was successful
func (p MyParser) apply(f func(interface{}) interface{}) MyParser {
	return func(i ParserInput) ParserResult {
		result := p(i)
		if result.result != nil {
			result.result = f(result.result)
		}
		return result
	}
}

// ExpectIdentifierCharacter expects one character that an be part of an identifier
func expectIdentifierCharacter(i ParserInput) ParserResult {
	result := ParserResult{}
	c := i.getCurrentCodePoint()
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		result.result = i.getCurrentCodePoint()
		result.restOfInput = i.getNextInput()
	}
	return result
}

var expectSpaces = (expectString(" ").orElse(expectString("\n")).orElse(expectString("\t")).orElse(expectString("\r"))).anyTimes()

func maybeSpacesInFront(p MyParser) MyParser {
	return expectSpaces.andThen(p).second()
}

// Nothing is the result of an optional value that is not present
type Nothing struct {
}

// optional parses zero or one times, if zero the result is Nothing{}
func (p MyParser) optional() MyParser {
	return func(i ParserInput) ParserResult {
		result := p(i)
		if nil == result.result {
			result.result = Nothing{}
			result.restOfInput = i
		}
		return result
	}
}

// Expect some string maybe with some spaces in front
func expect(s string) MyParser {
	return maybeSpacesInFront(expectString(s))
}

// getSecond extracts the second component of a pair if the argument is a pair. Otherwise it returns the argument.
func getSecond(x interface{}) interface{} {
	pair, isPair := x.(ResultPair)
	if isPair {
		return pair.second
	}
	return x
}

func newValue(s interface{}) interface{} {
	return newVal(s.(string))
}

func convertPairToNode(isOr bool) func(x interface{}) interface{} {
	return func(x interface{}) interface{} {
		pair, ok := x.(ResultPair)
		if ok {
			_, isEmpty := pair.second.(Nothing)
			if isEmpty {
				return pair.first.(node)
			}
			if isOr {
				return &or{pair.first.(node), pair.second.(node)}
			}
			return &and{pair.first.(node), pair.second.(node)}
		}
		return nil
	}
}

func convertPairToAndNode(x interface{}) interface{} {
	return convertPairToNode(false)(x)
}

func convertPairToOrNode(x interface{}) interface{} {
	return convertPairToNode(true)(x)
}

func dropParenthesis(x interface{}) interface{} {
	return x.(ResultPair).second.(ResultPair).first
}

// second extracts the second componend of a pair parsed with p.andThen and drops the first component
func (p MyParser) second() MyParser {
	return p.apply(getSecond)
}

func convertToNotNode(x interface{}) interface{} {
	return &not{x.(node)}
}
