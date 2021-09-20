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
