package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value string
	Data  int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value string, data int) error {
	if n == nil {
		return errors.New("Can't insert a value into a nil tree")
	}

	switch {
	case value == n.Value:
		return nil

	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
		}
		return n.Left.Insert(value, data)

	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func main() {

	//head.val = 1
	//insert(head, 2)
	var head *Node
	head = new(Node)

	head.Insert("Litecoin", 12)
	fmt.Println(head.Right.Data)

	//	fmt.Println(head.val)
	//fmt.Println(head.right.val)

}
