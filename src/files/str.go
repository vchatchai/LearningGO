package files

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func StrRead() {
	r := strings.NewReader("Test")

	fmt.Println("r length:", r.Len())

	b := make([]byte, 1)
	for {
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

func StrWrite() {
	w := strings.NewReader("This is an error\n")
	fmt.Println("r length:", w.Len())
	n, err := w.WriteTo(os.Stderr)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wrote %d bytes to os.Stderr\n", n)

}
