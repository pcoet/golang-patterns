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
	"reflect"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	g := CreateGraph(11)
	if g.E != 0 {
		t.Errorf("expected %v; got %v", 0, g.E)
	}
	if g.V != 11 {
		t.Errorf("expected %v; got %v", 11, g.V)
	}
	if len(g.Adj) != 11 {
		t.Errorf("expected %v; got %v", 11, g.Adj)
	}

	for _, v := range g.Adj {
		if len(v) != 0 {
			t.Errorf("expected %v; got %v", 0, len(v))
		}
	}
}

func TestAddEdge(t *testing.T) {
	g := CreateGraph(11)
	g.AddEdge(1, 9)
	g.AddEdge(9, 3)
	g.AddEdge(3, 7)
	g.AddEdge(3, 1)
	g.AddEdge(7, 4)
	g.AddEdge(4, 1)
	g.AddEdge(5, 10)
	g.AddEdge(10, 5) // this should be a noop
	g.AddEdge(0, 2)
	g.AddEdge(2, 8)
	g.AddEdge(8, 6)
	g.AddEdge(6, 0)

	if g.V != 11 {
		t.Errorf("expected %v; got %v", 11, g.V)
	}

	if g.E != 11 {
		t.Errorf("expected %v; got %v", 11, g.E)
	}

	testCases := []struct {
		name    string
		adjList []int
	}{
		{"0", []int{2, 6}},
		{"1", []int{9, 3, 4}},
		{"2", []int{0, 8}},
		{"3", []int{9, 7, 1}},
		{"4", []int{7, 1}},
		{"5", []int{10}},
		{"6", []int{8, 0}},
		{"7", []int{3, 4}},
		{"8", []int{2, 6}},
		{"9", []int{1, 3}},
		{"10", []int{5}},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			adjList := g.Adj[i]
			isEq := reflect.DeepEqual(adjList, tc.adjList)
			if isEq != true {
				t.Errorf("expected %v; got %v", tc.adjList, adjList)
			}
			d := g.Degree(i)
			l := len(adjList)
			if d != l {
				t.Errorf("expected %v; got %v", l, d)
			}
		})
	}
}

func TestDegree(t *testing.T) {
	g := CreateGraph(3)
	if g.Degree(0) != 0 {
		t.Errorf("expected %v; got %v", 0, g.Degree(0))
	}
	g.AddEdge(0, 1)
	if g.Degree(0) != 1 {
		t.Errorf("expected %v; got %v", 1, g.Degree(0))
	}

	// trying to add the same edge should be a noop
	g.AddEdge(0, 1)
	if g.Degree(0) != 1 {
		t.Errorf("expected %v; got %v", 1, g.Degree(0))
	}

	// trying to add the same edge with args in a different order should also be a noop
	g.AddEdge(1, 0)
	if g.Degree(0) != 1 {
		t.Errorf("expected %v; got %v", 1, g.Degree(0))
	}
	if g.Degree(1) != 1 {
		t.Errorf("expected %v; got %v", 1, g.Degree(1))
	}

	g.AddEdge(0, 2)
	if g.Degree(0) != 2 {
		t.Errorf("expected %v; got %v", 2, g.Degree(0))
	}
	if g.Degree(2) != 1 {
		t.Errorf("expected %v; got %v", 1, g.Degree(2))
	}
}

func ExampleGraph() {
	g := CreateGraph(11)
	g.AddEdge(1, 9)
	g.AddEdge(9, 3)
	g.AddEdge(3, 7)
	g.AddEdge(3, 1)
	g.AddEdge(7, 4)
	g.AddEdge(4, 1)
	g.AddEdge(5, 10)
	g.AddEdge(10, 5) // this should be a noop
	g.AddEdge(0, 2)
	g.AddEdge(2, 8)
	g.AddEdge(8, 6)
	g.AddEdge(6, 0)
	fmt.Print(g.String())
	// Output:
	// 11 vertices; 11 edges
	// 0: 2 6
	// 1: 9 3 4
	// 2: 0 8
	// 3: 9 7 1
	// 4: 7 1
	// 5: 10
	// 6: 8 0
	// 7: 3 4
	// 8: 2 6
	// 9: 1 3
	// 10: 5
}
