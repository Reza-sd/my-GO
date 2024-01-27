package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var nArr []rune
	reminder := 1
	//----------------
	if n == 0 {
		nArr = append(nArr, rune(48))
	} else {
		for {
			reminder = n % 10
			if reminder == 0 && n <= 9 {
				break
			}
			nArr = append(nArr, rune(48+reminder))
			n = n / 10
		}
	}
	//-------------------------
	var result []rune
	var ch rune = 48
	for i := 0; i <= 9; i++ {
		for _, v := range nArr {
			if v == ch {
				result = append(result, ch)
			}
		}
		ch++
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(result[i])
	}
}
