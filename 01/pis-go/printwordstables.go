package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(s []string) {
	for _, item := range s {

		for _, v := range item {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
