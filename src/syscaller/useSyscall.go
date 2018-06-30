package syscaller

import (
	"fmt"
	"syscall"
)

func SysCallUnix() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("My pid is", pid)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("User ID:", uid)
	time, _, _ := syscall.Syscall(13, 0, 0, 0)
	fmt.Println("System tie:", time)
	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	fd := 1
	syscall.Write(fd, message)
}
