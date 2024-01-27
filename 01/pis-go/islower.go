package piscine

func IsLower(s string) bool {
	count := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
