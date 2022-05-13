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
	validateVertex(s, g.V)
	bfs(g, s, marked, cb)
}
