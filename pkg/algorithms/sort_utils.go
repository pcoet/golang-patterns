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

import (
	"math/rand"
	"sort"
	"time"
)

// Less returns true if v is less than w; false otherwise.
func Less(v int, w int) bool {
	return v < w
}

// Exch swaps the item at index i with the item at index j.
func Exch(a []int, i int, j int) {
	swap := a[i]
	a[i] = a[j]
	a[j] = swap
}

// Shuffle rearranges the array in a pseudo-random order.
// See also: https://pkg.go.dev/math/rand#Shuffle
func Shuffle(a []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

// Sort sorts a slice of ints in place in ascending order.
func Sort(a []int) {
	sort.Ints(a)
}

// IsSorted returns true if the slice of ints is already sorted; false otherwise.
func IsSorted(a []int) bool {
	return sort.IntsAreSorted(a)
}
