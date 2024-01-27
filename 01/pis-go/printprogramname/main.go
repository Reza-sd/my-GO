package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var programName string = os.Args[0]
	for index, char := range programName {
		if index <= 1 {
			continue
		}
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
	// name := os.Args[0]
	// for i := 2; i < len(name); i++ {
	// 	z01.PrintRune(rune(name[i]))
	// }
	/*
		progName := os.Args[0]
		pRune := []rune(progName)
		lenP := len(progName)
		indexSlash := 0
		for i, v := range progName {
			if v == '/' {
				indexSlash = i
			}
		}
		for i := indexSlash + 1; i < lenP; i++ {
			z01.PrintRune(pRune[i])
		}
	*/
}
