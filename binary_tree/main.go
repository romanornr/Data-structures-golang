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

func main() {

	var head *Node
	head = new(Node)

	Litecoin := Altcoin{Name: "Litecoin", Price: 12.11}
	head.Insert(Litecoin)
	Decred := Altcoin{Name: "Decred", Price: 20.54}
	head.Insert(Decred)

	fmt.Println(head.Right.Value.Price)
	fmt.Println(head.Right.Left.Value.Price)
	d := string("Decred")
	var a float64
	a, found := head.Find(d)
	fmt.Println(a)
	fmt.Println(found)
}
