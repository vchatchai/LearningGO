package routine

import (
	"fmt"
	"sync"
	"time"
)

var (
	m  sync.Mutex
	v1 int
)

func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
}

func readMonitor() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func Mutex2() {
	var waitGroup sync.WaitGroup

	numGR := 21
	fmt.Printf("%d", readMonitor())
	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf("-> %d", readMonitor())
		}(i)
	}
	waitGroup.Wait()
	fmt.Printf("->%d\n", readMonitor())

}
