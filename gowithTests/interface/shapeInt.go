package main

import (
	"fmt"
	"math"
)

// ---------------
// interface group types together based on the methods
type shape interface {
	Area() (float64, string)
	Perimeter() float64
}

// -----------------Rectangle-------------------------
type Rectangle struct {
	Width  float64
	Height float64
}

// -----
// define method func Area()
func (r Rectangle) Area() (float64, string) {
	return r.Width * r.Height, "Rectangle"
}

// ---
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
//-------------------Circle----------------------
type Circle struct {
	Radius float64
}
// -----
// define method func Area()
func (c Circle) Area() (float64, string) {
	return math.Pi * c.Radius * c.Radius, "Circle"
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
// -----------------------
func main() {
	myShapes := []shape{
		Rectangle{Width: 10.5, Height: 5.1},
		Circle{Radius: 10.5},
		Rectangle{12.0, 6.0},
		Rectangle{5, 2.5},
		Rectangle{Width: 8, Height: 3},
		Circle{10},
		Circle{Radius: 4},
	}
	// fmt.Println(myShapes[0].Area())
	// fmt.Println(myShapes[1].Area())
	for _, shape := range myShapes {
		Area, Type := shape.Area() //Multiple Return Values
		Peri := shape.Perimeter()
		fmt.Println("\n Area of", Type, shape, "=", Area, ",and Perimeter =", Peri)
	}

}
