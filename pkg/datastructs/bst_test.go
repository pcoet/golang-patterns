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

import (
	"math/rand"
	"testing"
	"time"
)

func TestBST(t *testing.T) {
	b := BST{}

	if b.Size() != 0 {
		t.Errorf("expected %v; got %v", 0, b.Size())
	}

	if b.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, b.IsEmpty())
	}

	b.Put(11, "eleven")
	b.Put(17, "seventeen")

	if b.Size() != 2 {
		t.Errorf("expected %v; got %v", 2, b.Size())
	}

	if b.IsEmpty() != false {
		t.Errorf("expected %v; got %v", false, b.IsEmpty())
	}

	if !b.Contains(11) {
		t.Errorf("expected %v; got %v", true, b.Contains(11))
	}

	testCases := []struct {
		name string
		k    Key
		v    Value
	}{
		{"t1", 1, "one"},
		{"t2", 7, "seven"},
		{"t3", 0, "zero"},
		{"t4", 31, "thirty-one"},
		{"t5", 17, "seventeen"}, // overwrite key 17
		{"t6", 3, "three"},
		{"t7", 13, "three"},
		{"t8", 11, "eleven"}, // overwrite key 11
		{"t9", 5, "five"},
		{"t10", 37, "thirty-seven"},
		{"t11", 2, "two"},
		{"t12", 29, "twenty-nine"},
		{"t13", 19, "nineteen"},
		{"t14", 23, "twenty-three"},
	}

	for _, tc := range testCases {
		b.Put(tc.k, tc.v)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := b.Get(tc.k)
			if !(val == tc.v && ok == true) {
				t.Errorf("expected %v and %v; got %v and %v", tc.v, true, val, ok)
			}
		})
	}

	if b.Size() != 14 {
		t.Errorf("expected %v; got %v", 14, b.Size())
	}

	// make sure Get handles a non-existent pair as expected
	v, ok := b.Get(42)
	if !(v == "" && ok == false) {
		t.Errorf("expected %v and %v; got %v and %v", "", false, v, ok)
	}

	if b.Floor(30) != 29 {
		t.Errorf("expected %v; got %v", 29, b.Floor(30))
	}

	if b.Floor(2) != 2 {
		t.Errorf("expected %v; got %v", 2, b.Floor(2))
	}

	if b.Ceiling(30) != 31 {
		t.Errorf("expected %v; got %v", 31, b.Ceiling(30))
	}

	if b.Ceiling(2) != 2 {
		t.Errorf("expected %v; got %v", 2, b.Ceiling(2))
	}

	key := b.Select(0)
	if !(key == 0) {
		t.Errorf("expected %v; got %v", 0, key)
	}

	key = b.Select(5)
	if !(key == 7) {
		t.Errorf("expected %v; got %v", 7, key)
	}

	key = b.Select(13)
	if !(key == 37) {
		t.Errorf("expected %v; got %v", 37, key)
	}

	if b.Rank(0) != 0 {
		t.Errorf("expected %v; got %v", 0, b.Rank(0))
	}

	if b.Rank(1) != 1 {
		t.Errorf("expected %v; got %v", 1, b.Rank(1))
	}

	if b.Rank(17) != 8 {
		t.Errorf("expected %v; got %v", 8, b.Rank(17))
	}

	if b.Rank(37) != 13 {
		t.Errorf("expected %v; got %v", 13, b.Rank(37))
	}

	keys := b.Keys()
	if len(keys) != 14 {
		t.Errorf("expected %v; got %v", 14, len(keys))
	}

	if keys[0] != 0 {
		t.Errorf("expected %v; got %v", 0, keys[0])
	}

	if keys[1] != 1 {
		t.Errorf("expected %v; got %v", 1, keys[1])
	}

	if keys[2] != 2 {
		t.Errorf("expected %v; got %v", 2, keys[2])
	}

	if keys[7] != 13 {
		t.Errorf("expected %v; got %v", 13, keys[7])
	}

	if keys[13] != 37 {
		t.Errorf("expected %v; got %v", 37, keys[13])
	}

	k0 := b.KeysInRange(-2, -1)
	if len(k0) != 0 {
		t.Errorf("expected %v; got %v", 0, len(k0))
	}

	s0 := b.SizeOfRange(-2, -1)
	if s0 != 0 {
		t.Errorf("expected %v; got %v", 0, s0)
	}

	k2 := b.KeysInRange(0, 1)
	if len(k2) != 2 {
		t.Errorf("expected %v; got %v", 2, len(k2))
	}

	s2 := b.SizeOfRange(0, 1)
	if s2 != 2 {
		t.Errorf("expected %v; got %v", 2, s2)
	}

	k8 := b.KeysInRange(0, 13)
	if len(k8) != 8 {
		t.Errorf("expected %v; got %v", 7, len(k8))
	}

	s8 := b.SizeOfRange(0, 13)
	if s8 != 8 {
		t.Errorf("expected %v; got %v", 8, s8)
	}

	min := b.Min()
	if min != 0 {
		t.Errorf("expected %v; got %v", 0, min)
	}

	b.DeleteMin()
	min = b.Min()
	if min != 1 {
		t.Errorf("expected %v; got %v", 1, min)
	}

	b.DeleteMin()
	min = b.Min()
	if min != 2 {
		t.Errorf("expected %v; got %v", 2, min)
	}

	if b.Size() != 12 {
		t.Errorf("expected %v; got %v", 12, b.Size())
	}

	if b.Contains(0) {
		t.Errorf("expected %v; got %v", false, b.Contains(0))
	}

	if b.Contains(1) {
		t.Errorf("expected %v; got %v", false, b.Contains(1))
	}

	max := b.Max()
	if max != 37 {
		t.Errorf("expected %v; got %v", 37, max)
	}

	b.DeleteMax()
	max = b.Max()
	if max != 31 {
		t.Errorf("expected %v; got %v", 31, max)
	}

	b.DeleteMax()
	max = b.Max()
	if max != 29 {
		t.Errorf("expected %v; got %v", 29, max)
	}

	if b.Size() != 10 {
		t.Errorf("expected %v; got %v", 10, b.Size())
	}

	if b.Contains(37) {
		t.Errorf("expected %v; got %v", false, b.Contains(37))
	}

	if b.Contains(31) {
		t.Errorf("expected %v; got %v", false, b.Contains(31))
	}

	b.Delete(13)
	b.Delete(19)

	if b.Size() != 8 {
		t.Errorf("expected %v; got %v", 8, b.Size())
	}

	if b.Contains(13) {
		t.Errorf("expected %v; got %v", false, b.Contains(13))
	}

	if b.Contains(19) {
		t.Errorf("expected %v; got %v", false, b.Contains(19))
	}

	b.Put(7, "")
	if b.Size() != 7 {
		t.Errorf("expected %v; got %v", 7, b.Size())
	}

	if b.Contains(7) {
		t.Errorf("expected %v; got %v", false, b.Contains(7))
	}
}

func CreateBST() BST {
	b := BST{}
	a := []struct {
		name string
		k    Key
		v    Value
	}{

		{"t1", 1, "one"},
		{"t2", 2, "two"},
		{"t3", 3, "three"},
		{"t4", 4, "four"},
		{"t5", 5, "five"},
		{"t6", 6, "six"},
		{"t7", 7, "seven"},
		{"t8", 8, "eight"},
		{"t9", 9, "nine"},
		{"t10", 10, "ten"},
		{"t11", 11, "eleven"},
		{"t12", 12, "twelve"},
		{"t13", 13, "thirteen"},
		{"t14", 14, "fourteen"},
		{"t15", 15, "fifteen"},
		{"t16", 16, "sixteen"},
		{"t17", 17, "seventeen"},
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	for _, t := range a {
		b.Put(t.k, t.v)
	}
	return b
}

func TestSize(t *testing.T) {
	b := CreateBST()
	if b.Size() != 17 {
		t.Errorf("expected %v; got %v", 17, b.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	b := BST{}
	if b.IsEmpty() != true {
		t.Errorf("expected %v; got %v", true, b.IsEmpty())
	}

	b = CreateBST()
	if b.IsEmpty() != false {
		t.Errorf("expected %v; got %v", false, b.IsEmpty())
	}
}

func TestGet(t *testing.T) {
	b := CreateBST()

	v, ok := b.Get(-1)
	if v != "" && ok != false {
		t.Errorf("expected %v and %v; got %v and %v", "", false, v, ok)
	}

	v, ok = b.Get(1)
	if v != "one" && ok != true {
		t.Errorf("expected %v and %v; got %v and %v", "one", true, v, ok)
	}

	v, ok = b.Get(7)
	if v != "seven" && ok != true {
		t.Errorf("expected %v and %v; got %v and %v", "seven", true, v, ok)
	}

	v, ok = b.Get(-7)
	if v != "" && ok != false {
		t.Errorf("expected %v and %v; got %v and %v", "", false, v, ok)
	}

	v, ok = b.Get(11)
	if v != "eleven" && ok != true {
		t.Errorf("expected %v and %v; got %v and %v", "eleven", true, v, ok)
	}

	v, ok = b.Get(17)
	if v != "seventeen" && ok != true {
		t.Errorf("expected %v and %v; got %v and %v", "seventeen", true, v, ok)
	}
}

func TestContains(t *testing.T) {
	b := CreateBST()

	if b.Contains(-1) != false {
		t.Errorf("expected %v; got %v", false, b.Contains(-1))
	}

	if b.Contains(1) != true {
		t.Errorf("expected %v; got %v", true, b.Contains(1))
	}

	if b.Contains(7) != true {
		t.Errorf("expected %v; got %v", true, b.Contains(7))
	}

	if b.Contains(-7) != false {
		t.Errorf("expected %v; got %v", false, b.Contains(-7))
	}

	if b.Contains(11) != true {
		t.Errorf("expected %v; got %v", true, b.Contains(11))
	}

	if b.Contains(17) != true {
		t.Errorf("expected %v; got %v", true, b.Contains(17))
	}
}

func TestPut(t *testing.T) {
	b := CreateBST()

	b.Put(-1, "negative one")
	if !b.Contains(-1) {
		t.Errorf("expected %v; got %v", true, b.Contains(-1))
	}

	b.Put(18, "eighteen")
	if !b.Contains(18) {
		t.Errorf("expected %v; got %v", true, b.Contains(18))
	}

	if b.Size() != 19 {
		t.Errorf("expected %v; got %v", 19, b.Size())
	}

	// test delete op
	b.Put(11, "")
	if b.Contains(11) {
		t.Errorf("expected %v; got %v", false, b.Contains(11))
	}

	if b.Size() != 18 {
		t.Errorf("expected %v; got %v", 18, b.Size())
	}
}

func TestDeleteMin(t *testing.T) {
	b := CreateBST()

	if !b.Contains(1) {
		t.Errorf("expected %v; got %v", false, !b.Contains(1))
	}
	b.DeleteMin()
	if b.Contains(1) {
		t.Errorf("expected %v; got %v", false, b.Contains(1))
	}
}

func TestDeleteMax(t *testing.T) {
	b := CreateBST()

	if !b.Contains(17) {
		t.Errorf("expected %v; got %v", false, !b.Contains(17))
	}
	b.DeleteMax()
	if b.Contains(17) {
		t.Errorf("expected %v; got %v", false, b.Contains(17))
	}
}

func TestDelete(t *testing.T) {
	b := CreateBST()

	if b.Size() != 17 {
		t.Errorf("expected %v; got %v", 17, b.Size())
	}
	b.Delete(2)
	if b.Contains(2) {
		t.Errorf("expected %v; got %v", false, b.Contains(2))
	}
	b.Delete(15)
	if b.Contains(15) {
		t.Errorf("expected %v; got %v", false, b.Contains(15))
	}
	if b.Size() != 15 {
		t.Errorf("expected %v; got %v", 15, b.Size())
	}
}

func TestMin(t *testing.T) {
	b := CreateBST()
	min := b.Min()
	if min != 1 {
		t.Errorf("expected %v; got %v", 1, min)
	}
}

func TestMax(t *testing.T) {
	b := CreateBST()
	max := b.Max()
	if max != 17 {
		t.Errorf("expected %v; got %v", 17, max)
	}
}

func TestFloor(t *testing.T) {
	b := CreateBST()
	f := b.Floor(1)
	if f != 1 {
		t.Errorf("expected %v; got %v", 1, f)
	}

	f = b.Floor(10)
	if f != 10 {
		t.Errorf("expected %v; got %v", 10, f)
	}

	f = b.Floor(20)
	if f != 17 {
		t.Errorf("expected %v; got %v", 17, f)
	}
}

func TestCeiling(t *testing.T) {
	b := CreateBST()
	c := b.Ceiling(1)
	if c != 1 {
		t.Errorf("expected %v; got %v", 1, c)
	}

	c = b.Ceiling(10)
	if c != 10 {
		t.Errorf("expected %v; got %v", 10, c)
	}

	c = b.Ceiling(17)
	if c != 17 {
		t.Errorf("expected %v; got %v", 17, c)
	}
}

func TestSelect(t *testing.T) {
	b := CreateBST()
	n := b.Select(7)
	if n != 8 {
		t.Errorf("expected %v; got %v", 8, n)
	}

	n = b.Select(16)
	if n != 17 {
		t.Errorf("expected %v; got %v", 17, n)
	}
}

func TestRank(t *testing.T) {
	b := CreateBST()
	n := b.Rank(0)
	if n != 0 {
		t.Errorf("expected %v; got %v", 0, n)
	}

	n = b.Rank(10)
	if n != 9 {
		t.Errorf("expected %v; got %v", 9, n)
	}

	n = b.Rank(17)
	if n != 16 {
		t.Errorf("expected %v; got %v", 16, n)
	}
}

func TestKeysInRange(t *testing.T) {
	b := CreateBST()
	a := b.KeysInRange(0, 20)
	if len(a) != 17 {
		t.Errorf("expected %v; got %v", 17, len(a))
	}

	a = b.KeysInRange(5, 15)
	if !(len(a) == 11 && a[0] == 5 && a[10] == 15) {
		t.Errorf("expected %v, %v, and %v; got %v, %v, and %v", 11, 5, 15, len(a), a[0], a[10])
	}
}
