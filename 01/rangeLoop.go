package main

import "fmt"

func main() {

	mySlice := []string{
		"Word1",
		"Word2",
		"Word3",
		"Word4",
		"Word5",
		"Word6",
	}
	for index, word := range mySlice {
		fmt.Printf("The index is: %v the element is: %v\n", index, word)
	}
	for _ , word := range mySlice {
		fmt.Printf("the element is: %v\n", word)
	}

}
