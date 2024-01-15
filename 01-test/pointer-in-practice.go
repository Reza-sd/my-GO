package main

import "fmt"

func addTen(b *int) {
	*b += 10

}

func main() {

	a := 10
	addTen(&a)
	fmt.Println(a)

}
