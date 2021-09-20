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

func merge(a []int, aux []int, lo int, mid int, hi int) {
	// copy to aux
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	// merge back to a
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid { // if the left half is exhauxted...
			a[k] = aux[j]
			j++
		} else if j > hi { // if the right half is exhausted...
			a[k] = aux[i]
			i++
		} else if Less(aux[j], aux[i]) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func mergesort(a []int, aux []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	var mid int = lo + (hi-lo)/2
	mergesort(a, aux, lo, mid)
	mergesort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)

}

// Mergesort rearranges the array in order using a top-down, recursive mergesort algorithm.
func Mergesort(a []int) {
	aux := make([]int, len(a))
	mergesort(a, aux, 0, len(a)-1)
}
