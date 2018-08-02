package routine

import (
	"fmt"
	"sync"
	"time"
)

func A(a, b chan struct{}) {

	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
	orderWaitGroup.Done()
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
	orderWaitGroup.Done()
}
func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")

	orderWaitGroup.Done()
}

var orderWaitGroup sync.WaitGroup

func DefineOrder() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})
	orderWaitGroup.Add(1)
	go C(z)
	orderWaitGroup.Add(1)
	go A(x, y)
	orderWaitGroup.Add(1)
	go C(z)
	orderWaitGroup.Add(1)
	go B(y, z)
	orderWaitGroup.Add(1)
	go C(z)
	close(x)
	orderWaitGroup.Wait()
}
