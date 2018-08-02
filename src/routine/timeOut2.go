package routine

import (
	"fmt"
	"sync"
	"time"
)

func timeoutInternal(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}
func Timeout2() {
	var w sync.WaitGroup
	w.Add(1)
	duration := time.Duration(int32(10)) * time.Millisecond

	t := timeoutInternal(&w, duration)

	if t {
		fmt.Println("Timed Out!")
	} else {
		fmt.Println("OK!")
	}

	w.Done()

}
