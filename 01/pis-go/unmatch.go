package piscine

func Unmatch(nums []int) int {
	seen := make(map[int]int)
	for _, num := range nums {
		seen[num]++
	}
	for i := 0; i < len(nums); i++ {
		if seen[nums[i]]%2 != 0 {
			return nums[i]
		}
	}
	return -1
}
