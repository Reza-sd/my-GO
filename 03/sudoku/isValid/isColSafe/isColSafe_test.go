package main

import "testing"

// ===============================================
/*
It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
The test function takes one argument only t *testing.T
In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
*/

// ===============================================
func Test_IsColSafe(t *testing.T) {
	//------------

	t.Run("1-return bool (check if bool is true)", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{}
		row := 1
		col := 1
		nb := 0
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})

	//----------------------------------
	t.Run("2-return bool (check if bool is false )", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{1}}
		row := 1
		col := 1
		nb := 1
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})

	//----------------------------------
	t.Run("3-return bool (check if col is empty put 1 in row 1 )", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{}
		row := 1
		col := 1
		nb := 1
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		want := true
		assertCorrectMessage(t, got, want)
	})

	//----------------------------------
	t.Run("4- if the first cell board is full, can I put a new num? )", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{8}}
		row := 1
		col := 1
		nb := 1
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		want := false
		assertCorrectMessage(t, got, want)
	})

	//----------------------------------
	t.Run("5- if putting 1 in cell 2 safe when cell 1 is full?", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{8}}
		row := 2
		col := 1
		nb := 1
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		exp := true
		assertCorrectMessage(t, got, exp)
	})

	//----------------------------------
	t.Run("6- is row in [1...9]?", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{8}}
		row := 200
		col := 1
		nb := 1
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})

	//----------------------------------
	t.Run("7- is col in [1...9]?", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{8}}
		row := 2
		col := 10
		nb := 1
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})

	//----------------------------------
	t.Run("8- is number higher than 9?", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{8}}
		row := 2
		col := 1
		nb := 10
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	// ----------------------------------
	t.Run("9- is number less than 1?", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{8}}
		row := 2
		col := 1
		nb := -1
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	// ----------------------------------
	t.Run("10- check if i can put number 1 in position 6", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{8}, {4}, {2}, {3}, {5}, {}, {6}, {7}, {9}}
		row := 6
		col := 1
		nb := 1
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		exp := true
		assertCorrectMessage(t, got, exp)
	})
	// ----------------------------------
	t.Run("11- check if i can put number 4 in position 6", func(t *testing.T) {
		//-------setup-----
		board := [9][9]int{{8}, {4}, {2}, {3}, {5}, {}, {6}, {7}, {9}}
		row := 6
		col := 1
		nb := 4
		//-------ACT--------
		got := IsColSafe(board, row, col, nb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	// ----------------------------------

	//======================================
// ----------------------------------
t.Run("10- check if i can put number 1 in position 6", func(t *testing.T) {
	//-------setup-----
	board := [9][9]int{{8},{4},{2},{3},{5},{},{6},{7},{9}}
	row := 6
	col := 1
	nb := 1
	//-------ACT--------
	got := IsColSafe(board, row, col, nb)
	//------expectation------
	exp := true
	assertCorrectMessage(t, got, exp)
})
// ----------------------------------
t.Run("11- check if i can put number 4 in position 6", func(t *testing.T) {
	//-------setup-----
	board := [9][9]int{{8},{4},{2},{3},{5},{},{6},{7},{9}}
	row := 6
	col := 1
	nb := 4
	//-------ACT--------
	got := IsColSafe(board, row, col, nb)
	//------expectation------
	exp := false
	assertCorrectMessage(t, got, exp)
})
// ----------------------------------



//======================================

}

// =============asserion function ============
func assertCorrectMessage(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// ===============================================
