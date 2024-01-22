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
func TestHello(t *testing.T) {
	//------------
	t.Run("1-saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	//------------
	t.Run("2-empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World "
		assertCorrectMessage(t, got, want)
	})
	// ------------
}

// -----------------------------------------------------
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// ===============================================
