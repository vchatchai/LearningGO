package flags

import (
	"flag"
	"fmt"
)

func SimpleFlag() {
	minusK := flag.Bool("k1", true, "print description")
	minusO := flag.Int("O", 1, "O")

	flag.Parse()
	valueK := *minusK
	valueO := *minusO
	valueO++

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)

}
