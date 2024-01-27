package piscine

func IterativeFactorial(nb int) int {
	// stop condition
	if nb == 0 || nb == 1 {
		return 1
	}
	// Errors (non possible values or overflows) will return 0.
	if nb < 0 || nb > 25 {
		return 0
	}

	return nb * IterativeFactorial(nb-1)
}
