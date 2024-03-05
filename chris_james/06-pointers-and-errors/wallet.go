package wallet

import "fmt"

// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then
// it is private outside the package it's defined in.
type Wallet struct {
	balance int
}

// In Go, when you call a function or a method the arguments are copied.
func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
