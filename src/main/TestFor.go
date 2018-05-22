package main

import "fmt"

func testFor() {
	fmt.Println("For Test")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func testForNoInitial() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
