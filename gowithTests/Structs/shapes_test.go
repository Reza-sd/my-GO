package main

import (
	"testing"
)

// =======================Test_Perimeter==================================
func Test_Perimeter(t *testing.T) {
	//----------
	t.Run("t1-calculate the perimeter of a rectangle given a height and width", func(t *testing.T) {

		act := Perimeter(10.0, 10.0)  //Act
		exp := 40.0                   //Expectation
		assert_Perimeter(t, act, exp) //Assertion
	})
}

// =========================Test_Area=============================
func Test_Area(t *testing.T) {
	//----------
	t.Run("t1-calculate the Area of a rectangle given a height and width", func(t *testing.T) {

		act := Area(12.0, 6.0)        //Act
		exp := 72.0                   //Expectation
		assert_Perimeter(t, act, exp) //Assertion
	})
}

// ----------------assert_Perimeter-----------------------
func assert_Perimeter(t testing.TB, act, exp float64) {
	t.Helper()
	if act != exp {
		// Use %q for strings or specific verbs such as %d for integers and %f for float64.
		//The f is for our float64 and the .2 means print 2 decimal places.
		t.Errorf("---> got %.2f BUT expected %.2f <---", act, exp)
	}
}

//----------------------------------------------
