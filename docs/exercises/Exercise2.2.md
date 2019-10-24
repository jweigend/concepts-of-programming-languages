# Exercise 2.2 - Basic OOP

If you do not finish during the lecture period, please finish it as homework.

## Rational Numbers (Datatypes)
Write a type for rational numbers (a/b). Implement a constructor, methods for adding and multiplying 
rational numbers. 
- Implement the euclidean algorithm for reducing rational numbers. Make sure the == and != operators work correct for 1/2 and 2/4  
- Write a constructor function to create rational numbers
- The type should be immutable
- Make sure to implement the Stringer() interface
- Write an unit test to test the rational numbers
- Check the test coverage. It should be greater than 80%

### Questions
- What type has the receiver (pointer or value). Why?
- What does immutable mean? Why is it good design to make Rational immutable?

## Stack (Containers)
Write a generic LIFO container (stack) for all types. The stack should have at least four methods:
- Push(object)
- Pop()
- Size()
- GetAt(index) 
- Write a unit test which uses build in types and object types (Rational) inside the stack. Calculate the sum of all elements in the stack.

## After this Exercise
- You should know how to write OO code in Go
- You should how to reference other packages in your codebase
- You should know the difference between object and data types
- You should know how generic containers are built with Go

