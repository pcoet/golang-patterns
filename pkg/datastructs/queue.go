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

// Queue represents a first-in-first-out (FIFO) collection of ints.
type Queue struct {
	a []int
}

// IsEmpty returns true if the queue is empty; false otherwise.
func (q *Queue) IsEmpty() bool {
	return len(q.a) == 0
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	return len(q.a)
}

// Enqueue adds an item to the queue.
func (q *Queue) Enqueue(n int) {
	q.a = append(q.a, n)
}

// Dequeue removes and returns the least recently added item on the queue.
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("cannot dequeue from an empty queue")
	}
	var n int
	n, q.a = q.a[0], q.a[1:]
	return n
}

// Peek returns the least recently added item, without removing it.
func (q *Queue) Peek() int {
	if q.IsEmpty() {
		panic("cannot peek into an empty queue")
	}
	return q.a[0]
}

// Iterator returns a function for iterating over the items in the queue.
func (q *Queue) Iterator() func() (int, bool) {
	i := 0
	return func() (int, bool) {
		if i >= q.Size() {
			return 0, false
		}
		n := q.a[i]
		i = i + 1
		return n, true
	}
}
