package main

import "fmt"
/*
In Go, arrays are passed into functions by value, which means that changes remain local to the function. Slices, which are pointers, are passed by reference so changes affect the original variable. To modify an array within a function, it must be converted into a slice.
*/
func main() {
  myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
//Changes to the slice parameter will be permanent
  changeLastElement(myTutors, "Bobby")
  fmt.Println("main slice=", myTutors)

  // Changes to the array will only be local to the function.
  myArr :=[4]string{"a", "b", "c", "d"}
  changeFirst(myArr , "ooooo")
  fmt.Println("main Array=", myArr)

}
//-------------------------------------------------
func changeLastElement(tut []string, value string) {

  tut[len(tut)-1] = value
  fmt.Println("func slice=", tut)
}
//--------------------------------
func changeFirst(array [4]string, value string) {
    array[0] = value
  fmt.Println("func Array=", array)
}
//--------------------------------
