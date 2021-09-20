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

func partition(a []int, lo int, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if Less(a[j], pivot) {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	// Exch(a, a[i], a[hi])
	a[i], a[hi] = a[hi], a[i]
	return i
}

func quicksort(a []int, lo int, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		quicksort(a, lo, p-1)
		quicksort(a, p+1, hi)
	}
}

// Quicksort sorts an array of integers in ascending order.
func Quicksort(a []int) {
	Shuffle(a)
	quicksort(a, 0, len(a)-1)
}
