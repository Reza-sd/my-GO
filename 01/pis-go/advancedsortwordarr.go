package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	copy(a, sortSlice(a, f))
}

// --------------------------------------------------
func sortSlice(a []string, f func(a, b string) int) []string {
	var sortedArr []string
	var indexItem int
	middleArr := make([]string, len(a))
	copy(middleArr, a)
	for i := 0; i < len(a); i++ {
		indexItem = indexOfMin(middleArr, f)
		sortedArr = append(sortedArr, middleArr[indexItem])
		middleArr = append(middleArr[:indexItem], middleArr[indexItem+1:]...)
	}
	return sortedArr
}

// ------------------------------------------
func indexOfMin(a []string, f func(a, b string) int) int {
	var index int
	for i := 0; i < len(a); i++ {
		if amIminInSlice(a[i], a, f) {
			index = i
		}
	}
	return index
}

// ---------------------------------------------------
func amIminInSlice(s string, a []string, f func(a, b string) int) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if amILess(s, a[i], f) {
			count++
		}
	}
	if count == len(a) {
		return true
	} else {
		return false
	}
}

// ------------------------------------------------
func amILess(s, v string, f func(a, b string) int) bool {
	if f(s, v) == -1 || f(s, v) == 0 {
		return true
	} else {
		return false
	}
}

//-----------------------------------
