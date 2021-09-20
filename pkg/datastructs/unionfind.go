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
