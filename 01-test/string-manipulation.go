package main

import "fmt"

func main() {
	str := "Hello"
	//cast them first
	strch := []byte(str)
	strChRune := []rune(str)

	fmt.Println("str=", str)
	fmt.Println("strch(byte)=", strch)
	fmt.Println("strChRune=", strChRune)
	strch[0] = 'A'
	fmt.Println("strch=", strch)
	strFinal := string(strch)
	fmt.Println("strFinal=", strFinal)
}

/*
str= Hello
strch(byte)= [72 101 108 108 111]
strChRune= [72 101 108 108 111]
strch= [65 101 108 108 111]
strFinal= Aello
*/
