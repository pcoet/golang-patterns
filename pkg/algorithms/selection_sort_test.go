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

package algorithms

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		ex   []int
	}{
		{"t1", []int{2, 3, 1}, []int{1, 2, 3}},
		{"t2", []int{}, []int{}},
		{"t3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"t4", []int{11, 25, 12, 22, 64}, []int{11, 12, 22, 25, 64}},
		{"t5", []int{21, 1, 55, 5, 2, 3, 34, 8, 1, 13}, []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
		{"t6", []int{2, 0, -3, -11, -2, -5, -3, -13, -1, -7, 1, 3}, []int{-13, -11, -7, -5, -3, -3, -2, -1, 0, 1, 2, 3}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SelectionSort(tc.in)
			for i, n := range tc.in {
				if n != tc.ex[i] {
					t.Errorf("expected %v; got %v", tc.ex[i], n)
				}
			}
		})
	}
}

func ExampleSelectionSort() {
	a := []int{4, 8, 7, 1, 5, 3, 9, 6, 0, 2}
	SelectionSort(a)
	fmt.Println(a)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}
