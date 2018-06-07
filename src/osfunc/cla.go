package osfunc

import (
	"strconv"
	"fmt"
	"os"
)

func Cla() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1],64)
	max,_ := strconv.ParseFloat(arguments[1],64)

	for  i := 2 ; i  < len(arguments); i++  {
		n, _ := strconv.ParseFloat(arguments[i],65)
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("MAx:", max)
	
}