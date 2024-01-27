package piscine

import "github.com/01-edu/z01"

/*
Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.

These combinations are separated by a comma and a space.
*/

func PrintComb2() {
	var ch1 rune = 48 // 97
	var ch2 rune = 48 // 97
	var ch3 rune = 48
	var ch4 rune = 48

	// 0123456789

	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {

					if a <= c {
						if !(a == b && b == c && c == d) {
							if a == c && b < d || a < c {

								if a == 0 && b == 0 && c == 0 && d == 1 {
								} else {
									z01.PrintRune(32)
								}

								z01.PrintRune(ch1)
								z01.PrintRune(ch2)

								z01.PrintRune(32)

								z01.PrintRune(ch3)
								z01.PrintRune(ch4)

								// z01.PrintRune(44)
								if a == 9 && b == 8 && c == 9 && d == 9 {
								} else {
									z01.PrintRune(44)
								}

							}
						}
					}
					ch4++
				}
				ch4 = 48
				ch3++
			}
			ch3 = 48
			ch2++

		}
		ch2 = 48
		ch1++
	}

	z01.PrintRune('\n')
}
