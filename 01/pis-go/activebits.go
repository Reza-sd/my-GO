package piscine

// import "fmt"
func ActiveBits(n int) int {
	arr := []int{}

	r := 0
	//----------------------
	for {
		r = n % 2
		arr = append(arr, r)
		n = n / 2
		if n == 0 {
			break
		}

	}

	//fmt.Println ("arr=",arr)
	//-----------------------------------

	count := 0

	for _, v := range arr {
		if v == 1 {
			count++
		}
	}

	//--------------------------------

	return count
}
