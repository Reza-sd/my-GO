package piscine

func Max(a []int) int {
	return SortSliceMax(a)[len(a)-1]
}

// --------------------------------------------------
func SortSliceMax(a []int) []int {
	var sortedArr []int
	var indexItem int
	middleArr := make([]int, len(a))
	copy(middleArr, a)
	for i := 0; i < len(a); i++ {
		indexItem = IndexOfMinMax(middleArr)
		sortedArr = append(sortedArr, middleArr[indexItem])
		middleArr = append(middleArr[:indexItem], middleArr[indexItem+1:]...)
	}
	return sortedArr
}

// ------------------------------------------
func IndexOfMinMax(a []int) int {
	var index int
	for i := 0; i < len(a); i++ {
		if AmIminInSliceMax(a[i], a) {
			index = i
		}
	}
	return index
}

// ---------------------------------------------------
func AmIminInSliceMax(s int, a []int) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if AmILessMax(s, a[i]) {
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
func AmILessMax(s, v int) bool {
	if s <= v {
		return true
	} else {
		return false
	}
}

//-----------------------------------
