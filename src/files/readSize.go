package files

import (
	"fmt"
	"io"
	"os"
)

func ReadSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer[0:n]
}

func ReadBinary() (err error) {
	f, err := os.Open("readSize.go")
	if err != nil {
		return
	}
	defer f.Close()
	for {
		readData := ReadSize(f, 64)
		if readData != nil {
			fmt.Println(string(readData))
		} else {
			break
		}
	}
	return
}
