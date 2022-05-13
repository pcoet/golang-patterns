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

func TestIntBag(t *testing.T) {
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

func TestStringBag(t *testing.T) {
	b := Bag{}

	if b.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, b.IsEmpty())
	}

	strs := []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, s := range strs {
		b.Add(s)
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
		s    string
		ok   bool
	}{
		{"t1", "a", true},
		{"t2", "b", true},
		{"t3", "c", true},
		{"t4", "d", true},
		{"t5", "e", true},
		{"t6", "f", true},
		{"t7", "g", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, ok := next()
			if !((s == tc.s) && (ok == tc.ok)) {
				t.Errorf("expected %v and %v; got %v and %v", tc.s, tc.ok, s, ok)
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
