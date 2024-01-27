package piscine

func IsPrintable(s string) bool {
	sByte := []byte(s)
	for _, r := range sByte {
		if r <= 31 {
			return false
		}
	}
	return true
}
