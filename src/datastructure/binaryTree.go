package datastructure

import (
	"bytes"
	"errors"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

// var root = Node{}

func (node *Node) String() string {
	var deep int = 1
	var buffer bytes.Buffer
	for {

		nodes := getNode(deep, 0, node)

		if len(nodes) == 0 {
			return buffer.String()
		}
		for _, n := range nodes {
			buffer.WriteString(string(n.value))
			buffer.WriteString(" ")
		}
		buffer.WriteString("\n")
		deep++
	}

}

func getNode(deep, count int, current *Node) []Node {

	if deep == count {
		return []Node{*current}
	} else {
		nodes := make([]Node, deep)
		count++
		if current.left != nil {
			for _, n := range getNode(deep, count, current.left) {
				nodes = append(nodes, n)
			}
		}

		if current.right != nil {

			for _, n := range getNode(deep, count, current.right) {
				nodes = append(nodes, n)

			}
		}
		return nodes
	}

}

func (node *Node) Add(new *Node) (err error) {

	match, n := node.findNode(new.value)
	if match {
		err = errors.New("The value exist in system")
		return
	}
	if n.value > new.value {
		n.left = new
	}
	if n.value < new.value {
		n.right = new
	}
	return
}

func (node *Node) findNode(value int) (match bool, n *Node) {
	if node.value == value {
		return true, node
	}
	if node.value > value {
		if node.left == nil {
			return false, node
		}
		return node.left.findNode(value)
	}

	if node.value < value {
		if node.right == nil {
			return false, node
		}
		return node.right.findNode(value)
	}

	return false, nil
}

func BinaryTree() {

	root := Node{}
	root.value = 10

	root.Add(&Node{value: 11})
	root.Add(&Node{value: 15})
	root.Add(&Node{value: 9})

	root.Add(&Node{value: 16})
	root.Add(&Node{value: 1})
	root.Add(&Node{value: 5})
	root.Add(&Node{value: 6})
	root.Add(&Node{value: 12})
	root.Add(&Node{value: 4})

	fmt.Println("Done.", root.String())
}
