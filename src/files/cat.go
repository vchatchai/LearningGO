package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PrintFile(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}

	return nil
}

func Cat() {
	filename := ""
	arguments := os.Args
	if len(arguments) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(arguments); i++ {
		filename = arguments[i]
		err := PrintFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}
