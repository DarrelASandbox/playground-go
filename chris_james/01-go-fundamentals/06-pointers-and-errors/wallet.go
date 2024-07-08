package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

// Declare methods on the new type
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then
// it is private outside the package it's defined in.
type Wallet struct {
	balance Bitcoin
}

// In Go, when you call a function or a method the arguments are copied.
// Hence, set a pointer to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// The `var` keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("you're too poor")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
