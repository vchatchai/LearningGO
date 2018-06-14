package deferer

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
	fmt.Println()
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	fmt.Println()
}

func Deferer() {
	d1()
	d2()
	d3()
}
