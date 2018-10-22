# Exercise 4 - Functional Programming in Go

If you are not get finished during the lecture hours, please finish it as homework.

## Exercise 4.1 - Warm Up

Write a Go Programm which shows the following concepts:

- Functions as Variables
- Anonymous Lambda Functions
- High Order Functions (functions as parameters or return values)

Compare the Syntax with Java and discuss this in a group of students.

## Exercise 4.2 - Functional Composition

Write a Go function to compose two unknown unary functions with one parameter and one return value. The functions to compose should be parameters.
Write a Unit Test for that function.

## Exercise 4.2 - Map / Filter / Reduce
Map/Reduce is a famous functional construct implemented in many parallel and distributed collection frameworks like
HADOOP, SPARK, Java Streams (not distributed but parallel), C# Link

- Implement a custom Java like Stream API with the following interface:
 ```go
    type Stream interface {
    	Map(m Mapper) Stream
    	Filter(p Predicate) Stream
    	Reduce(a Accumulator) Any
    }
```
The usage of the interface should be like this:
```go
    stringSlice := []Any{"a", "b", "c", "1", "D"}

	// Map/Reduce
	result := ToStream(stringSlice).
		Map(toUpperCase).
		Filter(notDigit).
		Reduce(concat).(string)

	if result != "A,B,C,D" {
		t.Error(fmt.Sprintf("Result should be 'A,B,C,D' but is: %v", result))
    }
```

 Questions:
 - Describe the type of Mapper, Predicate and Accumulator
 - How can you make the types generic, so they work not only for strings?

## Exercise 4.3 - Wordcount
Wordcount is a famous Hello World Algorithm for demonstrating the power of distributed functional collection frameworks. 
Wordcount counts all words in a collection. After running this in parallel or even distributed, you have the following result:

INPUT:  "A" "B" "C" "A" "B" "A"
OUTPUT: "A:3" "B:2" "C:1"

- How can you implement the problem with the already built Map/Filter/Reduce functions?
- Write an Unit Test to prove that your solution works correct


