package routine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan<- int, end chan bool) {
	defer waitGroup.Done()
	for {

		select {
		case <-time.After(4 * time.Second):
			fmt.Println("\ngen time.After()!")
			return
		case createNumber <- rand.Intn(max-min) + min:
			time.Sleep(1 * time.Second)
		case <-end:
			close(end)
			return

		}
	}
}

func print(createNumber chan int) {
	defer waitGroup.Done()
	for {
		select {
		case x := <-createNumber:
			fmt.Println("wait print. ", x)
		case <-time.After(2 * time.Second):
			fmt.Println("\nprint time.After()!")
			return
		}
	}

}

var waitGroup sync.WaitGroup

func SelectGen() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	waitGroup.Add(1)
	go gen(0, 100, createNumber, end)
	go print(createNumber)
	waitGroup.Wait()
	// end <- true
	fmt.Println("Exitting...")
}
