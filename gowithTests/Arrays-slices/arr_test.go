package main

import (
	"reflect"
	"testing"
	//"slices"
	//"github.com/golang/exp/tree/master/slices"
)

// =======================Test_Sum==================================
func Test_Sum(t *testing.T) {
	//----------
	t.Run("t1-sum of input := [5]int{1, 2, 3, 4, 5}", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5} //input
		act := Sum(input)             //Act
		exp := 15                     //Expectation
		assert_Sum(t, act, exp)       //Assertion
	})
	//----------
	t.Run("t2-test negative numbers, sum of input := [5]int{1, 0, 1, -1, -2}", func(t *testing.T) {
		input := []int{1, 0, 1, -1, -2} //input
		act := Sum(input)               //Act
		exp := -1                       //Expectation
		assert_Sum(t, act, exp)         //Assertion
	})
	//------------
	t.Run("t3-sum of input := [2]int{1, 2}", func(t *testing.T) {
		input := []int{1, 2}    //input parameter
		act := Sum(input)       //Act
		exp := 3                //Expectation
		assert_Sum(t, act, exp) //Assertion
	})
	//------------
	//-----END Test_Sum---------
}

// ----------------assert_Sum-----------------------
func assert_Sum(t testing.TB, act, exp int) {
	t.Helper()
	if act != exp {
		// Use %q for strings or specific verbs such as %d for integers and %f for float64.
		t.Errorf("----------> got %d BUT expected %d <----------", act, exp)
	}
}

// =======================Test_SumAll==================================
func Test_SumAll(t *testing.T) {
	// ----------
	t.Run("t1-sum of notheing []int{}", func(t *testing.T) {

		act := SumAll([]int{})     //Act
		exp := []int{0}            //Expectation
		assert_SumAll(t, act, exp) //Assertion
	})
	//------------
	t.Run("t2-sum of notheing []int{1}", func(t *testing.T) {

		act := SumAll([]int{1})    //Act
		exp := []int{1}            //Expectation
		assert_SumAll(t, act, exp) //Assertion
	})
	//------------
	t.Run("t3-sum of slice 1 item include 2 number []int{1,2}", func(t *testing.T) {

		act := SumAll([]int{1, 2}) //Act
		exp := []int{3}            //Expectation
		assert_SumAll(t, act, exp) //Assertion
	})
	//------------
	t.Run("t3-sum of slice of 2 item include 1 number []int{1},[]int{1}", func(t *testing.T) {

		act := SumAll([]int{1}, []int{1}) //Act
		exp := []int{1, 1}                //Expectation
		assert_SumAll(t, act, exp)        //Assertion
	})
	//------------
	t.Run("t4-sum of slice of 2 item include 2 number []int{1,2},[]int{4,6}", func(t *testing.T) {

		act := SumAll([]int{1,2}, []int{4,6}) //Act
		exp := []int{3, 10}                //Expectation
		assert_SumAll(t, act, exp)        //Assertion
	})
	//------------
} //END
// -----------assert_SumAll----------
func assert_SumAll(t testing.TB, act, exp []int) {
	t.Helper()
	//if !slices.Equal(act, exp) {
	if !reflect.DeepEqual(act, exp) {
		//slices.Equal
		t.Errorf("----------> got %v BUT expected %v <----------", act, exp)
	}
}

//==============================================================
