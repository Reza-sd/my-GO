package integers

import "testing"

// =========================================================
func Test_Add(t *testing.T) {
	//----------
	t.Run("1-adding two integers", func(t *testing.T) {
		act := Add(1, 2)                //Act - my act
		exp := 3                        //Expected - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
}

// =============================================
func assertCorrectValue(t testing.TB, act, exp int) {
	t.Helper()
	if act != exp {
		t.Errorf("got %q want %q", act, exp)
	}
}

//-------------------------------
