package routine

import (
	"fmt"
	"sync"
)

func SyncGO(totalRoutine int) {

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)

	for i := 0; i < totalRoutine; i++ {
		go func(x int) {
			waitGroup.Add(1)
			fmt.Printf("%d	", x)
			defer waitGroup.Done()
		}(i)
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting....")
}
