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
