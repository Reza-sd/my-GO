package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	sArr := []rune(s)

	for _, v := range sArr {
		z01.PrintRune(v)
	}
}
