package routine

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

/**
Channel
*/
func Channel() {
	fmt.Println("Channel start")
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	// for i := 0; i < 10; i++ {
	// 	result := <-c
	// 	fmt.Println("result", result)
	// }

	time.Sleep(time.Second * 10)
	fmt.Println("Channel done.")

}
