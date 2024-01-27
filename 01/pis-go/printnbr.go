package piscine

import (
	//"fmt"
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var ch rune = 48
	sn := n
	if sn == -9223372036854775808 {
		n = 9223372036854775800
		z01.PrintRune(45)
	}

	if n < 0 {
		z01.PrintRune(45)

		n = n * -1

		// fmt.Println("--n>", n)
	}
	if n == 0 {
		z01.PrintRune(48)
	}
	var ra int
	var arr []int

	for n > 0 {

		ra = n % 10
		n = n / 10
		arr = append(arr, ra)

	}

	if sn == -9223372036854775808 {
		arr[0] = 8
	}

	for i := len(arr); i > 0; i-- {

		for x := 48; x < (48 + arr[i-1]); x++ {
			ch++
		}
		z01.PrintRune(ch)
		ch = 48
	}
}
