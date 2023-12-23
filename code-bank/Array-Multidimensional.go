package main

import "fmt"

func main() {
  // Create an array of three arrays containing two integers
  var twoDim1 [3][2]int

  // Accessing the first item of the first array
  twoDim1[0][0] = 1

  // Accessing the second item of the second array
  twoDim1[1][1] = 2

  // Accessing the first item of the third array
  twoDim1[2][0] = 3

  fmt.Println(twoDim1)

  // Create an array of two arrays containing two floating-point numbers
  twoDim2 := [2][2]float64{{3.14, 2.72}, {2.1, 3.7}}

  fmt.Println(twoDim2)
}
