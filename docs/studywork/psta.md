# Compare Object Oriented Programming in Go with Ruby
Alexander Hennecke

## Contents
1.  [Goals](#goals)
    1.  Ruby
    1.  Go
1.  [Basics](#basics)
    1.  Ruby
        1.  Interpreter
        1.  Classes
        1.  Visibility
        1.  Inheritance
        1.  Dynamic Typing
    1.  Go
        1.  Compiler
        1.  Structs
        1.  Interfaces
        1.  Embedding
        1.  Visibility
        1.  Strict Typing
1.  [Comparison OOP](#comparision-oop)
    1.  Inheritance
    1.  Polymorphism
1.  [Conclusion](#conclusion)
1.  [Sources](#sources)
        

## Goals
This paper compares object orientated programming in Go with Ruby and is part of the 
[University of Applied Sciences Rosenheim](https://th-rosenheim.de) master course [*Concepts of Programming Languages*](https://github.com/jweigend/concepts-of-programming-languages). 
Furthermore a Boolparser is being developed that is based on the Johannes Weigend 
implementation that is also available in the [GitHub repository](https://github.com/jweigend/concepts-of-programming-languages/tree/master/oop/boolparser). 


### Ruby
Ruby was one of the most popular programming languages back in the day. It was first released in 1995 under the 2-clause 
BSDL and was at its peak in 2006. The creator Yukihiro Matsumoto based Ruby on "his favorite programming languages (Perl, 
Smalltalk, Eiffel, Ada, and Lisp). Because of that Ruby uses imperative programming with functional programming 
properties.

### Go
Go is also an open source language and is being developed by Google since 2008. The language is used by the Cloud Native 
Stack and many other projects like Docker. 

## Basics
This chapter describes the basics of Ruby and Go.
 
### Ruby
Ruby is an object orientated Programming language that supports *classes*, *inheritance* and *polymorphism*. The 
language is not strict but rather uses *dynamic typing*.

#### Interpreter
The source code of a Ruby script is being interpreted while the program is running. This method is most of the time 
slower than a precompiled program, since the compiler can optimize the code before it's being executed.
Since Ruby 1.9 the *Yet Another RubyVM* (YARV) interpreter is being  used. YARV is a virtual machine that compiles a  
Ruby script to YARV instruction code sequences. It is possible to Ahead-of-Time (AOT) compile a C program that generates 
native machine code. This is more efficient than YARV instructions.

#### Classes
source 4

*Everything* in Ruby is based on a class, even the data types. The difference is that data types do not have to be 
initialized with a constructor like normal objects. For example, a string could be created by using the constructor of 
the `String` class or just with quotation marks. Both ways create the same `String` object.

```Ruby
s1 = String.new("Hello World")
s2 = "Hello World"
```


The definition of a class is shown in the example below. A class can have a constructor but is optional. The 
constructor is created by a method called `initialize` and can be given parameters.

```Ruby
class Square
  def initialize(length)

  end
end
```
An Instance of this class can be created by calling `<Class name>.new`. In this case `Square.new(4)`.

Instance or object variables are marked with an `@` symbol in front of them and are unique to every object. The variable 
has not to be declared in the constructor but could also be declared in a method.

Class variables can be used by every instance of that class. In this example the number of initiated squares is counted.
Before the class variable can be accessed, it has to be defined with a value. The method `define?` checks if the a 
value was already assigned to the variable, otherwise returns `nil`.  Since the instance variable in this example is 
always assigned with a value, it has not to be checked. If this is not the case, you should probably check the instance
 variable before accessing it.

```Ruby
class Square
  def initialize(length)
    @length = length
    if defined?(@@square_count)
      @@square_count += 1
    else
      @@square_count = 1
    end
  end

  def area
    @length * @length
  end
end
```

Methods can also be a class or instance method.
```Ruby
class Methods
  def self.important_method
    puts "Hello World from the class"
  end

  def important_method
    puts "Hello World from the object"
  end
end
```

The class method (or static method) can be called directly on the class like this:
```Ruby
Methods.important_method
```
and would return `"Hello World from the class"`. The keyword `self` points to the class itself when used inside of the 
Method definition. When used inside of a method, `self` points to the initiated object itself. (source 5)

##### Visibility
(source 4)
To control the visibility and accessibility of variables and methods of a class many programming languages use keywords 
like `public` and `private`, this is also the case for Ruby (at least for methods). The example below shows that every 
method is public if not specifically set otherwise. Every method defined under the keyword `private` is private and the 
same is also valid for the keyword `public`.

```Ruby
class Visibility
  def public_method_1
  end

private
  def private_method
  end

public
  def public_method_2
  end
end
```

A more readable way is to set the visibility after defining the method. The `private` keyword followed by the symbol of 
the method name, sets the visibility.

```Ruby
class Visibility
  def public_method_1
  end

  def private_method
  end

  def public_method_2
  end

private :private_method
end
``` 

In addition to that a method can also be `protected`. The same syntax as before is valid. If a method is protected, it 
can only be called by an object of the same class as the method itself.


Variables can either be accessed through getter and setter methods or with *accessors*. The example below shows both 
possibilities.

```Ruby
class Accessors
  attr_accessor :b

  def initialize(a, b)
    @a = a
    @b = b
  end
  
  def a
    @a
  end 

  def a=(a)
    @a = a
  end
end
```

`attr_accessor` allows the developer to directly read and write to the instance variable. `attr_reader` 
would only allow read access and `attr_writer` only write access to the variable. The getter simply returns the variable.
The setter method uses a virtual attribute by following the method name with an `=` sign (source 6). Outside of the 
class the method looks like an attribute and can also be set like one.

```Ruby
acc = Accessors.new(1,2)
acc.a = 2
acc.b = 1
```

#### Inheritance
Ruby only supports single inheritance and not multiple inheritance. A child class can only inherit from one parent class.
The methods of a parent class can be overridden.
```Ruby
class Parent
end

class Child < Parent
end
```

#### Dynamic Typing
Ruby is not a strict language and uses dynamic typing. You don't have to specify a specific type for variable. This is 
also called *duck typing* (**"If it quacks like a duck, it is a duck"**).

```Ruby
class Duck
  def quack
    'Quack'
  end
  
  def fly
    'Fly'
  end
end

class Frog
  def quack
    'Quack'
  end
end
```

In this example the duck and the frog have a method to quack. Only the duck can fly. In the next code snippet two 
methods are shown that call the `quack` or `fly` method of an object. The `duck_quack(duck)` test does not check if the 
parameter is really a duck, but just tries to call the method on the given object. In this case both calls work, since 
both animals can quack.
Since only the duck can fly, the second test `duck_fly(duck)` does not work and throws a `NoMethodError`, because the 
frog doesn't define the method `fly`.
```Ruby
def duck_quack(duck)
  duck.quack
end

def duck_fly(duck)
  duck.fly
end

duck_quack(Duck.new)
duck_quack(Frog.new)

duck_fly(Duck.new)
duck_fly(Frog.new)
```
source 3

### Go
Go provides an object oriented style of programming through interfaces. It does not support inheritance or type 
hierarchies.

#### Compiler
Go can be cross compiled for multiple architectures and operating systems. The AOT compiler is fast and generates 
performant code.

#### Structs
Since Go isn't an object oriented language, it does not have classes. Go uses structs and functions instead. The example 
below shows the definition of an struct. 
Go doesn't have constructors but the convention is to create a function `New<struc name>` that returns an initialized 
object. In this case `NewSquare` the parameter *length* sets the `length` variable of the struct.

```Go
type Square struct {
    length int
}

func NewSquare(length int) Square {
    return Square{length: length}
}
```

A function can be assigned to a struct by adding the type after the `func` keyword. In this example `s` is the instance 
of the called object.

```Go
// normal function
func area(length int) int {
    return length * length
}

// function assigned to struct
func (s *Square) area() int {
    return s.length * s.length
}
```

#### Interfaces
As described above Go doesn't support inheritance. Interfaces specify method headers that can be implemented by a struct.
In other languages is a keyword like *implements* needed to notify the compiler. In Go you must implement every 
method of an interface to use it.

In the example below a simple interface `World` with the method `hello() string` is created. The `Square` struct 
implements the method and thereby uses the interface. Since an interface is also a type, it can be used as a type for a 
variable.

```Go
type World interface {
    Hello() string
}

func (s Square) Hello() string {
    return "Hello World"
}
```

For example the `Stringer` interface describes the method `func (t T) String() string` for a struct, to convert it to a 
string.

```Go
func (s Square) String() string {
    return fmt.Sprintf("Square: length=%v", s.length)
}
```

#### Embedding
Instead of inheritance Go supports embedding. This functionality works with *structs* and *interfaces*.

The example below shows three interfaces. `Hello` and `Bye` both implement a method. The third interface `HelloBye` 
combines the two others in one interface. `HelloBye` has all method definitions of `Hello` and `Bye`. A 
*struct* that would implement the interface `HelloBye` would also be from the type `Hello` and `Bye`.

```Go
type Hello interface {
    SayHello() string
}

type Bye interface {
    SayBye() string
}

type HelloBye interface {
    Hello
    Bye
}
```

If you apply the same to *structs* the methods and variables of the embedded type are available to the "parent" struct. 
The methods can be overridden, but not overloaded.

#### Visibility
The encapsulation in Go is being restricted on package level. In most languages the visibility of methods and variables 
is achieved with keywords like `public` and `private`. In Go the visibility is described through the capitalizing of the 
first letter of an element. This can be used on methods, variables and types. Instead of calling an capitalized element 
*public*, it is called *exported*. Elements with a lowercase character are *unexported*.

```Go
package geometry

func (s *Square) ExportedMethod() {}
func (s *Square) unexportedMethod() {}
```

In this example the `ExportedMethod()` is visible outside of the package `geometry`. The method `unexportedMethod()` can 
only be used inside of the package. Since the `Square` struct from earlier is also capitalized, it is exported. The 
variable `length` cannot be accessed outside of the package.

#### Strict typing
Go is a strict language and is type safe. Every data type has a specific type and restricts the usage of a variable in 
certain situations to predefined requirements.
Since Go doesn't support inheritance, the language only supports polymorphism through interfaces.


## Comparison OOP
In this chapter both languages abilities for object orientated programming will be compared with the help of some 
examples of the boolparser implementation. The Ruby implementation can be found in this repository. The Go 
implementation can be found in the [GitHub repository](https://github.com/jweigend/concepts-of-programming-languages/tree/master/oop/boolparser) of 
Johannes Weigend.

### Inheritance
The only time that inheritance is used in the boolparser implementation is the abstract syntax tree. 

The Ruby implementation has a base class `Node` that implements a method `eval(vars)`. Ruby doesn't support abstract 
classes or interfaces, as it uses dynamic typing. Because of that the method `eval` raises an exception if called. The 
`Or` class inherits the `Node` class and must overwrite the `eval` method. The same is done for the other *node* types.

Another implementation option is shown in the polymorphism chapter below.

*boolparser/ast.rb*
```Ruby
# An abstract syntax tree for boolean expressions.
class Node
  def eval(vars)
    raise 'should be overwritten'
  end
end

# Or is the logical OR Operator in an AST
class Or < Node
  def initialize(lhs, rhs)
    @lhs = lhs
    @rhs = rhs
  end

  def eval(vars)
    @lhs.eval(vars) || @rhs.eval(vars)
  end
end
```

The Go implementation uses an *interface* for the `Node`. As every *Node* type, like `Or` and `Not`, implements this 
interface, they are also from the type `Node`. 

*oop/boolparser/ast/ast.go*
```Go
type Node interface {
    Eval(vars map[string]bool) bool
}

type Or struct {
    LHS Node
    RHS Node
}

func (o Or) Eval(vars map[string]bool) bool {
    return o.LHS.Eval(vars) || o.RHS.Eval(vars)
}
```

### Polymorphism

Ruby isn't type safe and uses dynamic typing. Because of that you have to be certain that the given object supports the 
needed functionality. In this implementation it is safe to say that the correct object is used, but if it is not certain
the input must be checked.  

*boolparser/ast.rb*
```Ruby
# Or is the logical OR Operator in an AST
class Or < Node
  def initialize(lhs, rhs)
    @lhs = lhs
    @rhs = rhs
  end

  def eval(vars)
    @lhs.eval(vars) || @rhs.eval(vars)
  end
end
```

A safer option would be to check if `lhs` *(left hand side)* and `rhs` *(right hand side)* is an instance of the 
superclass `Node`. This can be done with `is_a?`. An exception should be raised if this is not the case.

```Ruby
def initialize(lhs, rhs)
    @lhs = lhs if lhs.is_a?(Node)
    @rhs = rhs if rhs.is_a?(Node)
end
```

Another way would be to check if the object has the method implemented. This only verifies the method name not the 
parameters. So be sure that there is not another method with the same name.

```Ruby
    @lhs = lhs if lhs.respond_to?(:eval)
```


Since Go is a strictly typed language the variables `LHS` and `RHS` can only be assigned to a *struct* that implements 
the `Node` interface.

*oop/boolparser/ast/ast.go*
```Go
// Or is the logical OR Operator in an AST
type Or struct {
	LHS Node
	RHS Node
}

// Eval implements the Node interface
func (o Or) Eval(vars map[string]bool) bool {
	return o.LHS.Eval(vars) || o.RHS.Eval(vars)
}
```

## Conclusion
In the developer survey from stackoverflow in 2019  Ruby and Go had 8% as most popular programming language. As 
comparison Java and Python had 41%. Ruby is relatively old and loses popularity and Go is new and slowly gaining 
popularity. Both languages were created in different times but try to achieve some of the same things - to make life 
easier for the developers. Since the programming languages use different rudiments to get to that goal, the programming 
style is different, but both can achieve object oriented programming. 



## Sources
1.  [Ruby About](https://www.ruby-lang.org/en/about/)
1.  [Go Documentation](https://golang.org/doc/)
1.  L. Carlson, L. Richardson, Ruby Cookbook, 2006
1.  P. Cooper, Beginning Ruby: From Novice to Professional, Third Edition, 2016
1.  [Understanding `self` in Ruby](https://www.honeybadger.io/blog/ruby-self-cheat-sheet/)
1.  [Programming Ruby - The Pragmatic Programmer's Guide](http://ruby-doc.com/docs/ProgrammingRuby/html/tut_classes.html#UC)
1.  R. Olsen, Eloquent Ruby, 2011
1.  [YARV: Yet Another RubyVM](http://www.atdot.net/yarv/oopsla2005eabstract-rc1.pdf)
1.  [OOP and Go… Sorta](https://medium.com/behancetech/oop-and-go-sorta-c6682359a41b)
1.  [Concepts of Programming Languages](https://github.com/jweigend/concepts-of-programming-languages)
1.  [stackoverflow Developer Survey Results 2019](https://insights.stackoverflow.com/survey/2019#most-popular-technologies)