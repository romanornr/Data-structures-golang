package main

import "fmt"

type value struct {
	Name  string
	Price int
}

type Node struct {
	Value      //Embedded struct
	next, prev *Node
}

type List struct {
	head, tail *Node
}

func (l *List) First() *Node {
	return n.Next
}

func (n *Node) Next() *Node {
	return n.Next
}

func (n *Node) Prev() *Node {
	return n.prev
}

// create node with value
func (l *List) Push(v Value) *List {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n
	return l
}

func (l *List) Find(name string) *Node {
	found := false
	var x *Node = nil
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.Value.Name {
			found = true
			x = n
		}
	}
	return x
}

func (l *List) Delete(name string) bool {
	success := false
	node2del := l.Find(name)
	if node2del != nil {
		fmt.Println("Delete - FOUND: ", name)
		prev_node := node2del.prev
		next_node := node2del.next
	}
	return success
}
