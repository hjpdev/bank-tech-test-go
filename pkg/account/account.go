package account

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Account struct {
	balance decimal.Decimal
}

func NewAccount(initialBalance float64) *Account {
	return &Account{
		balance: decimal.NewFromFloat(initialBalance),
	}
}

func (a *Account) Balance() decimal.Decimal {
	return a.balance
}

func (a *Account) Withdraw(amount float64) error {
	if decimal.NewFromFloat(amount).GreaterThan(a.balance) {
		return errors.New("Insufficient Balance")
	}
	a.balance = a.balance.Sub(decimal.NewFromFloat(amount))
	return nil
}

func (a *Account) Deposit(amount float64) {
	a.balance = a.balance.Add(decimal.NewFromFloat(amount))
}
