package benchmarks

import "fmt"

func f1() int {

	fmt.Println("Enter f1()")
	return -10
	fmt.Println("Exiting f1()")
	return -1
}
func f2() int {
	if true {
		return 10
	}
	fmt.Println("Exiting f2()")
	return 0
}

func CannotReach() {
	fmt.Println(f1())
	fmt.Println("Exit program")
}
