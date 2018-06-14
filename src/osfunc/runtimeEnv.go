package osfunc

import (
	"fmt"
	"runtime"
)

func RuntimeEnv() {
	fmt.Println("You are using", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutine:", runtime.NumGoroutine())
}
