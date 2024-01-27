package piscine

func IsUpper(s string) bool {
	count := 0
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
