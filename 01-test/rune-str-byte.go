package main

import "fmt"

func main() {

	str := "你好"
	strRune := []rune(str) //[20320 22909]
	strByte := []byte(str) //[228 189 160 229 165 189]

	fmt.Println(strRune)
	fmt.Println(strByte)
	//------------------
	for index, v := range str {

		fmt.Printf("str : index: %v, value: %d => letter = %c \n", index, v, v)
	}
	//--------------
	for index, v := range strRune {

		fmt.Printf("strRune : index: %v, value: %d => letter = %c \n", index, v, v)
	}
	//-------------------------
	for index, v := range strByte {

		fmt.Printf("strByte : index: %v, value: %d => letter = %c \n", index, v, v)

	}
	/*
	[20320 22909]
	[228 189 160 229 165 189]
	str : index: 0, value: 20320 => letter = 你
	str : index: 3, value: 22909 => letter = 好
	strRune : index: 0, value: 20320 => letter = 你
	strRune : index: 1, value: 22909 => letter = 好
	strByte : index: 0, value: 228 => letter = ä
	strByte : index: 1, value: 189 => letter = ½
	strByte : index: 2, value: 160 => letter =
	strByte : index: 3, value: 229 => letter = å
	strByte : index: 4, value: 165 => letter = ¥
	strByte : index: 5, value: 189 => letter = ½
	*/
}
