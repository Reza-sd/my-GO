package piscine

// import "github.com/01-edu/z01"
func AlphaCount(s string) int {
	var count int

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}

	// z01.PrintRune('\n')
	return count
}
