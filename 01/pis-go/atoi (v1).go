package piscine

func Atoi1(s string) int { //<----rename Atoi
	//-------------------------------
	//+43 -45 <---

	sArr := []rune(s)

	firstCharacter := sArr[0]
	neg := +1
	//--------------------------
	if firstCharacter == 45 {
		neg = -1
	}

	if firstCharacter == 43 || firstCharacter == 45 {
		sArr = sArr[1:]
	}

	//---------------------------
	for _, v := range sArr {
		if v < 48 || v > 57 {
			return 0
		}
	}

	//-------------------------------
	var r int = 0
	dec := 1

	for i := len(sArr) - 1; i >= 0; i-- {
		item := int(sArr[i] - 48)
		r = item*dec + r

		dec = dec * 10
	}
	//-------------------------------

	return r * neg
}
