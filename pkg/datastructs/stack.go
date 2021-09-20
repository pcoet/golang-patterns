// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datastructs

// Stack represents a Stack of integers. It supports pushing, popping, and peeking.
type Stack struct {
	a []int // array of items
}

// IsEmpty returns true if the Stack is empty; false otherwise.
func (s *Stack) IsEmpty() bool {
	return len(s.a) == 0
}

// Size returns the number of items in the Stack.
func (s *Stack) Size() int {
	return len(s.a)
}

// Push pushes an item onto the Stack.
func (s *Stack) Push(item int) {
	// Note that the built-in append function runs in amortized O(1) time, i.e. it's performant.
	s.a = append(s.a, item)
}

// Pop removes and returns the item most recently added to the stack.
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("cannot pop from an empty stack")
	}
	n := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return n
}

// Peek returns, but does not remove, the most recently added item.
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("cannot peek into an empty stack")
	}
	return s.a[len(s.a)-1]
}

// Iterator returns a function for iterating over the items in the stack.
func (s *Stack) Iterator() func() (int, bool) {
	i := s.Size() - 1
	return func() (int, bool) {
		if i < 0 {
			return 0, false
		}
		n := s.a[i]
		i = i - 1
		return n, true
	}
}
