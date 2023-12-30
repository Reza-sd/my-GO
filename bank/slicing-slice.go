package main

import (
  "fmt"
)

func main() {
  mySlice := []string{"Pepper", "Elodie", "Ruth", "Pyewacket", "George"}

  fmt.Printf("mySlice = %v\n", mySlice)
  fmt.Printf("mySlice[2:] = %v\n", mySlice[2:])
  fmt.Printf("mySlice[:2] = %v\n", mySlice[:2])
}
/*
mySlice = [Pepper Elodie Ruth Pyewacket George]
mySlice[2:] = [Ruth Pyewacket George]
mySlice[:2] = [Pepper Elodie]
*/