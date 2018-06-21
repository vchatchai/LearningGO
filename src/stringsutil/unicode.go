package stringsutil

import (
	"fmt"
	"unicode"
)

func UnitCode() {
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c ", sL[i])
		} else {
			fmt.Print("No printable.")
		}
		fmt.Printf("   %c", sL[i])
		fmt.Printf("   %#U ", sL[i])
		fmt.Printf(", IsNumber  %t", unicode.IsNumber(rune(sL[i])))
		fmt.Printf(", IsLetter  %t", unicode.IsLetter(rune(sL[i])))
		fmt.Println()
	}

	fmt.Printf("ไ IsLetter  %t \n", unicode.IsLetter(rune('ไ')))
}
