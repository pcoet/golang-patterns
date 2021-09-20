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

package algorithms

import (
	"fmt"
	"testing"
)

func TestMergesort(t *testing.T) {
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
			Mergesort(tc.in)
			for i, n := range tc.in {
				if n != tc.ex[i] {
					t.Errorf("expected %v; got %v", tc.ex[i], n)
				}
			}
		})
	}
}

func ExampleMergesort() {
	a := []int{4, 8, 7, 1, 5, 3, 9, 6, 0, 2}
	Mergesort(a)
	fmt.Println(a)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}
