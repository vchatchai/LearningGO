package stringsutil

import "fmt"

func Rune() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	const r1 = '€'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1:%x\n", sLiteral)
	fmt.Printf("(as a String) r1: %s\n", sLiteral)
	fmt.Printf("(as a character) r1: %c\n", r1)

	fmt.Println("A string is a collection of runes:", []byte("Mihalis"))
	aString := []byte("Mihalis")
	for x, y := range aString {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aString[x])
	}
	fmt.Printf("%d\n", len(sLiteral))
}
