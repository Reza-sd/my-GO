package piscine

// import "fmt"
func Rot14(s string) string {
	str := ""
	var chRune rune
	var chString string
	//fmt.Println("s=", s)
	//-----------------------------
	for _, v := range s {
		// 26 letters
		//----------------------------------
		if v >= 'a' && v <= 'z' {
			if 'z'-v >= 14 {
				chRune = v + 14
			} else {
				chRune = 'a' + 14 - ('z' - v) - 1
			}
			//-----------------------
		} else if v >= 'A' && v <= 'Z' {
			if 'Z'-v >= 14 {
				chRune = v + 14
			} else {
				chRune = 'A' + 14 - ('Z' - v) - 1
			}
		} else {
			chRune = v
		}
		//--------------------------------------
		chString = string(chRune)
		// fmt.Println("chString=", chString)
		str += chString

	}

	return str
}
