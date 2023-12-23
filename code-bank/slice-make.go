package main

import "fmt"

func main() {

  // Declare and initialize a slice in one line
  n := []int{1, 2, 3, 4, 5}
  fmt.Println("New Slice:", n, "Length:", len(n), "Capacity:", cap(n), "\n")

  // Create a slice from an existing array
  a := [7]int{1, 2, 3, 4, 5, 6, 7}
  s := a[2:5]
  fmt.Println("Array:", a, "Length:", len(a), "Capacity:", cap(a))
  fmt.Println("Slice from Array:", s, "Length:", len(s), "Capacity:", cap(s), "\n")

  // Create an empty slice, setting length and capacity
  m := make([]int, 3, 8)
  fmt.Println("Empty Slice:", m, "Length:", len(m), "Capacity:", cap(m), "\n")

}
