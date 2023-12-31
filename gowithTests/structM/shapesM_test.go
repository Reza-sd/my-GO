package shape333

import (
	"testing"
)

// =========================Test_Area=============================
func Test_Area(t *testing.T) {
	//----------
	t.Run("t1-calculate the Area of a rectangle given a height and width", func(t *testing.T) {
		rectangle1 := Rectangle{12.0, 6.0}
		act := rectangle1.Area() //Act
		exp := 72.0              //Expectation
		assert_Area(t, act, exp) //Assertion
	})
	//-----------
	//----------
	t.Run("t2-calculate the Area of a Circle given a radius", func(t *testing.T) {
		circle1 := Circle{10}
		act := circle1.Area()    //Act
		exp := 314.1592653589793 //Expectation
		assert_Area(t, act, exp) //Assertion
	})
	//----------END----
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
