package main

import "fmt"

func main() {

	array := [5]int{2, 5, 7, 1, 3}
	// This is a slice of the whole array
	sliceVersion := array[:] //<-------------
	fmt.Println(sliceVersion)
	// [2 5 7 1 3]
	// This is a slice containing the elements at indices 2, 3, and 4
	partialSlice := array[2:5]
	fmt.Println(partialSlice)
	// [7 1 3]
	//----------------------------------
	var names = []string{"Kathryn", "Martin", "Sasha", "Steven"}
	fmt.Println(names[1])
	// Martin
	names[3] = "Bishop"
	fmt.Println(names[3])
	// Bishop
	//----------------len cap------------------
	slice := []string{"Fido", "Fifi", "FruFru"}
	// The slice begins at length 3 and capacity 3
	fmt.Println(slice, len(slice), cap(slice))
	// [Fido Fifi FruFru] 3 3
	slice = append(slice, "FroFro")
	// After appending an element when the slice is at capacity
	// The slice will double in capacity, but increase its length by 1
	fmt.Println(slice, len(slice), cap(slice))
	// [Fido Fifi FruFru FroFro] 4 6 (double in capacity)
	//---------------append---------------
	books := []string{"Tom Sawyer", "Of Mice and Men"}
	books = append(books, "Frankenstein")
	books = append(books, "Dracula")
	fmt.Println(books)
	// [Tom Sawyer Of Mice and Men Frankenstein Dracula]
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	fmt.Println(myTutors)
	myTutors = append(myTutors, "Josh")
	fmt.Println(myTutors)
	//---------------------------------
}
