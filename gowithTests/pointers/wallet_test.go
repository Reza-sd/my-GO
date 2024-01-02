package main

import (
	"fmt"
	"testing"
)

// =========================Test=============================
func Test_Area(t *testing.T) {
	t.Run("t1-calculate the Area of a rectangle given a height and width", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		want := 10
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})
}

// =========================END=============================
