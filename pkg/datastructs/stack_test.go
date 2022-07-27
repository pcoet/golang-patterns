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

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack[int]{}

	// new stack should be empty
	if s.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, s.IsEmpty())
	}

	// new stack should have size 0
	if s.Size() != 0 {
		t.Errorf("expected %v; got %v", 0, s.Size())
	}

	// add 7 ints to the stack
	nums := []int{1, 1, 2, 3, 5, 8, 13}
	for _, n := range nums {
		s.Push(n)
	}

	if s.IsEmpty() != false {
		t.Errorf("expected %v; got %v", false, s.IsEmpty())
	}

	if s.Size() != 7 {
		t.Errorf("expected %v; got %v", 7, s.Size())
	}

	n := s.Pop()

	if n != 13 {
		t.Errorf("expected %v; got %v", 13, n)
	}

	if s.Size() != 6 {
		t.Errorf("expected %v; got %v", 6, s.Size())
	}

	n = s.Peek()

	if n != 8 {
		t.Errorf("expected %v; got %v", 8, n)
	}

	// size of stack should not change after a peek
	if s.Size() != 6 {
		t.Errorf("expected %v; got %v", 6, s.Size())
	}

	// create an iterator
	next := s.Iterator()

	testCases := []struct {
		name string
		n    int
		ok   bool
	}{
		{"t1", 8, true},
		{"t2", 5, true},
		{"t3", 3, true},
		{"t4", 2, true},
		{"t5", 1, true},
		{"t6", 1, true},
		{"t7", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n, ok := next()
			if !((n == tc.n) && (ok == tc.ok)) {
				t.Errorf("expected %v and %v; got %v and %v", tc.n, tc.ok, n, ok)
			}
		})
	}
}

func ExampleStack() {
	s := Stack[int]{}
	nums := []int{1, 1, 2, 3, 5, 8, 13}
	for _, n := range nums {
		s.Push(n)
	}
	next := s.Iterator()
	// iterate over the items in the stack
	for {
		n, ok := next()
		if !ok {
			break
		}
		fmt.Println(n)
	}
	// Output:
	// 13
	// 8
	// 5
	// 3
	// 2
	// 1
	// 1
}
