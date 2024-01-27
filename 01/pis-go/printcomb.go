package piscine

import "github.com/01-edu/z01"

/*
Write a function that prints, in ascending order and on a single line:
all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.

These combinations are separated by a comma and a space.
*/
func PrintComb() {
	var ch1 rune = 48 // 97
	var ch2 rune = 48 // 97
	var ch3 rune = 48 // 97
	// 0123456789
	////48-57
	for a := 0; a <= 9; a++ {

		for b := 0; b <= 9; b++ {

			if a < b {

				for c := 0; c <= 9; c++ {

					if b < c && a != c {

						if a == 0 && b == 1 && c == 2 {
						} else {
							z01.PrintRune(44)
							z01.PrintRune(32)

						}

						z01.PrintRune(ch1)
						z01.PrintRune(ch2)
						z01.PrintRune(ch3)
						// z01.PrintRune(44)
						// z01.PrintRune(32)
					}
					ch3++

				}
				ch3 = 48

			}
			ch2++

		}
		ch2 = 48
		ch1++

	}

	// z01.PrintRune('\b')
	// z01.PrintRune('\b')
	// z01.PrintRune('\b')
	z01.PrintRune('\n')
	// c++
}
