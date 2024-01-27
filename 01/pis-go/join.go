package piscine

func Join(strs []string, sep string) string {
	sRune := []rune{}
	sepRune := []rune(sep)
	for i := 0; i < len(strs); i++ {
		for _, v := range strs[i] {
			sRune = append(sRune, v)
		}
		if i < len(strs)-1 {
			for _, value := range sepRune {
				sRune = append(sRune, value)
			}
		}
	}
	return string(sRune)
}
