package piscine

func Index(s string, toFind string) int {
	ls := len(s)
	lf := len(toFind)
	//--------------------------------
	if toFind == "" || s == toFind {
		return 0
	}
	//------------------------------
	if s != "" && toFind != "" && ls > lf {
		for index, value := range s {
			if value == []rune(toFind)[0] {
				if ls >= index+lf {
					if s[index:index+lf] == toFind {
						return index
					}
				}
			}
		}
	}
	//--------------------------
	return -1
}
