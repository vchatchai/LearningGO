package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5999; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	w.Done()
}

var w = sync.WaitGroup{}

func runGoRoutine() {
	w.Add(1)
	go say("world")
	say("hello")
	w.Wait()

	fmt.Println("Go Routine done.")
}

// func chanRountine(ch chan int){
// 	ch <- rand.Intn(100).
// }

// func channel() {
// 	ch := make(chan int, 2)

// 	go chanRountine(ch)
// 	go chanRountine(ch)

// 	for c := range ch{
// 		fmt.Println(c)
// 	}
// }
