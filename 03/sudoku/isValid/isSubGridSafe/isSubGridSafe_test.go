package main

import (
	"fmt"
	"testing"
)

// ===============================================
/*
It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
The test function takes one argument only t *testing.T
In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
*/

// ===============================================
func Test_IsSubGridSafe(t *testing.T) {
	//------------------------------------------

	t.Run("1- return false if nb is not in [1...9]", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{}}
		row := 1
		col := 1
		nb := 0
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})
	//-----------------------------------------------------
	t.Run("2- return true if nb is in [1...9]", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{}}
		row := 1
		col := 1
		nb := 3
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		want := true
		assertCorrectMessage(t, got, want)
	})

	//--------------------------------------------------------
	t.Run("3- return false if nb is > 9", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{}}
		row := 1
		col := 1
		nb := 3200
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})

	//--------------------------------------------
	t.Run("4- return false if col are > 9 while other(row, nb) args are valid", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{}}
		row := 1
		col := 10
		nb := 3
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})
	//-------------------------------------------------------

	t.Run("5- return false if col are < 1 while other(row, nb) args are valid", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{}}
		row := 1
		col := 0
		nb := 3
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})

	//-------------------------------------------------------

	t.Run("6- return false if rows > 9 while other args are valid)", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{}}
		row := 10
		col := 1
		nb := 3
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})

	//-------------------------------------------------------

	t.Run("7- return false if rows < 1 while other args are valid)", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{}}
		row := 0
		col := 1
		nb := 3
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})

	//-------------------------------------------------------

	t.Run("8-  Return true if all cell values are 0 in and position is (1,1) and other args are valid)", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{}}
		row := 1
		col := 1
		nb := 1
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		want := true
		assertCorrectMessage(t, got, want)
	})

	//-------------------------------------------------------

	t.Run("9-Return false if the cell being filled in is not empty while other conditions are valid", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{1}}
		row := 1
		col := 1
		nb := 3
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})

	//-------------------------------------------------------
	t.Run("10- Return false is I try to fill in (2,2) while are valid and position (2,2) is not empty", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{1, 5}, {0, 3}}
		row := 2
		col := 2
		nb := 4
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//-------------------------------------------------------
	t.Run("11- Return true is I try to fill in (1,3) while are valid and position (1,3) is  empty", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{1, 5}, {0, 3}}
		row := 1
		col := 3
		nb := 4
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		exp := true
		assertCorrectMessage(t, got, exp)
	})
	//-------------------------------------------------------
	t.Run("12- Return true if num=4 in position (1,3) is unique in its subgrid while other agrs are valid", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{1, 5}, {0, 3}}
		row := 1
		col := 3
		nb := 4
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		exp := true
		assertCorrectMessage(t, got, exp)
	})
	//-------------------------------------------------------
	t.Run("13- Return false if num=3 in position empty (1,3) is not unique in its subgrid while other agrs are valid", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{1, 5},{0, 3}}
		/*
		 board := [9][9]int{{1,5,0, 0,0,0,0,0,0},
		                    {0,3,0, 0,0,0,0,0,0},
							{0,0,0, 0,0,0,0,0,0},

							{0,0,0,0,0,0,0,0,0},
							{0,0,0,0,0,0,0,0,0},
							{0,0,0,0,0,0,0,0,0},
							{0,0,0,0,0,0,0,0,0},
							{0,0,0,0,0,0,0,0,0},
							{0,0,0,0,0,0,0,0,0}

						}
		*/
		fmt.Println(board)
		row := 1
		col := 3
		nb := 3
		//-------ACT--------
		got := IsSubGridSafe(board, row, col, nb)
		//------expectation------
		exp := false //bcz i realise (1,3) blongs to subgrid-1
		// in subgrid -1 I check the cells , I find out value of (2,2) = num
		assertCorrectMessage(t, got, exp)
	})
	//-------------------------------------------------------

	//----------end--------------------------

}

// =============asserion function ============
func assertCorrectMessage(t testing.TB, got, exp bool) {
	t.Helper()
	if got != exp {
		t.Errorf("got %v want %v", got, exp)
	}
}

// ===============================================
