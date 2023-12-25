package main

import (
  "fmt"
  "math/rand"
  "time"
)

// -----------------------------------------
func randomNumberGen(min float32, max float32) float32 {
  return min + (max-min)*rand.Float32()
}

// -----------------------------------------
// Task implementation goes here
type Stock struct {
  name  string
  price float32
  //updateStock()
}

// -----------------------------------------
func (mystock *Stock) updateStock() {
  change := randomNumberGen(-10000, 10000)
  mystock.price += change
}

// -----------------------------------------
func displayMarket(market []Stock) {
  fmt.Println()
  for _, value := range market {
    fmt.Println(value)
  }
}

// ========================================
func main() {
  rand.Seed(time.Now().UnixNano())
  // Function calls go here
  stockMarket := []Stock{
    {"GOOG", 2313.50},
    {"AAPL", 157.28},
    {"FB", 203.77},
    {"TWTR", 50.00},
  }

  //fmt.Println("\n", stockMarket)
  displayMarket(stockMarket)
  stockMarket[0].price = 2500.3
  stockMarket[1].updateStock()

  displayMarket(stockMarket)

}

//========================================
