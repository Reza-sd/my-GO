package piscine

func BasicAtoi(s string) int {
	sArr := []rune(s)
	var r int = 0
	dec := 1

	for i := len(sArr) - 1; i >= 0; i-- {
		item := int(sArr[i] - 48)
		r = item*dec + r

		dec = dec * 10
	}

	return r
}
