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

type BagItem interface{}

// Bag represents a bag (or multiset) of integers. It supports insertion and iteration.
type Bag struct {
	a []BagItem // array of items
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
func (b *Bag) Add(item BagItem) {
	b.a = append(b.a, item)
}

// Iterator returns a function for iterating over the items in the bag.
func (b *Bag) Iterator() func() (BagItem, bool) {
	i := 0
	return func() (BagItem, bool) {
		if i >= b.Size() {
			return 0, false
		}
		item := b.a[i]
		i = i + 1
		return item, true
	}
}
