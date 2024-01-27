package piscine

func ToLower(s string) string {
	sByte := []byte(s)
	for index, value := range sByte {
		if value >= 65 && value <= 90 {
			sByte[index] += 32
		}
	}
	return string(sByte)
}
