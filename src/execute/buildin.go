package execute

import (
	"fmt"
	"log"
	"os/exec"
)

func BuildinCommand() {
	cmd := exec.Command("/bin/bash", "-c", "source ~/.profile ")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
	fmt.Println(cmd.Env)
}
