package piscine

// ---------------------------
func ToLower2(s string) string {
	sByte := []byte(s)
	for index, value := range sByte {
		if value >= 65 && value <= 90 {
			sByte[index] += 32
		}
	}
	return string(sByte)
}

// ------------------------------------
func ToUpper2(s string) string {
	sByte := []byte(s)
	for index, value := range sByte {
		if value >= 97 && value <= 122 {
			sByte[index] -= 32
		}
	}
	return string(sByte)
}

// -------------------------------
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	s = ToLower2(s)
	sByte := []byte(s)
	sByte[0] = []byte(ToUpper2(string(sByte[0])))[0]
	for index, v := range sByte {
		if !((v >= 65 && v <= 90) || (v >= 97 && v <= 122) || (v >= 48 && v <= 57)) && index < len(sByte)-1 {
			sByte[index+1] = []byte(ToUpper2(string(sByte[index+1])))[0]
		}
	}
	return string(sByte)
}
