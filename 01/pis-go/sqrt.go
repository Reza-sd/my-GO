package piscine

func Sqrt(nb int) int {
	result := 0

	if nb == 1 {
		return 1
	}

	for i := 1; i < nb; i++ {
		if nb == i*i {
			return i
		}
	}

	return result
}
