package integers

import (
	//"fmt"
	"testing"
)

// =========================================================
func Test_Add(t *testing.T) {
	//----------
	t.Run("1-adding two zero", func(t *testing.T) {
		act := Add(0, 0)                //Act - my act
		exp := 0                        //Expected - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
	t.Run("2-adding two integers", func(t *testing.T) {
		act := Add(1, 2)                //Act - my act
		exp := 3                        //Expected - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
}

// =============================================
func assertCorrectValue(t testing.TB, act, exp int) {
	t.Helper()
	//fmt.Println("---> act=", act, "-- exp=", exp)
	if act != exp {
		t.Errorf("got %q want %q", act, exp)
	}
}

//-------------------------------
