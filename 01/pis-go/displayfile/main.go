package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	ArgArr := os.Args[1:]
	lenArg := len(ArgArr)
	var isArgExist bool
	var isMoreThanOne bool
	var finalMsg string
	//---------------------------------
	if lenArg == 0 {
		isArgExist = false
	} else {
		isArgExist = true
	}
	//----------------------------------
	if lenArg > 1 {
		isMoreThanOne = true
	}
	//--------------------------
	if isMoreThanOne {
		finalMsg = "Too many arguments"
	}
	//-------------------------------
	if isArgExist == false {
		finalMsg = "File name missing"
	}
	//-------------------------------------------
	if lenArg == 1 {
		arg := ArgArr[0]
		// Read the entire contents of a file into a byte slice
		content, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// Print the content as a string
		// fmt.Println("File content:")
		// fmt.Println(string(content))
		finalMsg = string(content[:len(content)-1])
	}
	//----------------------------
	fmt.Println(finalMsg)
}
