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
