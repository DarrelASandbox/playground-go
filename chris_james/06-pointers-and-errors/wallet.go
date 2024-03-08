package wallet

import "fmt"

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
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
