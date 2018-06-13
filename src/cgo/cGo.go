package cgo

// #include<stdio.h>
// void callC(){
// 	printf("Calling C code!\n");
// }
import "C"
import "fmt"

func CGO() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}
