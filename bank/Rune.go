package main

import "fmt"

func main() {
	str := "Hello"
	str = str[1:]
	fmt.Println(str)
	rs := []rune(str)
	fmt.Println(rs)

	fmt.Println(string(10084)) // utf-8 heart emoji

	for _, v := range str {

		fmt.Println(string(v))
		fmt.Printf("---%+q---", v)
		fmt.Println(v)
		fmt.Println(v == 101)
		fmt.Println(v == 'l')

	}
	fmt.Println("\n", int('a'))
	fmt.Println(-int(2 + '1'))

}
