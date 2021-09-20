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
