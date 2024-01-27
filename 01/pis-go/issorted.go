package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending_Count := 0
	descending_Count := 0

	for i := 0; i <= len(a)-2; i++ {
		//-------------
		if f(a[i], a[i+1]) <= 0 {
			ascending_Count++
		}
		//-------------
		if f(a[i], a[i+1]) >= 0 {
			descending_Count++
		}
		//------------
	}
	n := len(a) - 1
	if ascending_Count == n || descending_Count == n {
		return true
	} else {
		return false
	}
}
