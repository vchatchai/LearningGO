package inputoutput

import (
	"fmt"
	"bufio"
	"os"
)


func Stdin() { 
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}