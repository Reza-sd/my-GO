package piscine

// import "fmt"
func Abort(a, b, c, d, e int) int {
	inputArr := []int{a, b, c, d, e}
	// var result int
	// fmt.Println(inputArr)
	// fmt.Println(inputArr[4])
	sortedInput := []int{}
	midleArr := make([]int, len(inputArr))

	copy(midleArr, inputArr) // copy(destination, source)
	// fmt.Println("inputArr= ",inputArr)
	// fmt.Println("midleArr=",midleArr)
	// fmt.Println("sortedInput=",sortedInput)

	var minIndex int
	for i := 1; i <= len(inputArr); i++ {
		minIndex = whoMin(midleArr)

		sortedInput = append(sortedInput, midleArr[minIndex])

		// fmt.Println("sortedInput= ",sortedInput)
		// midleArr = midleArr[:minIndex]+midleArr[minIndex+1:]
		midleArr = append(midleArr[:minIndex], midleArr[minIndex+1:]...)
		// fmt.Println("midleArr= ",midleArr)
	}

	// fmt.Println(sortedInput)
	return sortedInput[2]
}

// --------------
func whoMin(arr []int) int {
	var iamMin int
	var index int

	for _, n := range arr {
		iamMin = n

		//------am I min?------------------
		for i, v := range arr {
			if iamMin > v {
				iamMin = v
				index = i
			}
		}
		//-----------------------------

	}
	return index
}
