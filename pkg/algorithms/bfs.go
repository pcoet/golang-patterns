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
	"github.com/pcoet/golang-patterns/pkg/datastructs"
)

func bfs(g datastructs.Graph, s int, marked []bool, cb Proc) {
	q := []int{}
	cb(s) // invoke the callback on the source vertex
	marked[s] = true
	q = append(q, s)

	for len(q) != 0 {
		var v int
		v, q = q[0], q[1:] // dequeue

		for _, w := range g.Adj[v] {
			if !marked[w] {
				cb(w) // invoke the callback on the current vertex
				marked[w] = true
				q = append(q, w)
			}
		}

	}
}

// BFS performs a breadth-first search on graph g, starting at vertex s. It invokes
// a callback function on each discovered vertex.
func BFS(g datastructs.Graph, s int, cb Proc) {
	marked := make([]bool, g.V)
	validateVertex(s, marked)
	bfs(g, s, marked, cb)
}
