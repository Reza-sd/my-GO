package iteration

import "testing"

// =========================================================
func Test_Add(t *testing.T) {
	//----------
	t.Run("1-one chatacter to 5 repeat", func(t *testing.T) {
		act := Repeat("a")              //Act - my act
		exp := "aaaaa"                  //Expecgoted - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
	t.Run("2-two chatacter to 5 repeat", func(t *testing.T) {
		act := Repeat("bc")             //Act - my act
		exp := "bcbcbcbcbc"             //Expected - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
	t.Run("3-three chatacter to 5 repeat", func(t *testing.T) {
		act := Repeat("e2f")            //Act - my act
		exp := "e2fe2fe2fe2fe2f"        //Expected - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
}

// ==============assertion-function======================
func assertCorrectValue(t testing.TB, act, exp string) {
	t.Helper()
	if act != exp {
		t.Errorf("expected %q but got %q", exp, act)
	}
}

// =================Benchmark=========================
func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a")
  }
}
// ====================================================