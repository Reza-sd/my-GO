package piscine

func RecursivePower(nb int, power int) int {
	// result := 1

	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}

	return nb * RecursivePower(nb, power-1)
}
