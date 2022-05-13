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

func TestQueue(t *testing.T) {
	q := Queue{}

	if q.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, q.IsEmpty())
	}

	if q.Size() != 0 {
		t.Errorf("expected %v; got %v", 0, q.Size())
	}

	// add 7 ints to the queue
	nums := []int{1, 1, 2, 3, 5, 8, 13}
	for _, n := range nums {
		q.Enqueue(n)
	}

	if q.IsEmpty() != false {
		t.Errorf("expected %v; got %v", false, q.IsEmpty())
	}

	if q.Size() != 7 {
		t.Errorf("expected %v; got %v", 7, q.Size())
	}

	n := q.Dequeue()

	if n != 1 {
		t.Errorf("expected %v; got %v", 1, n)
	}

	if q.Size() != 6 {
		t.Errorf("expected %v; got %v", 6, q.Size())
	}

	if q.Peek() != 1 {
		t.Errorf("expected %v; got %v", 1, q.Peek())
	}

	// create an iterator
	next := q.Iterator()

	testCases := []struct {
		name string
		n    int
		ok   bool
	}{
		{"t1", 1, true},
		{"t2", 2, true},
		{"t3", 3, true},
		{"t4", 5, true},
		{"t5", 8, true},
		{"t6", 13, true},
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

func ExampleQueue() {
	q := Queue{}
	nums := []int{1, 1, 2, 3, 5, 8, 13}
	for _, n := range nums {
		q.Enqueue(n)
	}
	next := q.Iterator()
	// iterate over the items in the queue
	for {
		n, ok := next()
		if !ok {
			break
		}
		fmt.Println(n)
	}
	// Output:
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
}
