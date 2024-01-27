package main

import (
	//"fmt"
	"os"
	"strconv"
)

// strconv.Atoi, os.*
func main() {
	//toString(-210)
	//-----------------------------
	if len(os.Args) != 4 {
		return
	}
	// fmt.Println(os.Args)
	// fmt.Println(os.Args[1:])
	aStr := os.Args[1]
	opStr := os.Args[2]
	bStr := os.Args[3]

	//--------------------
	a, errA := strconv.Atoi(aStr)
	if errA != nil {
		return
	}
	//--------------------
	b, errB := strconv.Atoi(bStr)
	if errB != nil {
		return
	}
	//----------------
	//fmt.Println(a,opStr,b)
	//	fmt.Println(opStr)
	//fmt.Println(opStr=="+")
	//----------------------
	if !(opStr == "+" || opStr == "-" || opStr == "/" || opStr == "main.go" || opStr == "*" || opStr == "%") {
		return
	}
	//---------------------
	//fmt.Println(a,opStr,b)
	//-----------------------
	var r int

	switch opStr {
	case "+":
		if a >= 9223372036854775807 && b > 0 {
			return
		}
		r = a + b
	case "-":
		if a <= -9223372036854775807 && b > 0 {
			return
		}
		r = a - b
	case "main.go", "*":
		if a >= 9223372036854775807 || b >= 9223372036854775807 {
			return
		}
		r = a * b
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")

			return
		}
		r = a / b
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		r = a % b
	}

	//-------------------------------
	//fmt.Println(r)
	os.Stdout.WriteString(toString(r) + "\n")
}

// ----------------------------------------------
func toString(n int) string {
	// fmt.Println("n=", n)

	var str string

	if n == 0 {
		return "0"
	}

	var isNeg bool

	if n < 0 {
		n = -1 * n
		isNeg = true
	}

	r := n
	var ch rune
	ch = '0'
	// fmt.Println("ch=", ch)
	var arrRune []rune

	// fmt.Println(string(ch + 1))
	for {

		r = n % 10
		n = n / 10
		// fmt.Println("r=", r)
		// fmt.Println("n=", n)

		for i := 0; i < r; i++ {
			ch++
		}

		// fmt.Println("ch=", ch)
		arrRune = append(arrRune, ch)
		ch = '0'

		if n == 0 {
			break
		}

	}
	//-------------------------
	//fmt.Println(string(arrRune))

	arrRuneRev := []rune{}

	for i := len(arrRune) - 1; i >= 0; i-- {
		arrRuneRev = append(arrRuneRev, arrRune[i])
	}
	//---------------------------
	//fmt.Println(string(arrRuneRev))

	str = string(arrRuneRev)
	// fmt.Println(str,isNeg)
	if isNeg {
		str = "-" + str
	}
	// fmt.Println(str,isNeg)
	return str
}

////------------
