package main

import "fmt"

func main() {

	a := 10 //int
	fmt.Println(a)

	// Use %p to print the pointer's value in hexadecimal
	fmt.Printf("address of a is %p \n", &a)

	var myPointers *int = &a

	fmt.Println(myPointers)

}

/*
10
address of a is 0xc0000b8000
-----------------------------
10
address of a is 0xc000120000
0xc000120000
----------------------------
*/
