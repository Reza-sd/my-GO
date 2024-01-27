package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, item := range os.Args {
		if i == 0 {
			continue
		}
		for _, char := range item {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
