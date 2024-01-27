package piscine

// import "fmt"
func Compact(ptr *[]string) int {
	// var result int
	// var arr []string
	var count int
	// l :=len(*ptr)
	for _, v := range *ptr {
		if v != "" {
			count++
		}
	}

	// fmt.Println(*ptr)

	arr := make([]string, count)

	// fmt.Println(arr, len(arr))

	i := 0
	for _, v := range *ptr {
		if v != "" {
			arr[i] = v
			i++
		}
	}

	// fmt.Println(arr, len(arr))

	*ptr = arr

	return len(arr)
}
