package main

import "fmt"

// Example function that takes a function as an argument
func applyFunction(value int, f func(int) int) int {
  return f(value)
}

// Example of a function that can be passed as an argument
func square(x int) int {
  return x * x
}

// Another example function
func double(x int) int {
  return x * 2
}

func main() {
  // Example 1: Using the square function with applyFunction
  result1 := applyFunction(5, square)
  fmt.Println("Square:", result1)

  // Example 2: Using the double function with applyFunction
  result2 := applyFunction(7, double)
  fmt.Println("Double:", result2)

  // Example 3: Define and use an anonymous function
  result3 := applyFunction(3, func(x int) int {
    return x + 1
  })
  fmt.Println("Anonymous Function:", result3)
}
