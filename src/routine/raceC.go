package routine

import (
	"fmt"
	"sync"
)

// run main with flag --race
func RacingCondition(numGR int) {
	var waitGroup sync.WaitGroup
	var i int
	var mux sync.Mutex

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(j int) {
			defer waitGroup.Done()
			mux.Lock()
			k[j] = j
			mux.Unlock()

		}(i)
	}
	k[2] = 10
	waitGroup.Wait()
	fmt.Printf("k = %#v\n", k)
}
