package piscine

func MakeRange(min, max int) []int {
	var result []int
	if max > min && max != 0 {
		arr := make([]int, max-min)
		counter := 0
		for i := min; i < max; i++ {
			arr[counter] = i
			counter++
		}
		result = arr
	}
	return result
}
