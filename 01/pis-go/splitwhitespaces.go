package piscine

func SplitWhiteSpaces(s string) []string {
	var resultArr []string
	word := ""
	ls := len(s)
	for index, r := range s {
		//-----------------------
		if r == ' ' || r == 9 || r == '\n' {
			// fmt.Println(index,word)
			resultArr = append(resultArr, word)
			word = ""
			continue
		}
		//------------------------------
		word = word + string(r)

		//---------last letter----------
		if index == ls-1 {
			resultArr = append(resultArr, word)
			break
		}
		//--------------------
	}
	//-------------remove ""----------------
	var lastArr []string
	for _, v := range resultArr {
		if v != "" {
			lastArr = append(lastArr, v)
		}
	}
	//---------------------------
	return lastArr
}
