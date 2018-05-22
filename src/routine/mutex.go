package routine

import (
	"fmt"
	"sync"
	"time"
)

type SaftCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SaftCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()

}

func (c *SaftCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func Mutex() {
	c := SaftCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
	fmt.Println(c)
}
