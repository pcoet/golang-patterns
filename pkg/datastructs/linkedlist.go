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
