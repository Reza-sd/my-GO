package main

import "fmt"

func main() {
	defer fmt.Println("first output")
	defer fmt.Println("second output")
	fmt.Println("finishing main function")
}

/*
finishing main function
second output
first output
*/
