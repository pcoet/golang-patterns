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
	"testing"

	"github.com/pcoet/golang-patterns/pkg/datastructs"
)

func TestDFS(t *testing.T) {
	g := datastructs.CreateGraph(13)
	g.AddEdge(8, 2)
	g.AddEdge(2, 10)
	g.AddEdge(8, 11)
	g.AddEdge(11, 4)
	g.AddEdge(8, 6)
	g.AddEdge(6, 5)
	g.AddEdge(5, 3)
	g.AddEdge(3, 12)
	g.AddEdge(8, 1)
	g.AddEdge(1, 7)
	g.AddEdge(1, 9)
	g.AddEdge(7, 9)
	g.AddEdge(9, 12)
	want := []int{8, 2, 10, 11, 4, 6, 5, 3, 12, 9, 1, 7}
	var got []int
	DFS(g, 8, func(v int) {
		got = append(got, v)
	})
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("expected %v; got %v", want[i], got[i])
		}
	}
}

func ExampleDFS() {
	g := datastructs.CreateGraph(13)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(0, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.AddEdge(5, 8)
	g.AddEdge(8, 9)
	g.AddEdge(0, 10)
	g.AddEdge(10, 11)
	g.AddEdge(11, 12)
	DFS(g, 0, func(v int) {
		fmt.Println(v)
	})
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
	// 12
}
