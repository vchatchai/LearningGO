package routine

import (
	"fmt"
	"sync"
	"time"
)

var password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}
func show(c *secret) string {
	c.RWM.RLock()
	time.Sleep(time.Second)
	fmt.Print("Show")
	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.M.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func RWMutex() {
	var waitGroup sync.WaitGroup

	s := secret{password: "1234"}

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass:", show(&s))
		}()
	}

	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()
		Change(&s, "5678")

	}()

	// for i := 0; i < 15; i++ {
	// 	waitGroup.Add(1)
	// 	go func() {
	// 		defer waitGroup.Done()
	// 		fmt.Println("Go Pass:", showWithLock(&s))
	// 	}()
	// }

	waitGroup.Wait()
	fmt.Println("Done")

}
