package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argArr := os.Args[1:]
	var ch byte = 0
	for i := 0; i <= 127; i++ {
		for _, item := range argArr {
			if []byte(item)[0] == ch {
				//-----print--------------
				for _, char := range item {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')
				//-----------------------
			}
		}
		ch++
	}
}
