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
	"strconv"
)

// Key is an integer key in a symbol table.
type Key int

// Value is a string value in a symbol table.
type Value string

type node struct {
	key         Key
	val         Value
	left, right *node
	size        int // number of nodes rooted at this node (i.e. number of nodes in the subtree)
}

// BST represents an ordered symbol table of int/string key-value pairs.
type BST struct {
	root *node // root of the BST
}

func (b *BST) size(x *node) int {
	if x == nil {
		return 0
	}
	return x.size
}

// Size returns the number of key-value pairs in the symbol table.
func (b *BST) Size() int {
	return b.size(b.root)
}

// IsEmpty returns true if the symbol table is empty, and false otherwise.
func (b *BST) IsEmpty() bool {
	return b.Size() == 0
}

func (b *BST) get(x *node, key Key) (Value, bool) {
	if x == nil {
		return "", false
	}
	if key < x.key {
		return b.get(x.left, key)
	}
	if key > x.key {
		return b.get(x.right, key)
	}
	return x.val, true
}

// Get returns the value associated with the given key.
func (b *BST) Get(key Key) (Value, bool) {
	return b.get(b.root, key)
}

// Contains returns true if the given key is in the symbol table; false otherwise.
func (b *BST) Contains(key Key) bool {
	_, ok := b.Get(key)
	return ok
}

func (b *BST) put(x *node, key Key, val Value) *node {
	if x == nil {
		return &node{key: key, val: val, size: 1}
	}

	if key < x.key {
		x.left = b.put(x.left, key, val)
	} else if key > x.key {
		x.right = b.put(x.right, key, val)
	} else {
		x.val = val
	}

	x.size = 1 + b.size(x.left) + b.size(x.right)
	return x
}

// Put inserts the specified key-value pair into the symbol table.
// If the key already exists, overwrites the old value with the new value.
func (b *BST) Put(key Key, val Value) {
	if val == "" {
		b.Delete(key)
		return
	}
	b.root = b.put(b.root, key, val)
}

func (b *BST) deleteMin(x *node) *node {
	if x.left == nil {
		return x.right
	}
	x.left = b.deleteMin(x.left)
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

// DeleteMin removes the smallest key and associated value from the symbol table.
func (b *BST) DeleteMin() {
	if b.IsEmpty() {
		panic("Symbol table underflow")
	}
	b.root = b.deleteMin(b.root)
}

func (b *BST) deleteMax(x *node) *node {
	if x.right == nil {
		return x.left
	}
	x.right = b.deleteMax(x.right)
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

// DeleteMax removes the largest key and associated value from the symbol table.
func (b *BST) DeleteMax() {
	if b.IsEmpty() {
		panic("Symbol table underflow")
	}
	b.root = b.deleteMax(b.root)
}

func (b *BST) delete(x *node, key Key) *node {
	if x == nil {
		return nil
	}
	if key < x.key {
		x.left = b.delete(x.left, key)
	} else if key > x.key {
		x.right = b.delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		t := x
		x = b.min(t.right)
		x.right = b.deleteMin(t.right)
		x.left = t.left
	}
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

// Delete removes the specified key and its associated value from the symbol table.
func (b *BST) Delete(key Key) {
	b.root = b.delete(b.root, key)
}

func (b *BST) min(x *node) *node {
	if x.left == nil {
		return x
	}
	return b.min(x.left)
}

// Min returns the smallest key in the symbol table.
func (b *BST) Min() Key {
	if b.IsEmpty() {
		panic("calls Min() with empty symbol table")
	}
	return b.min(b.root).key
}

func (b *BST) max(x *node) *node {
	if x.right == nil {
		return x
	}
	return b.max(x.right)
}

// Max returns the largest key in the symbol table.
func (b *BST) Max() Key {
	if b.IsEmpty() {
		panic("calls Max() with empty symbol table")
	}
	return b.max(b.root).key
}

func (b *BST) floor(x *node, key Key) *node {
	if x == nil {
		return nil
	}
	if key == x.key {
		return x
	}
	if key < x.key {
		return b.floor(x.left, key)
	}

	t := b.floor(x.right, key)
	if t != nil {
		return t
	}
	return x
}

// Floor returns the largest key in the symbol table less than or equal to `key`.
func (b *BST) Floor(key Key) Key {
	if b.IsEmpty() {
		panic("calls Floor() with empty symbol table")
	}
	x := b.floor(b.root, key)
	if x == nil {
		panic("argument to Floor() is too small")
	}
	return x.key
}

func (b *BST) ceiling(x *node, key Key) *node {
	if x == nil {
		return nil
	}
	if key == x.key {
		return x
	}
	if key < x.key {
		t := b.ceiling(x.left, key)
		if t != nil {
			return t
		}
		return x
	}
	return b.ceiling(x.right, key)
}

// Ceiling returns the smallest key in the symbol table greater than or equal to `key`.
func (b *BST) Ceiling(key Key) Key {
	if b.IsEmpty() {
		panic("calls Ceiling() with empty symbol table")
	}
	x := b.ceiling(b.root, key)
	if x == nil {
		panic("argument to Ceiling() is too large")
	}
	return x.key
}

func (b *BST) selectKey(x *node, rank int) (Key, bool) {
	if x == nil {
		return 0, false
	}
	leftSize := b.size(x.left)
	if leftSize > rank {
		return b.selectKey(x.left, rank)
	}
	if leftSize < rank {
		return b.selectKey(x.right, rank-leftSize-1)
	}
	return x.key, true
}

// Select returns the key of a given rank in the symbol table. This key has the
// property that there are `rank` keys in the symbol table that are smaller.
func (b *BST) Select(rank int) Key {
	if rank < 0 || rank >= b.Size() {
		panic("argument to Select() is invalid: " + strconv.Itoa(rank))
	}
	n, _ := b.selectKey(b.root, rank)
	return n
}

func (b *BST) rank(key Key, x *node) int {
	if x == nil {
		return 0
	}
	if key < x.key {
		return b.rank(key, x.left)
	}
	if key > x.key {
		return 1 + b.size(x.left) + b.rank(key, x.right)
	}
	return b.size(x.left)
}

// Rank returns the number of keys in the symbol table strictly less than the specified key.
func (b *BST) Rank(key Key) int {
	return b.rank(key, b.root)
}

func (b *BST) keysInRange(x *node, queue *[]Key, lo Key, hi Key) {
	if x == nil {
		return
	}
	if lo < x.key {
		b.keysInRange(x.left, queue, lo, hi)
	}
	if (lo <= x.key) && (hi >= x.key) {
		*queue = append(*queue, x.key)
	}
	if hi > x.key {
		b.keysInRange(x.right, queue, lo, hi)
	}
}

// KeysInRange returns all keys in the symbol table in the given range.
func (b *BST) KeysInRange(lo Key, hi Key) []Key {
	queue := []Key{}
	b.keysInRange(b.root, &queue, lo, hi)
	return queue
}

// Keys returns all keys in the symbol table.
func (b *BST) Keys() []Key {
	if b.IsEmpty() {
		return []Key{}
	}
	return b.KeysInRange(b.Min(), b.Max())
}

// SizeOfRange returns the number of keys in the symbol table in the given range.
func (b *BST) SizeOfRange(lo Key, hi Key) int {
	if lo > hi {
		return 0
	}
	if b.Contains(hi) {
		return b.Rank(hi) - b.Rank(lo) + 1
	}
	return b.Rank(hi) - b.Rank(lo)
}
