package stringsutil

import "fmt"

const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"

func PrintStrings() {
	s2 := "€£³"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)
	fmt.Printf("sLiteral length: %d\n", len(sLiteral))
	fmt.Println(s2)
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}
}
