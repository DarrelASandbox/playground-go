package main

type Transaction struct {
	From, To string
	Sum      float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

/*
This demonstrates the power of using concepts like `Reduce`.
The `NewBalanceFor` function appears more declarative, describing what happens rather than how it happens.
Often, when reading code, developers are navigating through numerous files,
attempting to understand what is happening rather than how it is implemented,
and this coding style supports that process effectively.

If a developer wishes to delve into the details,
they can do so and observe the business logic of `applyTransaction`
without concerning themselves with loops and mutating state; Reduce handles that separately.
*/
func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}

	return a
}
