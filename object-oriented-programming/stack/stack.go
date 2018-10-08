// Copyright 2018 Johannes Weigend
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//Package stack contains LIFO functions.
package stack

// Stack is a generic LIFO container for untyped object.
type Stack []interface{}

// NewStack constructs an empty stack.
func NewStack() *Stack {
	return new(Stack)
}

// Push pushes a value on the stack.
func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

// Pop pops a value from the stack. It returns an error if the stack is empty.
func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		panic("can not pop: empty stack")
	}
	var result = (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return result
}

// END OMIT
