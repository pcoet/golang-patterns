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
