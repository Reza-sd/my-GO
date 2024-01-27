package piscine

func BasicJoin(elems []string) string {
	sRune := []rune{}
	for _, item := range elems {
		for _, v := range item {
			sRune = append(sRune, v)
		}
	}
	return string(sRune)
}
