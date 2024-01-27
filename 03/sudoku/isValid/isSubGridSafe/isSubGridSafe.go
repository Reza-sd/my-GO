package main

import (
	"fmt"
)

func IsSubGridSafe(board [9][9]int, row, col, nb int) bool {
	// var IsSubgridVal bool // to set the value of subgrid depending on the value on each row and col
	// IsSubgridVal = true   // set default value to 0
	var isNumValid bool // check for validity of nb
	var result bool     // to store the final result
	var isColValid bool // store the value of validity of col
	// fmt.Println("--111111111------value of col =", col)
	var isRowValid bool
	var isPositionEmpty bool // true = the position can be used / false = position can't be used.

	//--------------Check validity of col--------------------
	if col < 1 || col > 9 {
		isColValid = false
	} else {
		isColValid = true
		// fmt.Println("----2222222---iSColValid =--- col =--", isColValid, col)
	}

	//-----------check validity of rows--------------
	if row < 1 || row > 9 {
		isRowValid = false
	} else {
		isRowValid = true
	}

	//-----checkPosition---------------------

	if isColValid && isRowValid {

		if board[row-1][col-1] != 0 { //
			isPositionEmpty = false
			//fmt.Println("-----9999----- isPositionEmpty = false", isPositionEmpty)
		} else {
			isPositionEmpty = true
			// fmt.Println("-----&&&&&&&&----- isPositionEmpty = True", isPositionEmpty)
		}

	}
	

	//-------set up result based on isPositioEmpty-----------
	if isPositionEmpty {
		result = true
		fmt.Println("------888888--->result = false", result)
	} else {
		result = false
		//fmt.Println("------888888--->result = false", result)

	}

	//------set up result based on isRowValid-----------
	if isRowValid {
		result = true
		fmt.Println("------xxxxxxxxxxxxxxxx--->result = true", result)
	} else {
		result = false
		//fmt.Println("------666666--->result = false", result)
	}
	//------set up result based on isColValid-----------

	if isColValid && isRowValid {
		result = true
		fmt.Println("------33333333333----->result = true", result)
	} else {
		result = false
		//fmt.Println("------33333333333----->result = false", result)
	}
	//-------------check for validy -----------
	if nb >= 1 && nb <= 9 && isColValid {
		isNumValid = true
		// fmt.Println("------>isNumValid = true =", isNumValid)
	} else {
		isNumValid = false
		// fmt.Println("##########>isNumValid = false =", isNumValid)
	}
	//--------block of cheking if the number in the subgrid are unique or not----------

	var isNumunique bool

	


	//-----------set up the result based on the isNumValid----------
	if isNumValid && isRowValid && isPositionEmpty {
		result = true
		 fmt.Println("------4444444--->result = true", result)
	} else {
		result = false
		//fmt.Println("------4444444--->result = false", result)
	}
	//------------------------------------
//sudo code : subgrig1, subgrid2,...,subgrid9

  fmt.Println()
	//--------------------------------------
	return result
}
