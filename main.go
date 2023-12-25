package main

import "fmt"

// ======================================
type Pet struct {
	name    string
	petType string
	age     int
}

// --------------------------------
type Country struct {
	name      string
	capital   string
	latitude  float32
	longitude float32
}

// --------------------------------
type Student struct {
	firstName string
	lastName  string
	age       int
	grade     int
}

// --------------------------------
type Restaurant struct {
	name             string
	typeOfRestaurant string
	yearEstablished  int
}

// ----------------------------------
type Triangle struct {
	height float32
	base   float32
	//area()
	//updateBase ()
}

func (triIns Triangle) area() float32 {
	return triIns.height * triIns.base / 2
}

//---------------pointer to struct-----------

func (triIns *Triangle) updateBase(newBase float32) {
	triIns.base = newBase
}

// -----------------Arrays of Structs-----------------------
type Cake struct {
	typeOfCake string
	weight     int // Weight in grams
}

// ------------------Nested Structs-------------
type Name struct {
	firstName string
	lastName  string
}

type Employee struct {
	name  Name
	age   int
	title string
}

//-----------------anonymous field struct------------

type Point struct {
	x int
	y int
}

// Checkpoint 1 goes here
type Circle struct {
	Point  //<-----------------anonymous
	radius int
}

//-----------------------------------------------------

// ======================================
func main() {
	//---------------------------------
	nuggets := Pet{"Nuggets", "dog", 4}
	mittens := Pet{"Mittens", "cat", 7}
	robin := Pet{"Robin", "bird", 2}
	fmt.Println(nuggets)
	fmt.Println(mittens)
	fmt.Println(robin)
	//---------------------------------
	var france Country
	fmt.Println(france)
	//--------------------------------
	peter := Student{"Peter", "Bookman", 16, 11}
	fmt.Println(peter)
	scott := Student{firstName: "Scott", lastName: "Peterson", grade: 12}
	fmt.Println(scott)
	//--------------------------------
	restaurant := Restaurant{"Codecademy Steakhouse", "Japanese", 2011}
	fmt.Println(restaurant)
	fmt.Println(restaurant.name)
	fmt.Println(restaurant.typeOfRestaurant)
	fmt.Println(restaurant.yearEstablished)
	restaurant.name = "Skillsoft Steakhouse"
	restaurant.yearEstablished = 2022
	fmt.Println(restaurant)
	//---------------------------------
	//triangle
	myTri := Triangle{10, 4}
	fmt.Println(myTri)

	// Call the function here
	fmt.Println(myTri.area())

	//-----------Pointers to a Struct---------------

	triangle := Triangle{10, 4}
	fmt.Println(triangle)
	// Call the function here
	triangle.updateBase(13)
	fmt.Println(triangle)
	//-----------------Arrays of Structs-----------------------

	// Checkpoint 1 code goes here
	cakes := []Cake{{"Chocolate", 1000}, {"Carrot", 500}, {"Ice Cream", 2000}} // Checkpoint 2 code goes here

	fmt.Println(cakes)
	// Checkpoint 2 code goes here
	fmt.Println(cakes[0].weight)
	// Checkpoint 3 code goes here
	cakes[1].weight = 1500
	fmt.Println(cakes)
	//------------------Nested Structs-------------

	carl := Employee{Name{"Carl", "Carlson"}, 32, "Engineer"}
	fmt.Println(carl)
	fmt.Println(carl.name.lastName) // Output will be “Carlson”
	//-----------------anonymous field struct------------

	circle := Circle{Point{4, 5}, 2}

	// Checkpoint 2 code goes here
	fmt.Println(circle.x)
	fmt.Println(circle)
	//---------------------------------
} //END main

//======================================
