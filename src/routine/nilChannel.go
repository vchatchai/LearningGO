package routine

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)
	for {
		time.Sleep(200 * time.Microsecond)
		select {
		case input := <-c:

			sum = sum + input
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func NilChannel() {
	c := make(chan int)
	go add(c)
	go send(c)
	time.Sleep(30 * time.Second)
}
