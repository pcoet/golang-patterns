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

package datastructs

// MaxPQ represents a max priority queue. It's implemented with an unsorted array.
type MaxPQ struct {
	pq []int // store items at indices 1 to n
	n  int   // number of items in the queue
}

func (q *MaxPQ) less(i int, j int) bool {
	return q.pq[i] < q.pq[j]
}

func (q *MaxPQ) exch(i int, j int) {
	swap := q.pq[i]
	q.pq[i] = q.pq[j]
	q.pq[j] = swap
}

func (q *MaxPQ) swim(k int) {
	for k > 1 && q.less(k/2, k) {
		q.exch(k, k/2)
		k = k / 2
	}
}

func (q *MaxPQ) sink(k int) {
	for 2*k <= q.n {
		j := 2 * k
		if j < q.n && q.less(j, j+1) {
			j++
		}
		if !q.less(k, j) {
			break
		}
		q.exch(k, j)
		k = j
	}
}

// IsEmpty returns true if the priority queue is empty; false otherwise.
func (q *MaxPQ) IsEmpty() bool {
	return q.n == 0
}

// Size returns the number of items in the priority queue.
func (q *MaxPQ) Size() int {
	return q.n
}

// Insert adds an item to the priority queue.
func (q *MaxPQ) Insert(x int) {
	// if there are no items on the array, first add a placeholder 0, for indexing at 1
	if len(q.pq) == 0 {
		q.pq = append(q.pq, 0)
	}
	q.pq = append(q.pq, x)
	q.n = q.n + 1
	q.swim(q.n)
}

// Max returns the largest item in the priority queue.
func (q *MaxPQ) Max() int {
	if q.IsEmpty() {
		panic("cannot return max from an empty queue")
	}
	return q.pq[1]
}

// DelMax removes and returns the largest item in the priority queue.
func (q *MaxPQ) DelMax() int {
	if q.IsEmpty() {
		panic("cannot return max from an empty queue")
	}
	max := q.pq[1]
	q.exch(1, q.n)
	q.n = q.n - 1
	q.sink(1)
	q.pq = q.pq[:len(q.pq)-1]
	return max
}
