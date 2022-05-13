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

func TestBFS(t *testing.T) {
	g := datastructs.CreateGraph(13)
	g.AddEdge(9, 3)
	g.AddEdge(9, 11)
	g.AddEdge(9, 5)
	g.AddEdge(3, 12)
	g.AddEdge(3, 0)
	g.AddEdge(3, 7)
	g.AddEdge(11, 6)
	g.AddEdge(5, 4)
	g.AddEdge(5, 10)
	g.AddEdge(10, 1)
	g.AddEdge(10, 2)
	g.AddEdge(10, 8)
	want := []int{9, 3, 11, 5, 12, 0, 7, 6, 4, 10, 1, 2, 8}
	var got []int
	BFS(g, 9, func(v int) {
		got = append(got, v)
	})
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("expected %v; got %v", want[i], got[i])
		}
	}
}

func ExampleBFS() {
	g := datastructs.CreateGraph(13)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(4, 8)
	g.AddEdge(4, 9)
	g.AddEdge(6, 10)
	g.AddEdge(6, 11)
	g.AddEdge(7, 12)
	BFS(g, 0, func(v int) {
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
