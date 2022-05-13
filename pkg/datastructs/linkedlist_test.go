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

import "testing"

func TestNode(t *testing.T) {
	tests := []struct {
		nodeData interface{}
	}{
		{17},
		{nil},
		{"foo"},
	}
	for _, tc := range tests {
		node := Node{tc.nodeData, nil}
		want := tc.nodeData
		got := node.value
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}
func TestNodeToString(t *testing.T) {
	tests := []struct {
		nodeVal    interface{}
		nodeStrVal interface{}
	}{
		{17, "17"},
		{nil, "<nil>"},
		{"foo", "foo"},
	}
	for _, tc := range tests {
		node := Node{tc.nodeVal, nil}
		want := tc.nodeStrVal
		got := node.toString()
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}
func TestLinkedListHead(t *testing.T) {
	tests := []struct {
		nodeData interface{}
	}{
		{17},
		{nil},
		{"foo"},
	}
	for _, tc := range tests {
		node := Node{tc.nodeData, nil}
		list := LinkedList{&node}
		want := tc.nodeData
		got := list.head.value
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}
func TestInsertAfter(t *testing.T) {
	// create a list with three nodes
	node3 := Node{17, nil}
	node2 := Node{11, &node3}
	node1 := Node{7, &node2}
	list := LinkedList{&node1}
	// create a new node and insert it after node2
	newNode := Node{13, nil}
	list.insertAfter(&node2, &newNode)
	// test that the new node comes after node2
	want := &newNode
	got := node2.next
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
	// test that the new node comes before node3
	want = &node3
	got = newNode.next
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
func TestInsertHead(t *testing.T) {
	// create a list with three nodes
	node3 := Node{17, nil}
	node2 := Node{13, &node3}
	node1 := Node{11, &node2}
	list := LinkedList{&node1}
	// create a new node and insert it at the head of the list
	newNode := Node{7, nil}
	list.insertHead(&newNode)
	// test that the new node is at the head of the list
	want := &newNode
	got := list.head
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
	// test that the new node comes before node1
	want = &node1
	got = newNode.next
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
func TestRemoveAfter(t *testing.T) {
	// create a list with three nodes
	node3 := Node{17, nil}
	node2 := Node{13, &node3}
	node1 := Node{11, &node2}
	list := LinkedList{&node1}
	// remove node2
	list.removeAfter(&node1)
	// test that node3 comes after node1
	want := &node3
	got := node1.next
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
	// try to remove non-existent node; expect no error.
	list.removeAfter(&node3)
	want = nil
	got = node3.next
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
func TestRemoveHead(t *testing.T) {
	// create a list with three nodes
	node3 := Node{17, nil}
	node2 := Node{13, &node3}
	node1 := Node{11, &node2}
	list := LinkedList{&node1}
	// remove head node
	list.removeHead()
	// test that node2 is at the head
	want := &node2
	got := list.head
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
