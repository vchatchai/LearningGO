package routine

import (
	"fmt"
	"sync"
)

var forgot sync.Mutex

func function() {
	forgot.Lock()
	fmt.Println("Locked!")
}

func ForgetMutex() {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		function()
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		function()
	}()

	w.Wait()
}
