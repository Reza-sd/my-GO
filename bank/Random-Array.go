package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  guestArr := [3]string{"sarah", "jack", "john"}
  catArr := [3]string{"Toy", "Crate", "Box"}
  fmt.Println("guests list", guestArr)
  fmt.Println("Cat objects", catArr)
  rand.Seed(time.Now().UnixNano())

  //To pass our array into getRandomElement we will need to convert it into a slice.
  // sliceVersion := array[:]

  fmt.Println(getRandomElement(catArr[:]))
  fmt.Println(getRandomElement(guestArr[:]))
}

func getRandomElement(Arr []string) string {
  return Arr[rand.Intn(len(Arr))]
}
