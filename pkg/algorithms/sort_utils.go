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
