package main

import (
	//"fmt"
	"github.com/01-edu/z01"
)

// --------struct--------------
type point struct {
	x int
	y int
}

//----------setPoint----------
//The function setPoint() must work with int.

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// -----------------------
func toRune(num int) []rune {
	r := []rune{}
	var remainder int
	for num > 0 {
		remainder = num % 10
		num = num / 10
		r = append(r, rune(remainder+48))

	}
	lenR := len(r)
	// fmt.Println(lenR)
	rev := make([]rune, lenR)
	for i, v := range r {
		rev[lenR-i-1] = v
	}
	return rev
}

// ----------main---------------------
func main() {
	points := &point{}
	setPoint(points)
	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
	var msgRune []rune
	x := string(toRune(points.x))
	y := string(toRune(points.y))
	msg := "x = " + x + ", y = " + y
	msgRune = []rune(msg)
	for _, v := range msgRune {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
