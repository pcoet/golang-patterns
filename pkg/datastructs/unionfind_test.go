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

package datastructs

import "testing"

func TestNewUnionFind(t *testing.T) {
	n := 3
	uf := NewUnionFind(n)
	if uf.count != n {
		t.Errorf("unexpected count; expected %v; got %v", n, uf.count)
	}
	// make sure each element initially has itself as a parent
	for i := 0; i < n; i++ {
		if uf.parent[i] != i {
			t.Errorf("unexpected parent value; expected %v; got %v", i, uf.parent[i])
		}
	}
	// make sure each element rank is initially set to 0
	for i := 0; i < n; i++ {
		if uf.rank[i] != 0 {
			t.Errorf("unexpected parent value; expected %v; got %v", i, uf.parent[i])
		}
	}
}

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(10)

	if uf.Count() != 10 {
		t.Errorf("expected %v; got %v", 10, uf.Count())
	}

	uf.Union(4, 3)
	if uf.Count() != 9 {
		t.Errorf("expected %v; got %v", 9, uf.Count())
	}
	if uf.Connected(4, 3) != true {
		t.Errorf("expected %v; got %v", true, uf.Connected(4, 3))
	}

	uf.Union(3, 8)
	if uf.Count() != 8 {
		t.Errorf("expected %v; got %v", 8, uf.Count())
	}
	if uf.Connected(8, 3) != true {
		t.Errorf("expected %v; got %v", true, uf.Connected(8, 3))
	}
	if uf.Connected(4, 8) != true {
		t.Errorf("expected %v; got %v", true, uf.Connected(4, 8))
	}

	uf.Union(6, 5)
	if uf.Count() != 7 {
		t.Errorf("expected %v; got %v", 7, uf.Count())
	}
	if uf.Connected(5, 6) != true {
		t.Errorf("expected %v; got %v", true, uf.Connected(5, 6))
	}

	uf.Union(9, 4)
	uf.Union(2, 1)

	if uf.Count() != 5 {
		t.Errorf("expected %v; got %v", 5, uf.Count())
	}
	if uf.Connected(0, 7) != false {
		t.Errorf("expected %v; got %v", false, uf.Connected(0, 7))
	}
	if uf.Connected(8, 9) != true {
		t.Errorf("expected %v; got %v", true, uf.Connected(8, 9))
	}

	uf.Union(5, 0)
	uf.Union(7, 2)
	uf.Union(6, 1)
	uf.Union(1, 0)

	if uf.Count() != 2 {
		t.Errorf("expected %v; got %v", 2, uf.Count())
	}
	if uf.Connected(0, 7) != true {
		t.Errorf("expected %v; got %v", false, uf.Connected(0, 7))
	}
}
