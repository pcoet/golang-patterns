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

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}

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
	s := Stack{}
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
