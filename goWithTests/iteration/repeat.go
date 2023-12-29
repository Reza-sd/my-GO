package iteration

func Repeat(character string) string {
	result := ""

	for i := 0; i < 5; i++ {
		result += character
	}

	return result
}
