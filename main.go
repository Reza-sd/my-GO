package main

import (
	"fmt"
)

func main() {

	var currencies = map[string]float32{
		"JPY": 130.2,
		"EUR": 0.95,
	}

	var dollarAmount float32
	var currency string

	fmt.Println(currencies, currency)

	fmt.Println("Please enter dollar Amount:")
	fmt.Scan(&dollarAmount)

	if dollarAmount == 0 {
		fmt.Println("Invalid dollarAmount.")
	} else {
		// rest of your code
		fmt.Println("Please enter your  the target currency:")
		fmt.Scan(&currency)

		rate, isInMap := currencies[currency]
		if isInMap {
			fmt.Println("we found the rate", rate)
			fmt.Println("The amount of target currency is : ", rate*dollarAmount)

		} else {
			//zero value if key is not in the map)
			fmt.Println("no such rate!", rate)
		}

	}

}
