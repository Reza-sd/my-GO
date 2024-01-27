package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		argArr := os.Args[1:]
		firtArg := argArr[0]
		isUpper := false
		alphArr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25"}
		loArr := []rune("abcdefghijklmnopqrstuvwxyz")
		upArr := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		if firtArg == "--upper" {
			isUpper = true
			argArr = argArr[1:]
		}
		isMatch := false
		for _, item := range argArr {
			for index, v := range alphArr {
				if item == v {
					if isUpper {
						z01.PrintRune(upArr[index])
					} else {
						z01.PrintRune(loArr[index])
					}
					isMatch = true
				}
			}
			if !isMatch {
				z01.PrintRune(' ')
			}
			isMatch = false
		}
		z01.PrintRune('\n')
	}
}
