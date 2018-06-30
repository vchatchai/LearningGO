package syscaller

import (
	"fmt"
	"os"
	"syscall"
)

func SysCallUnixCmd() {
	fmt.Println("Using syscall.Exec()")
	command := "/bin/ls"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
