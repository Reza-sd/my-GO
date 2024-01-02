package main  
import (
  "fmt"
  
)

type Wallet struct{
  balance int
}
//In Go, when you call a function or a method the arguments are copied.

func (w *Wallet) Deposit(amount int) {
  fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
  w.balance += amount
}

func (w *Wallet) Balance() int {
  //return (*w).balance  //same = automatically dereferenced
  return w.balance
}
