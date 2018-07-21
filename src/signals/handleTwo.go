package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HandleTwo() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGABRT:
				handleSignal(sig)
				return
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}
}

func handleSignal(signal os.Signal) {
	fmt.Println("handler Signal() Caught:", signal)
}
