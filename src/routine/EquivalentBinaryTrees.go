package routine

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) String() string {
	// var left string = ""
	// if t.Left == nil {
	// 	return fmt.Sprintf(" v:%v  left:nil ", t.Value)
	// }

	// if t.Right == nil {
	// 	return fmt.Sprintf(" v:%v  right:nil ", t.Value)
	// }
	// return fmt.Sprintf(" v:%v  left:%v right:%v ", t.Value, t.Left, t.Right)
	return fmt.Sprintf(" v:%v  ", t.Value)
}

func (t *Tree) setValue(v int) Tree {
	if t.Left == nil {
		t.Left = &Tree{Value: v}
		// fmt.Println(t.Left)
	} else if t.Value == 0 {
		t.Value = v
		// fmt.Println(t.Value)
	} else if t.Right == nil {
		t.Right = &Tree{Value: v}
		// fmt.Println(t.Right)
	} else {
		t = &Tree{Left: t, Value: v}
	}

	return *t
}

func BinaryTree() {
	var t Tree = Tree{}
	trees := []int{1, 1, 2, 3, 5, 8, 13}
	for _, tree := range trees {
		fmt.Println(tree)
		t = t.setValue(tree)
	}

	fmt.Println("Test start.")
	fmt.Println(t)
	t2 := t.Left

	fmt.Println(t)
	t3 := t2.Left

	fmt.Println(t)

	fmt.Println(t2)
	fmt.Println(t3)

	fmt.Println(*t3.Left)

	// for {
	// 	fmt.Println(t)
	// 	if t.Left != nil {
	// 		temp := t.Left
	// 		t = *temp

	// 	} else {
	// 		break
	// 		return
	// 	}

	// }
	fmt.Println("Test done.")
}
