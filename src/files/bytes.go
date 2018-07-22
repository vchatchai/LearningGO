package files

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Bytes() {
	var buffer bytes.Buffer

	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n")

	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stderr)

	buffer.Reset()
	buffer.Write([]byte("Mastering GO"))
	r := bytes.NewReader([]byte(buffer.String()))

	fmt.Println(buffer.String())
	for {
		b := make([]byte, 1)
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}
}
