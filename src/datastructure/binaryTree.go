package datastructure

import (
	"bytes"
	"errors"
	"fmt"
)

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func (Tree *Tree) String() string {
	var deep int = 0
	var buffer bytes.Buffer
	for {

		Trees := getTree(deep, 0, Tree)

		if len(Trees) == 0 {
			return buffer.String()
		}
		for _, n := range Trees {
			buffer.WriteString(fmt.Sprintf(" %d ", n.value))
			// buffer.WriteString(" ")
		}
		buffer.WriteString("\n")
		deep++
	}

}

func getTree(deep, count int, current *Tree) []Tree {

	trees := []Tree{}
	if deep == count {
		if current != nil {
			trees = append(trees, *current)
		}
		return trees
	} else {
		count++
		if current.left != nil {
			for _, n := range getTree(deep, count, current.left) {
				trees = append(trees, n)
			}
		}

		if current.right != nil {

			for _, n := range getTree(deep, count, current.right) {
				trees = append(trees, n)

			}
		}
		return trees
	}

}

func (tree *Tree) Add(new *Tree) (err error) {

	// match, n := Tree.findTree(new.value)
	// if match {
	// 	err = errors.New("The value exist in system")
	// 	return
	// }

	if tree.value > new.value {
		if tree.left != nil {
			tree.left.Add(new)
		} else {
			tree.left = new
		}
	}
	if tree.value < new.value {
		if tree.right != nil {
			tree.right.Add(new)
		} else {
			tree.right = new
		}
	}
	err = errors.New("The value exist in system")
	return
}

func (tree *Tree) findTree(value int) (match bool, n *Tree) {
	if tree.value == value {
		return true, tree
	}
	if tree.value > value {
		if tree.left == nil {
			return false, tree
		}
		return tree.left.findTree(value)
	}

	if tree.value < value {
		if tree.right == nil {
			return false, tree
		}
		return tree.right.findTree(value)
	}

	return false, nil
}

func BinaryTree() {

	root := Tree{}
	root.value = 10

	root.Add(&Tree{value: 5})
	root.Add(&Tree{value: 7})
	root.Add(&Tree{value: 2})
	root.Add(&Tree{value: 1})
	root.Add(&Tree{value: 3})
	root.Add(&Tree{value: 4})
	root.Add(&Tree{value: 6})
	root.Add(&Tree{value: 8})
	root.Add(&Tree{value: 9})
	root.Add(&Tree{value: 15})
	root.Add(&Tree{value: 13})
	root.Add(&Tree{value: 17})
	root.Add(&Tree{value: 11})
	root.Add(&Tree{value: 12})
	root.Add(&Tree{value: 14})
	root.Add(&Tree{value: 16})
	root.Add(&Tree{value: 18})
	root.Add(&Tree{value: 19})

	fmt.Println("Done.", root.String())
}
