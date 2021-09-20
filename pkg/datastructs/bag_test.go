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

func TestBag(t *testing.T) {
	b := Bag{}

	if b.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, b.IsEmpty())
	}

	nums := []int{1, 1, 2, 3, 5, 8, 13}
	for _, n := range nums {
		b.Add(n)
	}

	if b.IsEmpty() != false {
		t.Errorf("expected %v; got %v", false, b.IsEmpty())
	}

	if b.Size() != 7 {
		t.Errorf("expected %v; got %v", 7, b.Size())
	}

	// create an iterator
	next := b.Iterator()

	testCases := []struct {
		name string
		n    int
		ok   bool
	}{
		{"t1", 1, true},
		{"t2", 1, true},
		{"t3", 2, true},
		{"t4", 3, true},
		{"t5", 5, true},
		{"t6", 8, true},
		{"t7", 13, true},
		{"t8", 0, false},
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

func ExampleBag() {
	b := Bag{}
	nums := []int{1, 1, 2, 3, 5, 8, 13}
	for _, n := range nums {
		b.Add(n)
	}
	next := b.Iterator()
	// iterate over the items in the bag
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
