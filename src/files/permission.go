package files

import (
	"fmt"
	"os"
)

func Permission() {
	fileInfo, err := os.Stat("permission.go")
	if err != nil {
		panic(err)
	}

	fmt.Println(fileInfo.Mode().String())
}
