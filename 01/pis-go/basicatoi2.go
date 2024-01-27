package piscine

/*
Atoi returns 0 if the string is not considered as a valid number.

For this exercise non-valid string chains will be tested. Some will contain non-digits characters.
*/
func BasicAtoi2(s string) int {
	sArr := []rune(s)

	for _, v := range sArr {
		if v < 48 || v > 57 {
			return 0
		}
	}

	var r int = 0
	dec := 1

	for i := len(sArr) - 1; i >= 0; i-- {
		item := int(sArr[i] - 48)
		r = item*dec + r

		dec = dec * 10
	}

	return r
}
