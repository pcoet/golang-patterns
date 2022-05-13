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
	"testing"
)

func TestMaxPQ(t *testing.T) {
	q := MaxPQ{}

	if q.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, q.IsEmpty())
	}

	if q.Size() != 0 {
		t.Errorf("expected %v; got %v", 0, q.Size())
	}

	a := []int{8, 1, 89, 2, 21, 3, 0, 1, 55, 5, 13, 34}
	for _, n := range a {
		q.Insert(n)
	}

	if q.IsEmpty() != false {
		t.Errorf("expected %v; got %v", false, q.IsEmpty())
	}

	if q.Size() != 12 {
		t.Errorf("expected %v; got %v", 12, q.Size())
	}

	if q.Max() != 89 {
		t.Errorf("expected %v; got %v", 89, q.Max())
	}

	testCases := []struct {
		name string
		max  int
		size int
	}{
		{"t1", 89, 11},
		{"t2", 55, 10},
		{"t3", 34, 9},
		{"t4", 21, 8},
		{"t5", 13, 7},
		{"t6", 8, 6},
		{"t7", 5, 5},
		{"t8", 3, 4},
		{"t9", 2, 3},
		{"t10", 1, 2},
		{"t11", 1, 1},
		{"t12", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := q.DelMax()
			s := q.Size()
			if !((m == tc.max) && (s == tc.size)) {
				t.Errorf("expected %v and %v; got %v and %v", tc.max, tc.size, m, s)
			}
		})
	}

	if q.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, q.IsEmpty())
	}

	a = []int{21, 34, 55, 89}
	for _, n := range a {
		q.Insert(n)
	}

	moreCases := []struct {
		name  string
		del   int
		max   int
		size  int
		empty bool
	}{
		{"t1", 89, 55, 3, false},
		{"t2", 55, 34, 2, false},
		{"t3", 34, 21, 1, false},
	}

	for _, tc := range moreCases {
		t.Run(tc.name, func(t *testing.T) {
			d := q.DelMax()
			m := q.Max()
			s := q.Size()
			e := q.IsEmpty()
			if !((d == tc.del) && (m == tc.max) && (s == tc.size) && (e == tc.empty)) {
				t.Errorf("expected %v, %v, %v, and %v; got %v, %v, %v, and %v", tc.del, tc.max, tc.size, tc.empty, d, m, s, e)
			}
		})
	}

	q.DelMax()
	if q.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, q.IsEmpty())
	}
}
