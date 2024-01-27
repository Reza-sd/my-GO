package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get file arguments
	args := os.Args[1:]

	// Handle no arguments (read from stdin)
	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("ERROR: ", err)
			os.Exit(1)
		}
		fmt.Println(string(data))
		return
	}

	// Handle file arguments
	for _, file := range args {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("ERROR:", err)
			if len(args) == 1 {
				fmt.Println("exit status 1")
			}
			continue
		}
		fmt.Print(string(data))
	}
}
