package piscine

import (
	"sort"
	"strings"
)

func MaxWordCountN(text string, n int) map[string]int {
	seenMap := make(map[string]int)
	wordArr := strings.Split(text, " ")
	for _, word := range wordArr {
		seenMap[word]++
	}
	var numSlice []int
	for key := range seenMap {
		numSlice = append(numSlice, seenMap[key])
	}
	sort.Ints(numSlice)
	singleNumSlice := removeDuplicatesInt(numSlice)
	var r []string
	var finalRes []string
	var x int
	//-----------------------------------
	for i := len(singleNumSlice) - 1; i >= 0; i-- {
		x = singleNumSlice[i]
		for key := range seenMap {
			if seenMap[key] == x {
				r = append(r, key)
			}
		}
		sort.Strings(r)
		finalRes = append(finalRes, r...)
		r = []string{}
	}
	//-------------------------------------------
	finalMap := make(map[string]int)

	for i := 0; i < n; i++ {
		finalMap[finalRes[i]] = seenMap[finalRes[i]]
	}
	//--------------------------------
	return finalMap
}

// -------------------------------------------
func removeDuplicatesInt(input []int) []int {
	unique := make(map[int]bool)
	result := []int{}
	for _, value := range input {
		if !unique[value] {
			unique[value] = true
			result = append(result, value)
		}
	}
	return result
}

//------------------------------
