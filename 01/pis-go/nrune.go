package piscine

func NRune(s string, n int) rune {
	sRune := []rune(s)
	itemindex := n - 1
	lenStr := len(sRune)

	if itemindex < lenStr && itemindex >= 0 {
		return sRune[itemindex]
	} else {
		return 0
	}
}
