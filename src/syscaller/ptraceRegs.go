package syscaller

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func Ptrace(command string, args ...string) (err error) {
	var r syscall.PtraceRegs
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		// return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		// return
	}

	wpid := cmd.Process.Pid
	err = syscall.PtraceGetRegs(wpid, &r)
	if err != nil {
		fmt.Println(err)
		// return
	}

	fmt.Printf("Registers: %#v\n", r)
	fmt.Printf("R15=%d, Gs=%d\n", r.R15, r.Gs)

	time.Sleep(2 * time.Second)
	return
}
