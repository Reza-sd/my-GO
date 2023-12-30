package iteration

import "testing"

// =========================================================
func Test_Add(t *testing.T) {
	//----------
	t.Run("t1-one chatacter to 5 repeat", func(t *testing.T) {
		act := Repeat("a", 5)           //Act - my act
		exp := "aaaaa"                  //Expecgoted - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
	t.Run("t2-two chatacter to 5 repeat", func(t *testing.T) {
		act := Repeat("bc", 5)          //Act - my act
		exp := "bcbcbcbcbc"             //Expected - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
	t.Run("t3-three chatacter to 5 repeat", func(t *testing.T) {
		act := Repeat("e2f", 5)         //Act - my act
		exp := "e2fe2fe2fe2fe2f"        //Expected - my expectation
		assertCorrectValue(t, act, exp) //assertion - my assertion
	})
	//-----------
	t.Run("t4-twe chatacter to 3 repeat", func(t *testing.T) {
		act := Repeat("ab", 3)          //Act - my act
		exp := "ababab"                 //Expected - my expectation
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
func Benchmark_Repeat(b *testing.B) {
	b.Run("b1-one chatacter to 10 repeat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 10)
		}
	})
	//-----------
	b.Run("b2-two chatacter to 15 repeat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("bc", 15)
		}
	})
	//-----------
	b.Run("b3-three chatacter to 20 repeat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("e2f", 20)
		}
	})
	//-----------
}

// ====================================================
