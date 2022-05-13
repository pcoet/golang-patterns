// Copyright 2022 Google LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

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
