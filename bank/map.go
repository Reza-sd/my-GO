package main

import "fmt"

func main() {
	// ---------------------------------------------
	// Create a simple contact list.
	contacts := map[string]int{
		"Joe":    2126778723,
		"Angela": 4089978763,
		"Shawn":  3143776876,
		"Terell": 5026754531,
	}
	// Print out all the contacts
	fmt.Println(contacts)
	//--------------------true/false----------------------------
	//customer,status := customers["billy"]
	// To test for a key without retrieving the value, use an underscore in place of the first value:

	_, status := contacts["Joe"]
	fmt.Println(status)

	if status {
		fmt.Println("we found the contact")
	} else {
		fmt.Println("no such contact!")
	}

	//---------------Declaring an Empty Map----------------
	// Create a empty map called emptyMap
	emptyMap := make(map[string]int)

	// Print map
	fmt.Println(emptyMap)
	//---------------Declaring a Map with Values-------------
	// Initialize map gradebook with values
	gradebook := map[string]float32{"John": 85.2, "Ana": 95.6}

	// Print map gradebook
	fmt.Println(gradebook)
	//----------------Accessing Elements-----------------------
	// Print the value with key "John"
	fmt.Println(gradebook["John"])
	//--------Iterating All Key/Value Pairs in a Map----------
	// Iterate all key/value pairs of the gradebook map
	for key, value := range gradebook {
		fmt.Printf("(%s, %.1f)\n", key, value)
	}
	//----------Storing a Map Value in a Variable--------------
	// Store the value that has a key of "Ana" in anaScore
	anaScore := gradebook["Ana"]

	fmt.Println(anaScore)

	//----------Accessing a Key That Doesn’t Exist-------------
	// Store the value that has a key of "John" in johnScore
	johnScore := gradebook["David"]

	// Since "David" does not exist in the map, 0 will be printed
	fmt.Println(johnScore)
	//--------------Adding Values---------------------------
	// Print the initialized map
	fmt.Println(gradebook)

	// Add more key-value pairs
	gradebook["George"] = 76.4
	gradebook["Emma"] = 90

	// Print the map again
	fmt.Println(gradebook)
	//------------Removing Values-------------------------------
	// Print the initialized map
	fmt.Println(gradebook)

	// Delete an item
	delete(gradebook, "John")

	// Print the map again
	fmt.Println(gradebook)
	//----------------------len(map)-----------------------
	var bakingSupplies = map[string]int {
		"sugar": 4 ,
		"apples": 4,
		"flour": 4,
		"eggs":3,
	}

	fmt.Println(len(bakingSupplies))
	//---------------------------------
}
