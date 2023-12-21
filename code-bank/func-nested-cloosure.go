//closure is a nested function that helps us access the outer function's variables even after the outer function is closed.

package main

import "fmt"

// outer function
func greet() func() string {

	// variable defined outside the inner function
	name := "John"

	// return a nested anonymous function
	return func() string {
		name = "Hi " + name
		return name
	}
}

func main() {

	// call the outer function
	message := greet()

	// call the inner function
	fmt.Println(message())

}
