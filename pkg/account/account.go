package account

import "math"

type Account struct {
	balance float64
}

func NewAccount(initialBalance float64) *Account {
	return &Account{
		balance: initialBalance,
	}
}

func (a *Account) Balance() float64 {
	bal := math.Floor(a.balance*100) / 100
	return bal
}
