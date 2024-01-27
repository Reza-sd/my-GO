package piscine

func StrRev(s string) string {
	sArr := []rune(s)
	var sRev []rune

	for i := len(sArr) - 1; i >= 0; i-- {
		sRev = append(sRev, sArr[i])
	}

	return string(sRev)
}
