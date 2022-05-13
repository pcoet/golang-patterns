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
	"fmt"

	"github.com/pcoet/golang-patterns/pkg/datastructs"
)

// Proc is a callback for processing nodes in the graph. It will be invoked once per node.
type Proc func(int)

func validateVertex(v int, V int) {
	if v < 0 || v >= V {
		msg := fmt.Sprintf("vertex %v is not between 0 and %v", v, V-1)
		panic(msg)
	}
}

func dfs(g datastructs.Graph, v int, marked []bool, cb Proc) {
	marked[v] = true
	cb(v)
	for _, w := range g.Adj[v] {
		if !marked[w] {
			dfs(g, w, marked, cb)
		}
	}
}

// DFS performs a depth-first search on graph g, starting at vertex s. It invokes
// a callback function on each discovered vertex.
func DFS(g datastructs.Graph, s int, cb Proc) {
	marked := make([]bool, g.V)
	validateVertex(s, g.V)
	dfs(g, s, marked, cb)
}
