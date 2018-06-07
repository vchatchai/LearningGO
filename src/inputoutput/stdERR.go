package inputoutput

import (
	"io"
	"os"
)

func Stderr(){
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me on argument!"
	}  else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is standard output")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}