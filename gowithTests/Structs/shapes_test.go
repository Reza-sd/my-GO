package main

import (
	"testing"
)

// =======================Test_Perimeter==================================
func Test_Perimeter(t *testing.T) {
	//----------
	t.Run("t1-calculate the perimeter of a rectangle given a height and width", func(t *testing.T) {
    rectangle1 := Rectangle{10.0, 10.0}
		act := Perimeter(rectangle1)  //Act
		exp := 40.0                   //Expectation
		assert_Perimeter(t, act, exp) //Assertion
	})
	//---------
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

// =========================Test_Area=============================
func Test_Area(t *testing.T) {
	//----------
	t.Run("t1-calculate the Area of a rectangle given a height and width", func(t *testing.T) {
    rectangle1 := Rectangle{12.0, 6.0}
		act := Area(rectangle1)   //Act
		exp := 72.0              //Expectation
		assert_Area(t, act, exp) //Assertion
	})
	//-----------
}

// ------------------assert_Area-----------------
func assert_Area(t testing.TB, act, exp float64) {
	t.Helper()
	if act != exp {
		// Use %q for strings or specific verbs such as %d for integers and %f for float64.
		//The f is for our float64 and the .2 means print 2 decimal places.
		t.Errorf("---> got %.2f BUT expected %.2f <---", act, exp)
	}
}

//=====================END===========================
