package account

type Account struct {
	balance float64
}

func NewAccount(initialBalance float64) *Account {
	return &Account{
		balance: initialBalance,
	}
}

func (a *Account) Balance() float64 {
	return a.balance
}
