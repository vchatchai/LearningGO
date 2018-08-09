package routine

import (
	"fmt"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
	worker int
	period time.Duration
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup, workerId int) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square, workerId, time.Duration(random(0, 100))}
		data <- output
		time.Sleep(time.Millisecond * output.period)
	}
	w.Done()
}
func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w, i)

	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func WorkerPool(workers, jobs int) {

	go create(jobs)
	finished := make(chan interface{})

	go func() {
		for d := range data {
			fmt.Printf("Worker ID: %d\t: ", d.worker)
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("Period: %d\tint: ", d.period)
			fmt.Printf(" %d tsquare: %d\n", d.job.integer, d.square)
		}

		finished <- true
	}()

	makeWP(workers)
	fmt.Printf(": %v\n", <-finished)

}
