package piscine

// import "fmt"
func CollatzCountdown(start int) int {
	// var result int
	if start <= 0 {
		return -1
	}

	// fmt.Println(isEven(start))
	count := 0
	n := start
	for {
		if n == 1 {
			break
		} else if isEven(n) {
			n = n / 2
		} else if !isEven(n) {
			n = (3 * n) + 1
		}
		count++
	}

	return count
}

// ---------------------
func isEven(n int) bool {
	return n%2 == 0
}
