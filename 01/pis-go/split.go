package piscine

func Split(str, sep string) []string {
	var result []string
	lsep := len(sep)
	lstr := len(str)
	p := 0
	for i := 0; i+lsep <= lstr; i++ {
		if str[i:i+lsep] == sep {
			result = append(result, str[p:i])
			p = i + lsep
		}
		if i+lsep == lstr {
			result = append(result, str[p:])
		}
	}
	return result
}
