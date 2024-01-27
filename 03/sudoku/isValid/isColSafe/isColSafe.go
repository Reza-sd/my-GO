package main

// import (
// 	"fmt"
// )

func IsColSafe(board [9][9]int, row, col, nb int) bool {
	var isNotNumExit bool   // check if the col is safe to put nb
	var isEmpty bool = true // check if the cell is empty
	isNotNumExit = true     // default value of col.
	var cell int            // the value of cell on the grid
	var isNumValid bool     // check is nb=[1..9]
	//--------check if row and col are valid-----------
	var isRow_Col_valid bool
	if row >= 1 && row <= 9 && col >= 1 && col <= 9 {
		isRow_Col_valid = true
	} else {
		isRow_Col_valid = false
	}
	//----------if number valid [1-9]-------------------
	if nb >= 1 && nb <= 9 {
		isNumValid = true
	} else {
		isNumValid = false
	}

	//-------isempty-------------
	if isRow_Col_valid {
		if board[row-1][col-1] == 0 {
			isEmpty = true

		} else {
			isEmpty = false
		}
	}

	//--------------------
	if isNumValid && isEmpty && isRow_Col_valid { //check if number is between 1 to 9
		//--------Seaching for cells------------------
		for i := 0; i < 9; i++ { // i represent the row
			cell = board[i][col-1]

			//-------check if any cell value is equal to number(nb)---
			if cell == nb {
				//fmt.Println("not empty,col=", col, "row=", row)
				isNotNumExit = false
				break
			}

		}
		//---------------------------
	} else { //0
		isNotNumExit = false
	}
	//-------------------
	return isNotNumExit
}
