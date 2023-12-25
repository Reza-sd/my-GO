package main

import (
  "fmt"
)

func main() {
  //----------------------------------------
  letters := []string{"A", "B", "C", "D"}
  for index, value := range letters {
    fmt.Println("Index:", index, "Value:", value)
  }
  //-----------------------------------------
  addressBook := map[string]string{
    "John":   "12 Main St",
    "Janet":  "56 Pleasant St",
    "Jordan": "88 Liberty Ln",
  }
  for key, value := range addressBook {
    fmt.Println("Name:", key, "Address:", value)
  }
  // --------------------------------------------
}
