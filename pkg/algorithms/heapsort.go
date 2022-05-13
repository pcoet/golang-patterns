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

func less(pq []int, i int, j int) bool {
	return pq[i-1] < pq[j-1]
}

func exch(pq []int, i int, j int) {
	swap := pq[i-1]
	pq[i-1] = pq[j-1]
	pq[j-1] = swap
}

func sink(pq []int, k int, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && less(pq, j, j+1) {
			j++
		}
		if !(less(pq, k, j)) {
			break
		}
		exch(pq, k, j)
		k = j
	}
}

// Heapsort rearranges the array in ascending order using heapsort.
func Heapsort(pq []int) {
	n := len(pq)

	// heapify the input array
	for k := n / 2; k >= 1; k-- {
		sink(pq, k, n)
	}

	k := n
	for k > 1 {
		exch(pq, 1, k)
		k--
		sink(pq, 1, k)
	}
}
