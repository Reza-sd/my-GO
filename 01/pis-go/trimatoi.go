package piscine

func TrimAtoi(s string) int {
	result := 0
	sByte := []byte(s)
	dec := 1
	sign := 1
	for i := len(sByte) - 1; i >= 0; i-- {
		if sByte[i] >= 48 && sByte[i] <= 57 {
			result = result + ((int(sByte[i] - 48)) * dec)
			dec = dec * 10
		}
	}
	firstNumberPosition := 0
	for index, r := range s {
		if r >= '0' && r <= '9' {
			firstNumberPosition = index
			break
		}
	}
	for i := 0; i < firstNumberPosition; i++ {
		if s[i] == '-' {
			sign = -1
		}
	}
	return sign * result
}
