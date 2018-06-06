package routine

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()

}

func printNumbers2() {
	for i := 0; i < 10; i++ {
		// time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}

}

func printLetters2() {
	for i := 'A'; i < 'A'+10; i++ {
		// time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
}

func Print2() {
	printNumbers2()
	printLetters2()
}

func goPrint2() {
	go printNumbers2()
	go printLetters2()
	time.Sleep(1 * time.Microsecond)
}

func printNumbers3(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		// time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()

}

func printLetters3(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		// time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func goPrint3() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go printNumbers3(&wg)
	go printLetters3(&wg)
	wg.Wait()
}
