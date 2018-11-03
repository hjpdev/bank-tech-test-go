package account

import (
	"testing"
)

func TestAccountZeroBalance(t *testing.T) {
	acc := NewAccount(0.00)
	if acc.Balance() != 0.00 {
		t.Errorf("Test failed -> Expected balance to be zero, but was %f", acc.Balance())
	}
}

func TestAccountWithdraw(t *testing.T) {
	acc := NewAccount(100.00)
	acc.Withdraw(25.50)
	if acc.Balance() != 74.50 {
		t.Errorf("Test failed -> Expected balance to be 74.50, but was %f", acc.Balance())
	}
}
