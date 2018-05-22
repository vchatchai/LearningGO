package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 5, 7, 8, 9}
	value := 3
	expectIndex := 6

	index, ok := BinarySearch(array, value)

	if ok != nil {
		fmt.Println(" ok", ok)
	}

	if index != expectIndex {
		t.Error("   is expected value is not equal result value  ", index)
	}

}

func TestMaxParallelism(t *testing.T) {
	fmt.Println("TestMaxParallelism Start")
	MaxParallelism()
	fmt.Println("TestMaxParallelism Done")
}
