package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initial2")
}

func deferFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover")
		}
	}()
	fmt.Println("Counting")
	for i := 1; i < 10; i++ {
		defer fmt.Println(i)
	}
	panic("Fail")
	// fmt.Println("Hello")
}
