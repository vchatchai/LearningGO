package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	var err error

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			return nil
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		fmt.Print(line)
	}
	return nil
}
