package piscine

func AppendRange(min, max int) []int {
	var result []int
	if max > min && max != 0 {
		for i := min; i < max; i++ {
			result = append(result, i)
		}
	}
	return result
}
