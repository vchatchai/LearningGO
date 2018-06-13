package osfunc

import (
	"fmt"
	"os"
)

func GetENV() {
	for _, env := range os.Environ() {
		fmt.Println(env)
		// fmt.Fprintln("ENV: %s", env)
	}
}

func SetENV(key, value string) (err error) {

	err = os.Setenv(key, value)
	
	return
}
