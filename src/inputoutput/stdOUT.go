package inputoutput

import (
	"fmt"
	"io"
	"os"
)

func Stdout() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)

	fmt.Fprintf(os.Stdout, "\nTEst %s", myString)

}
