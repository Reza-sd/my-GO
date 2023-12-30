package iteration

func Repeat(character string, iteration int) string {
	result := ""
	for i := 0; i < iteration; i++ {
		result += character
	}
	return result
}
