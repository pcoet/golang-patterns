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

import "fmt"

// UnionFind is a data structure that stores a collection of disjoint (non-overlapping) sets.
type UnionFind struct {
	parent []int  // parent[i] = parent of i
	rank   []byte // rank[i] = rank of subtree rooted at i (never more than 31)
	count  int    // number of components
}

// NewUnionFind returns a new instance of UnionFind with n elements,
// each of which is in its own set.
func NewUnionFind(n int) *UnionFind {
	if n < 0 {
		panic("n cannot be negative")
	}
	uf := UnionFind{parent: make([]int, n), rank: make([]byte, n), count: n}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 0
	}
	return &uf
}

// Find returns the canonical element (i.e. the root) of the set containing element p.
func (uf *UnionFind) Find(p int) int {
	uf.validate(p)
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]] // path compression by halving
		p = uf.parent[p]
	}
	return p
}

// Count returns the number of sets.
func (uf *UnionFind) Count() int {
	return uf.count
}

// Connected returns true if p and q are in the same set.
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Union merges the set containing element p with the set containing element q.
func (uf *UnionFind) Union(p int, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	// make root of smaller rank point to root of larger rank
	if uf.rank[rootP] < uf.rank[rootQ] {
		uf.parent[rootP] = rootQ
	} else if uf.rank[rootP] > uf.rank[rootQ] {
		uf.parent[rootQ] = rootP
	} else {
		uf.parent[rootQ] = rootP
		uf.rank[rootP]++
	}
	uf.count--
}

func (uf *UnionFind) validate(p int) {
	n := len(uf.parent)
	if p < 0 || p >= n {
		panic(fmt.Sprintf("index %v is not between 0 and %v", p, n-1))
	}
}
