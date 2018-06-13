package inputoutput

import (
	"fmt"
	"os"
	"strconv"
)

func SUM() {
	var sum float64
	for _, arg := range os.Args {
		n, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			sum = n + sum
		}
	}
	fmt.Println(sum)
}

func Average() {
	var total float64
	var count int
	for _, arg := range os.Args {

		n, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			total = n + total
			count++
		}
	}
	result := total / float64(count)
	fmt.Println(result)
}

func StopBreak() {
	for _, arg := range os.Args {
		if arg == "STOP" {
			break
		}
		value, err := strconv.ParseInt(arg, 10, 64)
		if err == nil {
			fmt.Println(value)
		}
	}
}
