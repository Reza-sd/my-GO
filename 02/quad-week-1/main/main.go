package main

import (
	"fmt" //<----when test block actgive
	"piscine"
)

func main() {


	testQuad := "A"
	x := 5
	y := 3

	for !(testQuad == "q" || testQuad == "Q") {
		fmt.Println()
	fmt.Println("Which quad do you want to test(A,B,C,D,E)(q = quit) ?")
	fmt.Scan(&testQuad)
	if testQuad=="q" || testQuad=="Q"{break}
	fmt.Println("X?")
	fmt.Scan(&x)
	fmt.Println("Y?")
	fmt.Scan(&y)
	fmt.Printf("\nHere is your share:\n")
	fmt.Println()

	
//--------------------------------------
fmt.Printf("\n-------We are testing-quad-%v--X=%d--Y=%d--------\n\n",testQuad,x,y)


	switch testQuad {
	case "A","a":
		
		piscine.QuadA(x, y)
	case "B","b":
		
		piscine.QuadB(x, y)
	case "C","c":
		
		piscine.QuadC(x, y)
	case "D","d":
		
		piscine.QuadD(x, y)
	case "E","e":
		
		piscine.QuadE(x, y)
	}
//---------------------------------------------	
}


}
