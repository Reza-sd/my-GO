package piscine

func ToUpper(s string) string {
	sByte := []byte(s)
	for index, value := range sByte {
		if value >= 97 && value <= 122 {
			sByte[index] -= 32
		}
	}
	return string(sByte)
}
