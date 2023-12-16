package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  a := time.Now()
  b := a.UnixNano()

  fmt.Println("time.Now()=",a)
  fmt.Println("time.Now().UnixNano()=",b)
  rand.Seed(b)
  fmt.Println("Random Number=",rand.Intn(100))
}