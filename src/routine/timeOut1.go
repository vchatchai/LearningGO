package routine

import (
	"fmt"
	"time"
)

func Timeout() {
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout C1")
	}
}
