package piscine

func Itoa(n int) string {
	var str string
	if n == 0 {
		return "0"
	}
	var isNeg bool
	if n < 0 {
		n = -1 * n
		isNeg = true
	}
	r := n
	var ch rune
	ch = '0'
	var arrRune []rune
	for {

		r = n % 10
		n = n / 10

		for i := 0; i < r; i++ {
			ch++
		}

		arrRune = append(arrRune, ch)
		ch = '0'
		if n == 0 {
			break
		}

	}
	arrRuneRev := []rune{}
	for i := len(arrRune) - 1; i >= 0; i-- {
		arrRuneRev = append(arrRuneRev, arrRune[i])
	}

	str = string(arrRuneRev)

	if isNeg {
		str = "-" + str
	}
	return str
}
