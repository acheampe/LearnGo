package main

//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors

import (
	"errors"
	"fmt"
)

// creating new types from existing types
type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

// w is a variable that takes a copy of Wallet. Thus we must reference Wallet via (*)
// to update main address
func (w *Wallet) Balance() Bitcoin {
	return w.balance
	// return (*w).balance in go, we do not have to have an explicit dereferencing
}
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount

}

// In Go, errors are values
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Refactor code
