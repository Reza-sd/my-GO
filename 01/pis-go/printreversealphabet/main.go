package main

import "github.com/01-edu/z01"

func main() {
	var c rune = 122 // 97
	// abcdefghijklmnopqrstuvwxyz

	for i := 0; i <= 25; i++ {

		z01.PrintRune(c)
		c--
	}

	z01.PrintRune('\n')
}
