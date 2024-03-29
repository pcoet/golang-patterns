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

import (
	"fmt"
	"strings"
)

// Graph represents a simple undirected graph of vertices named 0 through V – 1.
// Multiple edges and self loops are disallowed.
type Graph struct {
	V   int
	E   int
	Adj [][]int
}

// CreateGraph initializes an empty graph with v vertices and 0 edges.
func CreateGraph(v int) Graph {
	if v < 0 {
		panic("number of vertices must be non-negative")
	}
	g := Graph{}
	g.V = v
	g.E = 0
	g.Adj = make([][]int, v)
	return g
}

func (g *Graph) validateVertex(v int) {
	if v < 0 || v >= g.V {
		msg := fmt.Sprintf("vertex %v is not between 0 and %v", v, g.V-1)
		panic(msg)
	}
}

func (g *Graph) edgeExists(v int, w int) bool {
	result := false
	for _, vtx := range g.Adj[v] {
		if vtx == w {
			result = true
		}
	}
	return result
}

// AddEdge adds the undirected edge v-w to the graph.
func (g *Graph) AddEdge(v int, w int) {
	g.validateVertex(v)
	g.validateVertex(w)
	// disallow self loops
	if v == w {
		panic("self loops are not allowed")
	}
	// disallow multiple (or parallel) edges
	if g.edgeExists(v, w) {
		return
	}
	g.E = g.E + 1
	g.Adj[v] = append(g.Adj[v], w)
	g.Adj[w] = append(g.Adj[w], v)
}

// Degree returns the degree of vertex v.
func (g *Graph) Degree(v int) int {
	g.validateVertex(v)
	return len(g.Adj[v])
}

// String returns a string representation of the graph.
func (g *Graph) String() string {
	s := fmt.Sprintf("%v vertices; %v edges\n", g.V, g.E)
	for v := 0; v < g.V; v++ {
		s = s + fmt.Sprintf("%v: ", v)
		for _, w := range g.Adj[v] {
			s = s + fmt.Sprintf("%v ", w)
		}
		s = strings.TrimSpace(s)
		s = s + "\n"
	}

	return s
}
