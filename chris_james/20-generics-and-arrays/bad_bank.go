package main

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	return 0.0
}
