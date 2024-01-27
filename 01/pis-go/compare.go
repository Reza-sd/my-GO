package piscine

func Compare(a, b string) int {
	result := 0

	if a < b {
		result = -1
	} else if a > b {
		result = 1
	} else {
		result = 0
	}

	return result
}
