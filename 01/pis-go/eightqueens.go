package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

// -------------------------------------------
func EightQueens() {
	myBoard := makeBoard(8)
	myBoard[0][0] = 1 // queen starts at 1

	// myBoard[0][1] = 1
	printBoard(myBoard)
	// issafe
	fmt.Println(isSafeRow(1, myBoard))
	fmt.Println(isSafeRow(2, myBoard))

	fmt.Println(isSafeColumn(1, myBoard))
	fmt.Println(isSafeColumn(2, myBoard))
	z01.PrintRune('\n')
}

//------------isSafeColumn---------

func isSafeColumn(column int, board [][]int) bool {
	sunColumn := 0
	columnIndex := column - 1

	for i := 0; i < len(board); i++ {
		sunColumn += board[i][columnIndex]
	}

	if sunColumn == 0 {
		return true
	}
	return false
}

//---------------------isSafeRow-------------------

func isSafeRow(row int, board [][]int) bool {
	sumRow := 0
	rowIndex := row - 1

	for i := 0; i < len(board[rowIndex]); i++ {
		sumRow += board[rowIndex][i]
	}

	if sumRow == 0 {
		return true
	}
	return false
}

// --------------printBoard-----------------------
func printBoard(a [][]int) {
	for i := 0; i < len(a[0]); i++ {
		fmt.Println(a[i])
	}
}

// -----------makeBoard----------------
func makeBoard(s int) [][]int {
	rows, cols := s, s

	a := make([][]int, rows)
	for i := range a {
		a[i] = make([]int, cols)
	}
	return a
}

//------------------------------------
