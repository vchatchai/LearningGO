package compile

import "fmt"

func init() {
	fmt.Println("init cPackage")
}

func A() {
	fmt.Println("This is function A!")
}

const MyConstant = 123
const privateConstant = 21

/*
a.go
	package a
	import "fmt"
	func init() {
		fmt.Println("a")
	}

b.go
	package b
	import (
		"fmt"
	)
	func init() {
		fmt.Println("b")
	}

c.go
	package c
	import "fmt"
	func init() {
		fmt.Println("c")
	}

main.go
	package main
	import (
		_ "a"
		_ "c"
		_ "b"
		"fmt"
	)
func main(){
}

*/
