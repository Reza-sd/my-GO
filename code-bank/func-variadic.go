//Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

package main

import "fmt"

// This variadic function takes an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print("The sum of ", nums) // Also a variadic function.
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(" is", total) // Also a variadic function.
}

//-------------------------------------
/*
for _, num := range nums {:

This line initiates a for loop specifically designed to iterate over elements within a collection like a slice or array.
It uses the range keyword to access each element and its index within nums.
The _ indicates that we're not interested in using the index of each element, so it's discarded.
The num := range nums part assigns each element to the variable num during each iteration.
*/

// ----------------------------------
func main() {
	// Variadic functions can be called in the usual way with individual
	// arguments.
	sum(1, 2)    // "The sum of [1 2] is 3"
	sum(1, 2, 3) // "The sum of [1 2 3] is 6"

	// If you already have multiple args in a slice, apply them to a variadic
	// function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...) // "The sum of [1 2 3 4] is 10"
}
