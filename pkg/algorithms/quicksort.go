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
