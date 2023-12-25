// Program to return an anonymous function

package main

import "fmt"

// function that returns an anonymous function
func displayNumber() func() int {

	number := 10
	return func() int {
		number++
		return number
	}
}

/*
Here, the func() denotes that displayNumber() returns a function, whereas the int denotes that the anonymous function returns an integer. Notice that displayNumber() is a regular function.
*/
func main() {

	a := displayNumber()

	fmt.Println(a())

}
