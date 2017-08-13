package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value Altcoin
	Left  *Node
	Right *Node
}

type Altcoin struct {
	Name  string
	Price float64
}

func (n *Node) Insert(a Altcoin) error {
	if n == nil {
		return errors.New("Can't insert a value")
	}

	switch {
	case a.Name == n.Value.Name:
		return nil
	case a.Name < n.Value.Name:
		if n.Left == nil {
			n.Left = &Node{Value: a}
		}
		return n.Left.Insert(a)
	case a.Name > n.Value.Name:
		if n.Right == nil {
			n.Right = &Node{Value: a}
		}
		return n.Right.Insert(a)
	}

	return nil
}

func (n *Node) Find(s string) (float64, bool) {
	if n == nil {
		return 0, false
	}

	switch {
	case s == n.Value.Name:
		return n.Value.Price, true
	case s < n.Value.Name:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("ReplaceNode() not allowed on nil node")
	}
	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (n *Node) Delete(s string, parent *Node) error {
	if n == nil {
		return errors.New("Value does not exist in the binary tree")
	}

	switch {
	case s < n.Value.Name:
		return n.Left.Delete(s, n)
	case s > n.Value.Name:
		return n.Right.Delete(s, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}

		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		replacement, replaceParent := n.Left.findMax(n)

		n.Value = replacement.Value
		n.Value.Price = replacement.Value.Price
		return replacement.Delete(replacement.Value.Name, replaceParent)
	}
}

func main() {

	var head *Node
	head = new(Node)

	Litecoin := Altcoin{Name: "Litecoin", Price: 12.11}
	head.Insert(Litecoin)
	Decred := Altcoin{Name: "Decred", Price: 20.54}
	head.Insert(Decred)

	fmt.Println(head.Right.Value.Price)
	fmt.Println(head.Right.Left.Value.Price)

	head.Delete("Decred", head)

	d := string("Decred")
	_, found := head.Find(d)
	fmt.Println(found)

}
