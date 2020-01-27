# Compare functional programming in Go with JavaScript

This paper compares functional programming in JavaScript with functional programming in Go.
At the beginning, it gives a short overview over JavaScript and its history.
Furthermore, an implementation of a parser for Boolean expressions is used as a practical example to compare functional programming in the two programming languages.
In addition, JavaScript and Go are compared on their support of basic and advanced functional programming concepts.
These concepts of functional programming are explained and looked at in the paper.
At the end of the paper there will be an evaluation and summary on how suitable JavaScript is for implementing a parser in a functional style.

## Table of contents

1. [JavaScript Overview](#javascript-overview)
1. [Parser for Boolean expressions](#parser-for-boolean-expressions)
1. [Functional programming concepts](#functional-programming-concepts)
    1. [Type system](#type-system)
    1. [Immutability](#immutability)
    1. [First-class functions](#first-class-functions)
    1. [Closures and lambda expressions](#closures-and-lambda-expressions)
    1. [Higher-order functions](#higher-order-functions)
    1. [Function composition](#function-composition)
    1. [Pure functions](#pure-functions)
    1. [Lazy evaluation](#lazy-evaluation)
    1. [Recursion and tail-call optimization](#recursion-and-tail-call-optimization)
    1. [Algebraic data types](#algebraic-data-types)
    1. [Pattern matching](#pattern-matching)
1. [Summary](#summary)
1. [References](#references)

## JavaScript Overview

JavaScript is a multi-paradigm programming language and a core technology of the internet.
It is a general purpose programming language and runs in the browser as well as on the server.
Despite often deceived as an object-oriented programming language, JavaScript also follows functional and imperative paradigms.
In addition, JavaScript is event-driven and has good support for asynchronous programming [moz01].

Interestingly, the original plan of Netscape was to integrate Scheme, a Lisp dialect, into their browser.
But for marketing reasons, it was decided, to create a new language with a syntax similar to Java.
Later the newly created language was called JavaScript and integrated into the Netscape browser.
Nevertheless, JavaScript has taken the functional concepts of Scheme and integrated them in to the language [ant16].

However, to stay within the scope of this paper, the focus will be on the functional aspects of JavaScript. These will be presented in the following sections.

## Parser for Boolean expressions

The parser for Boolean expressions is a practical example to compare functional programming in Go with JavaScript.
It is implemented in a functional style and will be used to provide code examples for various functional programming concepts discussed later.
Generally speaking the parser is built using parser combinators, which are well suited to be implemented with functional programming.
The resulting parser parses Boolean expressions of the following EBNF grammar.

```ebnf
<expression> ::= <term> { <or> <term> }
<term> ::= <factor> { <and> <factor> }
<factor> ::= <var> | <not> <factor> | (<expression>)
<or>  ::= '|'
<and> ::= '&'
<not> ::= '!'
<var> ::= '[a-zA-Z0-9]*'
```

`A & B | !C` is an example of an expression, that can be parsed by the parser.
Depending on the values of A, B and C, which can be `True` or `False`, the parser determines the result of the expression.
The expression is then parsed by building an abstract syntax tree (AST), consisting of `Or`, `And`, `Not` and `Value` nodes, mimicking the EBNF grammar.
The created AST allows it to determine the results of the expression, when replacing A, B and C by either `True` or `False`.

## Functional programming concepts

Functional programming and the functional programming paradigm have various unique concepts.
To see how well JavaScript and Go are suited for functional programming, we will take a look on these concepts and their support in both languages.

### Type system

JavaScript is a dynamic and weakly typed programming language, that also features duck-typing.
Weakly typed means, types are implicitly cast depending on the used operation.
Furthermore, dynamic typing allows types to change their type at runtime, when their assigned value changes [moz05].

In the context of functional programming, the dynamic and weakly typing of JavaScript simplifies writing highly reusable functions.
This is useful for higher-order functions and function composition, because there is no need to use `any` types or make frequent type casts.
The `getFirst()` function, for example, is highly flexible, because it can be used for any argument without having to specify the possible types of the `pair` argument in advance.

```javascript
export const getFirst = pair => (pair instanceof Pair ? pair.first : pair);
```

Go on the other hand is a strict and strongly typed programming language.
This means types are explicitly assigned and cannot change after assignment, except when explicitly cast by the developer.

Furthermore, there are no generic types in Go, so we have to use an empty interface to simulate an `any` type.
This makes the code more verbose and less readable than JavaScript code, while providing no benefit to the developer.
Generally speaking, the Go type system is not tailored to functional programming.

Additionally, using empty interfaces to simulate generic types has an impact on performance, especially when doing many type conversions.
This is no problem in the rather small parser example, but might become one at a larger scale and should therefore be mentioned [she17].

However, by using empty interfaces, it's possible to write flexible and reusable functions in Go, as we can see in the following example of the `GetFirst()` function.

```go
func GetFirst(argument interface{}) interface{} {
  var pair, isPair = argument.(Pair)
  if isPair {
    return pair.First
  }
  return argument
}
```

### Immutability

Immutability is a desired property, especially for functional programming, because it reduces unintended side effects.

Unfortunately, true immutability can't be achieved in JavaScript.
Although it's possible to create constructs that are sort of immutable, there is no built-in or default immutability as in pure functional programming languages like Haskell.

One way to achieve immutability in JavaScript is to use the `const` keyword, that allows to define constant variables, functions or objects.
But while primitive types as strings are truly constant, when declared with the `const` keyword, objects declared with the `const` keyword are still mutable [moz07].
As we can see in the example below, the properties of the `result` object can still be reassigned, despite the `result` object being declared as `const`.
So, used on an object, the `const` keyword only prevents to assign a new value to the object.

```javascript
export const optional = parser => input => {
  const result = parser(input);

  if (result.result === null) {
    result.result = new Nothing();
    result.remainingInput = input;
  }

  return result;
};
```

Another way to achieve immutability is to _freeze_ an object after creation.
This makes the object truly immutable, but still has the caveat, that it doesn't effect nested objects.
Therefore, it's necessary to call _freeze_ recursively on an object that should be truly immutable [moz08].

However, this is prone to developer mistakes and may not play nicely with libraries expecting mutable objects.

In Go immutability is quite similar to JavaScript and can't be easily achieved.
Except strings, data types in Go are mutable by default and only primitive data types like `bool` and `int` can declared to be constant.
The immutability of composite data types like Go's `structs` on the other hand, is in the responsibility of the developer.

For Go 2.0 however, there is a proposal to introduce new immutable data types, so this might change in the future [git01].

To sum it up, as of today there is some support for immutability in JavaScript and in Go, but not by default and not easily usable.

### First-class functions

First-class functions are the foundation of supporting functional programming in a programming language.
A language with first-class functions has to meet the following criteria:

- Allow passing functions as parameters to other functions.
- Allow functions to be return values of other functions, so that functions can return functions.
- Allow functions to be assigned to variables.
- Allow functions to be stored in data structures like arrays.

The listed properties also allow for concepts such as higher-order functions and function composition, both of which are described later.

In JavaScript, all the criteria is met.
Therefore, functions in JavaScript are first-class functions and are treated like first-class citizens [moz06][fog13].
This allows us to assign the `optional()` and the `expectSeveral()` function to the `expectSpaces` variable in the example below.

```javascript
const expectSpaces = optional(expectSeveral(isSpaceChar, isSpaceChar));
```

The same applies to Go, which has the same support for first-class functions as JavaScript [she17][gol01].
Like in JavaScript it's possible to assign the `ExpectSeveral()` and the `Optional()` function to the `ExpectSpaces` variable.
The only difference is the chaining of function calls instead of nesting and the explicit `Parser` type of the variable.

```go
var ExpectSpaces Parser = ExpectSeveral(isSpaceChar, isSpaceChar).Optional()
```

### Closures and lambda expressions

Closures or lambda expressions, also called anonymous functions, are unnamed functions, often returned from another function.
To be precise, a closure is the reference to the local state of the function, that returns an anonymous function.
However, both terms are often used interchangeably, since both concepts belong to the concept of anonymous functions, returned by an outer function [fog13].

Support for closures is found in all programming languages with first-class functions, because closures are needed for anonymous functions to work.
Without closures and thus without references to the _outer_ function, the _inner_ function would stop working when called directly [moz02][fog13].

```javascript
const expectString = expectedString => input => expectCodePoints(expectedString)(input);
```

In JavaScript, lambda expressions can be written very concisely using the arrow function syntax introduced with ECMAScript 6 [moz04].
This can be seen in the example above where the `expectString()` function takes `expectedString` as its argument and returns an anonymous function with `input` as its argument.

```go
func ExpectString(expectedString string) Parser {
  return func(Input Input) Result {
    var result = ExpectCodePoints([]rune(expectedString))(Input)
    var runes, isRuneArray = result.Result.([]rune)

    if isRuneArray {
      result.Result = string(runes)
    }

    return result
  }
}
```

In Go lambda expressions are more verbose, mainly because the type system requires explicit types.
This can be seen in the example of the `expectString()` function above, where we have explicit types on the arguments and the return values, on both, the inner and the outer function.
Apart from that, lambda expressions in Go are equal to lambda expressions in JavaScript [she17].

### Higher-order functions

Higher-order functions are functions that accept other functions as arguments or return a function as a result.
As discussed in the section on first-class functions, JavaScript has first-class functions and thus allows writing and using higher-order functions.

The convert function of the Boolean parser for example takes two functions as its arguments.
A parser function to be executed and a converter function to convert the result of the parser function.
This flexible and easily reusable higher order function can be used for any parser function and with any converter function.

```javascript
export const convert = (parser, converter) => input => {
  const result = parser(input);

  if (result.result !== null) {
    result.result = converter(result.result);
  }

  return result;
};
```

The converter function written in Go looks quite similar to the JavaScript implementation.
This is the case, because Go has similar support for higher-order functions as JavaScript.
The main differences between the Go and the JavaScript implementation are the explicit empty interfaces to satisfy the Go type system and the higher verbosity of the Go code [med02][she17].

```go
func (parser Parser) Convert(converter func(interface{}) interface{}) Parser {
  return func(Input Input) Result {
    var result = parser(Input)

    if result.Result != nil {
      result.Result = converter(result.Result)
    }

    return result
  }
}
```

### Function composition

Function composition is quite similar to higher-order functions and can be seen as an application of higher-order functions.
Function composition generally describes the act of combining multiple functions together to create a more complex function [hac01].

The Boolean parser utilizes function composition to compose the Boolean parser out of simpler parsers, that only parse parts of a Boolean expression.
The following example of the `parseOr()` function shows how the rather complex function is composed out of many simpler functions.

```javascript
const parseOr = input =>
  convert(
    andThen(parseAnd, optional(second(andThen(expect("|"), parseOr)))),
    makeOr,
  )(input);
```

Therefore, function composition is an important part of functional programming, because it allows us to compose complex software out of simple functions.
Because both, Go and JavaScript, have support for higher-order functions, the same differences, as mentioned in the higher-order functions section, are applicable to function composition.

### Pure functions

Pure functions are functions that have no side effects and no hidden inner state.
This means, a function, given the same input, always produces the same output.
To achieve this, a pure function uses only its input and does not use or mutate the internal state.
This property of pure functions is called referential transparency and allows to replace a function with its result without changing the behaviour of a program.

```javascript
const isDigit = codePoint => "0" <= codePoint && codePoint <= "9";
```

The `isDigit()` function in the above example is pure, because it doesn't mutate the given `codepoint` and always returns true if a digit between `0` and `9` is given.
JavaScript thus allows the writing of pure functions, but has no special constructs to enforce side effect free and therefore pure functions [fog13].

The same applies to Go.
As in JavaScript, it's possible to write pure and side effect free functions in Go, but there are no special constructs to enforce these concepts.
Since Go also doesn't support tail-call optimization, which will be discussed later, there is a performance impact on the heavy use of pure functions and recursion.
So as long as there is no tail-call optimization in Go, pure functions should be used with precaution [she17].

Therefore, pure functions are possible in both languages, but it's the responsibility of the developer to keep them pure.

### Lazy evaluation

There are two ways to evaluate functions, eager and lazy evaluation.
Programming languages that use eager evaluation, evaluate a function as soon as it's assigned or defined.
Lazy evaluation, on the other hand, means that functions are evaluated when they are executed, which may happen much later than the assignment.

In functional programming with heavy use of functions, lazy evaluation is useful for performance optimization.
This is possible, because functions are only evaluated, when they are actually used and therefore no unnecessary calculations are done.

Unfortunately both, Go and JavaScript, use eager evaluation for functions with no built-in support for lazy evaluation.
However, it's possible to simulate lazy evaluation in both languages, but it's no core part of the two programming languages [med02][med04].

### Recursion and tail-call optimization

To avoid mutating state functional programming makes heavy use of recursion.
Recursion happens when a function calls itself with new parameters to compute something instead of mutating state inside the function.
Unfortunately recursion is less efficient than iteration.

A technique for remedying the performance issues of recursion is called tail-call optimization (TCO).
Without TCO, a new stack frame is added to the call stack each time a function is called recursively.
Therefore, the call stack grows with every function call and results in high memory consumption for deeply nested recursion.
TCO prevents this by overwriting the unneeded stack frames of previous function calls.
TCO is thus required for efficient functional programming.

JavaScript supports TCO since the introduction of ECMAScript 6 in 2015.

Go on the other side has no support for TCO and according to the Go developers, they don't see this as a problem that affects many people, so they won't add support for TCO [git04].
This means that the heavy use of recursion and functional programming in Go will have an impact on performance.
In fact, there are some workarounds for this issue, but they are out of the scope of this paper [med02][she17].

Summarized, JavaScript has more advanced support for efficient recursive programming than Go.

### Algebraic data types

Algebraic data types, also named sum/product types or discriminate unions, are a concept in functional programming languages for representing data structures by composing them with other types.

An example for this is the AST of our parser example.
The AST is a data structure consisting of nodes with the `OR`, `AND`, `NOT` or `VALUE` type.

With support of algebraic data types we could write the AST much more concisely, as we can see in following example.

```typescript
type Or = { lhs: Node, rhs: Node };
type And = { lhs: Node, rhs: Node };
type Not = { ex: Node };
type Value = { name: string };

type Node = Or | And | Not | Value;
```

Unfortunately this is not possible in JavaScript because it's missing explicit types.
However, by using TypeScript, a JavaScript superset, it would be possible to use this syntax today [typ01].

This feature is also missing in Go, but there is an ongoing discussion on the introduction of sum types along with generic types in Go 2.0 [gol02][git03].
This means that Go could receive support for sum types in the future, thus allowing easy representation of AST nodes.

### Pattern matching

Pattern matching is a concept to work with data structures from primary functional programming languages like Haskell [has01].
It's often used in conjunction with algebraic data types to select different behaviour depending on the data type.

In our parser example this would be useful for the `evaluate()` function, which could be written in a more functional style instead of using JavaScript classes.
There is a stage 1 proposal to introduce pattern matching to ECMAScript in the future [git02].
This means, that in the future, the `evaluate()` function could be written as concisely as in the following example.

```javascript
const evaluate = (vars, node) => case (node) {
  when (node instanceof Or) -> evaluate(vars, node.lhs) || evaluate(vars, node.rhs)
  when (node instanceof And) -> evaluate(vars, node.lhs) && evaluate(vars, node.rhs)
  when (node instanceof Not) -> !evaluate(vars, node.ex)
  when (node instanceof Value) -> vars.get(node.name)
}
```

In Go, by contrast, there is no support for pattern matching and there are no plans to introduce it to the language.
But it even in absent of pattern matching in Go, something similar can be achieved by using interfaces and switch statements [eli01].

## Summary

After implementing the Boolean parser, the author is convinced that JavaScript is well suited to implement a parser.
As we have seen, JavaScript has good support for all basic concepts of functional programming.
Only some advanced concepts like pattern matching or algebraic data types are missing, which would have simplified the implementation of the parser.

The only real downside is the dynamic and weakly typed type system of JavaScript.
While this makes it easy to write reusable functions, it also makes it difficult to detect errors.
This is especially the case with some parser functions that take many inputs, such as the `convert()` function.
However, this issue is a more general problem in JavaScript than a specific problem in functional programming.

When it comes to implementing the parser, the differences between Go and JavaScript are subtle.
The only outstanding difference is the higher verbosity of the Go code, due to type annotations and type casts.

On the whole, JavaScript shows more focus on functional programming compared to Go, especially on topics like tail-call optimization or short lambda expressions.
Therefore, support for functional programming in JavaScript is more advanced than in Go.

## References

- [ant16] JavaScript: Functional Programming for JavaScript Developers, Ved Antani; Simon Timms; Dan Mantyla, Packt Publishing, 2016-08-31
- [eli01] [Go and Algebraic Data Types](https://eli.thegreenplace.net/2018/go-and-algebraic-data-types/) (viewed 2020-01-05)
- [fog13] Functional JavaScript, Michael Fogus, O'Reilly Media, Inc., 2013-06-10
- [git01] [proposal: Go 2: immutable type qualifier](https://github.com/golang/go/issues/27975) (viewed 2019-12-26)
- [git02] [ECMAScript Pattern Matching](https://github.com/tc39/proposal-pattern-matching) (viewed 2019-12-27)
- [git03] [proposal: spec: add sum types / discriminated unions #19412](https://github.com/golang/go/issues/19412) (viewed 2020-01-03)
- [git04] [proposal: Go 2: add become statement to support tail calls](https://github.com/golang/go/issues/22624) (viewed 2020-01-05)
- [gol01] [Codewalk: First-Class Functions in Go](https://golang.org/doc/codewalk/functions/) (viewed 2019-12-26)
- [gol02] [Go FAQ: Why does Go not have variant types?](https://golang.org/doc/faq#variant_types) (viewed 2020-01-05)
- [hac01] [Functional programming paradigms in modern JavaScript: Function Composition](https://hackernoon.com/functional-programming-paradigms-in-modern-javascript-function-composition-109670038859) (viewed 2020-01-08)
- [has01] [Case Expressions and Pattern Matching](https://www.haskell.org/tutorial/patterns.html) (viewed 2020-01-03)
- [ker17] Mastering Javascript Functional Programming, Federico Kereki, Packt Publishing, 2017-12-29
- [med01] [Introduction to Functional JavaScript](https://medium.com/functional-javascript/introduction-to-functional-javascript-45a9dca6c64a) (viewed 2019-12-21)
- [med02] [Functional Go](https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4) (viewed 2019-12-26)
- [med03] [JS ES6 Recursive Tail Call Optimization](https://medium.com/@mlaythe/js-es6-recursive-tail-call-optimization-feaf2dada3f6) (viewed 2019-12-31)
- [med04] [Lazy Evaluation in Javascript](https://medium.com/hackernoon/lazy-evaluation-in-javascript-84f7072631b7) (viewed 2020-01-05)
- [moz01] [MDN JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript) (viewed 2019-12-25)
- [moz02] [MDN Closures](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Closures) (viewed 2019-12-23)
- [moz03] [MDN Functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Functions) (viewed 2019-12-25)
- [moz04] [MDN Arrow function expressions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Arrow_functions) (viewed 2019-12-25)
- [moz05] [MDN JavaScript data types and data structures](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Data_structures) (viewed 2019-12-25)
- [moz06] [MDN First-class Function](https://developer.mozilla.org/en-US/docs/Glossary/First-class_Function) (viewed 2019-12-25)
- [moz07] [MDN const](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/const) (viewed 2019-12-26)
- [moz08] [MDN Object.freeze()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/freeze) (viewed 2019-12-26)
- [she17] Learning Functional Programming in Go, Lex Sheehan, Packt Publishing, 2017-11-24
- [typ01] [Discriminated Unions](https://www.typescriptlang.org/docs/handbook/advanced-types.html#discriminated-unions) (2020-01-05)
