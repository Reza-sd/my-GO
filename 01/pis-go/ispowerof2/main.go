package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// --------------------------------------
func main() {
	if !(len(os.Args) == 2) {
		return
	}

	str := os.Args[1]
	//----------------------------------
	num, err := strconv.Atoi(str)
	// Check for errors
	if err != nil {
		// fmt.Println("Error:", err)
		return
	}
	//---------------------------------------
	if num <= 0 {
		return
	}

	msg := "false"

	if isPowerTwo(num) {
		msg = "true"
	}

	for _, v := range msg {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

// ------------------------------------------
func isPowerTwo(n int) bool {
	var r int

	for {

		if n == 1 && r == 0 {
			return true
		} else if n == 1 && r > 0 {
			return false
		}
		r += n % 2
		n = n / 2

	}
}
