package main

import (
	"fmt"
	"os"
)

func sort(str string) string {
	result := ""
	var ch rune = 0
	for i := 0; i <= 127; i++ {
		for _, v := range str {
			if v == ch {
				result += string(v)
			}
		}
		ch++
	}
	return result
}

func main() {
	helpStr := "--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument."
	if len(os.Args) > 1 {
		argArr := os.Args[1:]
		lt := 0
		str := ""
		strInsert := ""
		var argRune []rune
		isOrder := false
		isHelp := false
		for _, arg := range argArr {
			argRune = []rune(arg)
			lt = len(argRune)
			if lt >= 2 {
				firstCh := argRune[0]
				secondCh := argRune[1]
				if arg == "--order" || arg == "-o" {
					isOrder = true
				} else if arg == "--help" || arg == "-h" {
					fmt.Println(helpStr)
					isHelp = true
				} else if lt >= 4 && firstCh == '-' && secondCh == 'i' {
					if argRune[2] == '=' {
						strInsert += string(argRune[3:])
					}
				} else if lt >= 10 && firstCh == '-' && secondCh == '-' {
					if string(argRune[2:8]) == "insert" && argRune[8] == '=' {
						strInsert += string(argRune[9:])
					}
				} else {
					str += arg
				}
			} else if lt == 1 && arg != "-" {
				str += arg
			}
		}
		str += strInsert
		if !isHelp {
			if !isOrder {
				fmt.Println(str)
			} else {
				fmt.Println(sort(str))
			}
		}
	} else {
		fmt.Println(helpStr)
	}
}
