package main

import "fmt"

func main() {
	treasure := "The friends we make along the way."

	// Add your code below:
	fmt.Println(&treasure)

	star := "Polaris"

	// Add your code below:
	starAddress := &star

	fmt.Println("The address of star is", starAddress)
}
