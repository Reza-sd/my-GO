package main

import (
	"testing"
	// "fmt"
)

// ===============================================
/*
It needs to be in a file with a name like xxx_test.go
The test function must start with the word Test
The test function takes one argument only t *testing.T
In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
*/

// ===============================================
func Test_IsRowSafe(t *testing.T) {
	//-----------------------------------------------------------------------------

	t.Run("1-check if the row argument is [1..9](check if int)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{}
		myrow := 10
		mycol := 1
		mynb := 1

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("2-check if the col argument is [1..9](check if int)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{}
		myrow := 1
		mycol := 10
		mynb := 1

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("3-check if the number argument is [1..9](check if int)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{}
		myrow := 1
		mycol := 1
		mynb := 10

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})

	//-----------------------------------------------------------------------------

	t.Run("4-check if the number argument is [1..9](check if int)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{}
		myrow := 1
		mycol := 1
		mynb := 10

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("5-check if the row 1 col 1 is empty and put 1 there)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{}
		myrow := 1
		mycol := 1
		mynb := 1

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := true
		assertCorrectMessage(t, got, exp)
	})

	//-----------------------------------------------------------------------------

	t.Run("6-check if the row 1 col 1 is not empty and put 1 there)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{{1}}
		myrow := 1
		mycol := 1
		mynb := 1

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("7-check if the row 1 col 6 is not empty and put 4 there)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{{1, 3, 2, 5, 7, 0, 8, 9, 6}}
		myrow := 1
		mycol := 6
		mynb := 4

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := true
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("8-it should return false if the address is an occupied cell)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{{1, 3, 2, 5, 0, 7, 8, 9, 6}}
		myrow := 1
		mycol := 6
		mynb := 4

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("9-it should return false if a number is duplicated on the same row)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{{1}}
		myrow := 1
		mycol := 2
		mynb := 1

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("10-it should return true if my number doesn't exist on same row and in an empty position)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{{1, 3, 2, 5, 0, 7, 8, 9, 6}}
		myrow := 1
		mycol := 5
		mynb := 4

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := true
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("11-it should return false if the row has only one empty space and is not unique number)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{{1, 3, 2, 5, 0, 7, 8, 9, 6}}
		myrow := 1
		mycol := 5
		mynb := 9

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})

	//-----------------------------------------------------------------------------

	t.Run("12-it should return false if the row has only one empty space and is not unique number)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{{1, 3, 2, 5, 0, 7, 8, 9, 6}}
		myrow := 1
		mycol := 5
		mynb := 9

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//-----------------------------------------------------------------------------

	t.Run("13-it should return false if the board is invalid)", func(t *testing.T) {
		//-------setup-----
		myboard := [9][9]int{{1, 300, 2, 5, 0, 7, 8, 0, 9}}
		myrow := 1
		mycol := 5
		mynb := 4

		//-------ACT--------
		got := IsRowSafe(myboard, myrow, mycol, mynb)
		//------expectation------
		exp := false
		assertCorrectMessage(t, got, exp)
	})
	//--------------------END---------------------------------------------------------
}

// =============asserion function ============
func assertCorrectMessage(t testing.TB, got, exp bool) {
	t.Helper()
	if got != exp {
		t.Errorf("got %v want %v", got, exp)
	}
}

// ===============================================
// assertbool
func assertbool(t testing.TB, got, exp bool) {
	t.Helper()

}

//===============================
