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
func TestSum(t *testing.T) {
	//------------
	
	t.Run("1-return 0 (check if int)", func(t *testing.T) {
		//-------setup-----
		a := 0
		b := 0
		//-------ACT--------
		got := Sum(a,b)
		//------expectation------
		want := 0
		assertCorrectMessage(t, got, want)
	})
	//--------------
	t.Run("2-return 1 if a=0 , b=1", func(t *testing.T) {
		got := Sum(0,1)
		exp := 1
		assertCorrectMessage(t, got, exp)
	})

	//--------------
	t.Run("3-return 2 a=1 , b=1", func(t *testing.T) {
		got := Sum(0,1)
		exp := 1
		assertCorrectMessage(t, got, exp)
	})
	//--------------
	t.Run("4-return 2 a=1 , b=1", func(t *testing.T) {
		got := Sum(0,1)
		exp := 1
		assertCorrectMessage(t, got, exp)
	})
	//--------------

}

// =============asserion function ============
func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// ===============================================
