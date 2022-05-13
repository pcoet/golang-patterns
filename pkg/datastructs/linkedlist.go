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

import "fmt"

// Node represents an element of a linked list.
type Node struct {
	value interface{}
	next  *Node
}

func (n *Node) toString() string {
	return fmt.Sprintf("%v", n.value)
}

// LinkedList implements a singly linked list data structure.
type LinkedList struct {
	head *Node
}

// insertAfter inserts newNode after node.
func (l *LinkedList) insertAfter(node *Node, newNode *Node) {
	newNode.next = node.next
	node.next = newNode
}

// insertHead inserts newNode at the head of the list.
func (l *LinkedList) insertHead(newNode *Node) {
	newNode.next = l.head
	l.head = newNode
}

// removeAfter removes the node following node. If there is no next node, it does nothing.
func (l *LinkedList) removeAfter(node *Node) {
	if node.next != nil {
		node.next = node.next.next
	}
}

// removeHead removes the head node.
func (l *LinkedList) removeHead() {
	if l.head != nil {
		l.head = l.head.next
	}
}
