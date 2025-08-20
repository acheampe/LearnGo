package main

import "fmt"

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

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
