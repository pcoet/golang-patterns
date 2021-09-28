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
