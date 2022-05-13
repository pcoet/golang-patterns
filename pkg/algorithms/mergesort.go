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
