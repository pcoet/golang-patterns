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

// Bag represents a bag (or multiset) of integers. It supports insertion and iteration.
type Bag struct {
	a []int // array of items
}

// IsEmpty returns true if the bag is empty; false otherwise.
func (b *Bag) IsEmpty() bool {
	return len(b.a) == 0
}

// Size returns the number of items in the bag.
func (b *Bag) Size() int {
	return len(b.a)
}

// Add adds an item to the bag.
// Note that the built-in append function runs in amortized O(1) time, i.e. it's performant.
func (b *Bag) Add(item int) {
	b.a = append(b.a, item)
}

// Iterator returns a function for iterating over the items in the bag.
func (b *Bag) Iterator() func() (int, bool) {
	i := 0
	return func() (int, bool) {
		if i >= b.Size() {
			return 0, false
		}
		n := b.a[i]
		i = i + 1
		return n, true
	}
}
