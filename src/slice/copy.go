package slice

import "fmt"

func Copy() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	// copy(a6, a4)

	copy(a4, a6)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	fmt.Println()
	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	fmt.Println("array4:", array4[0:])
	fmt.Println("s6:", s6)
	copy(s6, array4[0:])
	fmt.Println("array4:", array4[0:])
	fmt.Println("s6:", s6)
	fmt.Println()
	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, -7, 7, -7, 7}

	// copy(array5[0:], s7)
	copy(array5[0:], s7)
	fmt.Println("array5:", array5)
	fmt.Println("s7:", s7)
	copy(array5[0:], s7)
	fmt.Println("array5:", array5)
	fmt.Println("s7:", s7)
}
