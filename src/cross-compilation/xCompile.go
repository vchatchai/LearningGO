package compilation

import (
	"fmt"
	"runtime"
)

func XCompile() {
	fmt.Println("You are using", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("With GO version", runtime.Version())

}
