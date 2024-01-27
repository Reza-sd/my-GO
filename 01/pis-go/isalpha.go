package piscine

func IsAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r == ' ') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}
