package files

import (
	"encoding/binary"
	"fmt"
	"os"
)

func Random() (err error) {
	f, err := os.Open("/dev/random")
	if err != nil {
		return
	}
	defer f.Close()

	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("seed:", seed)
	return
}
