package files

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Save() (err error) {

	s := []byte("Data to write\n")
	f1, err := os.Create("f1.txt")
	if err != nil {
		return
	}
	defer f1.Close()

	fmt.Fprintf(f1, string(s))

	f2, err := os.Create("f2.txt")
	if err != nil {
		return
	}
	defer f2.Close()

	n, err := f2.WriteString(string(s))
	if err != nil {
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("f3.txt")
	if err != nil {
		return
	}
	w := bufio.NewWriter(f3)

	n, err = w.WriteString(string(s))
	if err != nil {
		return
	}
	fmt.Printf("wrote %d bytes \n", n)
	w.Flush()

	f4 := "f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		return
	}

	f5, err := os.Create("f5.txt")
	if err != nil {
		return
	}

	n, err = io.WriteString(f5, string(s))
	if err != nil {
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	return
}
