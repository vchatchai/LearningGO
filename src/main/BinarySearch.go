package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

func BinarySearch(array []int, value int) (int, error) {
	return search(array, value, 0, len(array))

}
func search(array []int, value, startIndex, endIndex int) (int, error) {

	middle := (endIndex - startIndex) / 2
	valueInPosition := array[startIndex+middle]
	if valueInPosition == value {
		return startIndex + middle, nil
	} else if endIndex-startIndex == 1 {
		return 0, errors.New("data not found")
	} else if valueInPosition > value {
		return search(array, value, startIndex, startIndex+middle)
	} else {
		return search(array, value, startIndex+middle, endIndex)
	}
}

func MaxParallelism() {
	// maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	log.Println("numCPU:", numCPU)
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.HeapSys)

	fmt.Println("numCPU", numCPU)
	fmt.Printf("\n%v", mem.TotalAlloc)
}
