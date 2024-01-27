package main

// import "fmt"

// Set-a-2D-array-of-9x9-to-represent-sudoku-table.
func IsRowSafe(board [9][9]int, row, col, nb int) bool {

	//-----------------------setup----------------------
	//Set-variables
	var IsEmpty bool //checks if a row col intersecition is 0 or not (empty)
	var result bool
	var IsColVal bool // save if column is valid or not
	var IsRowVal bool //save if row is valid or not
	var IsNbVal bool  //save if number is valid or not
	var isRowAcceptable bool
	//-----------validate-the-row-----------------
	if row >= 1 && row <= 9 {

		IsRowVal = true //true if the number is between 1 and 9
		// fmt.Println("<----------")
	} else {
		IsRowVal = false //else False if the number is not in between 1 and 9
		result = false
		// fmt.Println("<------***********----", IsRowVal)
	}

	//----------------check the validity of board---------------------
	counter := 0
	if IsRowVal {
		
		
		var vcell int
		for i := 0; i <= 8; i++ { // iterates through the board
			vcell = board[row-1][i]

			//------check if elements of board are valid-----
			if vcell >= 0 && vcell <= 9 {
				counter ++ //count through the loop
				
			} 
			//-----------
			// fmt.Println("cell=", vcell, "isRowAcceptable=", isRowAcceptable)
		}

	}
	if counter == 9 {
		isRowAcceptable = true
	} else {
		isRowAcceptable = false
	}

	// fmt.Println("isRowAcceptable=", isRowAcceptable)

	// true?????
	//--------------------------------------------------

	//-----------validate the column-----------------
	if col >= 1 && col <= 9 && isRowAcceptable {
		IsColVal = true //true if the number is between 1 and 9
	} else {
		IsColVal = false //else False if the number is not in between 1 and 9
		result = false
	}
	// -----------validate-the-number-----------------

	if nb >= 1 && nb <= 9 && isRowAcceptable {
		IsNbVal = true //true if the number is between 1 and 9
	} else {
		IsNbVal = false //else False if the number is not in between 1 and 9
		result = false
	}
	// -------------------check if cell empty or no--------------------------------
	if IsColVal && IsRowVal && IsNbVal && isRowAcceptable {

		if board[row-1][col-1] == 0 {
			IsEmpty = true
			result = true
			// fmt.Println("11111",isRowAcceptable)
		} else {
			IsEmpty = false
			result = false
		}

	}

	//------------------------not-same-value and unique--------------------------------
	if IsColVal && IsRowVal && IsNbVal && IsEmpty && isRowAcceptable {
		var cell int
		// fmt.Println(board)
		for i := 0; i <= 8; i++ {
			cell = board[row-1][i]
			if cell == nb {
				result = false

			}
		}
	}
	//-------------set result value if board is not valid----------------
	if !isRowAcceptable {
		result = false
		
	}



	
	//------------------------------

	return result
}
