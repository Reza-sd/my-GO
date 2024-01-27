package main

import (
	"os"

	"github.com/01-edu/z01"
)

// ---------------------------
//  takes int as argument and "it passed the argumet to even func"return true if it is negative
// it seems even function return a int value , is seems 1 one the value is 1

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

// ---------------------
func even(num int) int {
	var result int

	//---sucodoed : Dear my 2yrs son :  an even num is : is num when / 2 the remainder is 0
	var remainder int
	remainder = num % 2

	if remainder == 0 {
		// this number is even
		result = 1
	}

	return result
}

// ------------------------------
// -----------printStr--------------------
// takes  string as argument and print its characters followed by a new line
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// ----------------------------
func main() {
	arg := os.Args
	var isNumofArgEven bool
	numOfArg := len(arg) - 1
	msgEven := "I have an even number of arguments"
	msgOdd := "I have an odd number of arguments"
	var whatIPrint string
	isNumofArgEven = isEven(numOfArg)
	//----------------
	if isNumofArgEven == true {
		whatIPrint = msgEven
	} else {
		whatIPrint = msgOdd
	}
	printStr(whatIPrint)
	//----------------------------
}
