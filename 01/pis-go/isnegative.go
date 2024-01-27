/*
Write a function that prints 'T' (true) on a single line if the int passed as parameter is negative, otherwise it prints 'F' (false).

*/

package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb >= 0 {
		z01.PrintRune(70) // F = 70
	} else {
		z01.PrintRune(84) // F = 84
	}

	z01.PrintRune('\n')
}
