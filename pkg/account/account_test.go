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

func TestTableWithdraw(t *testing.T) {
	var withdrawalTests = []struct {
		initialBalance  float64
		withdrawal      float64
		expectedBalance float64
	}{
		{100.00, 50.00, 50.00},
		{100.00, 34.65, 65.35},
		{100.00, 78.21, 21.79},
		{100.00, 22.34, 77.66},
		{100.00, 56.94, 43.06},
		{100.00, 11.45, 88.55},
	}

	for _, test := range withdrawalTests {
		acc := NewAccount(test.initialBalance)
		acc.Withdraw(test.withdrawal)
		if result := acc.Balance(); result != test.expectedBalance {
			t.Errorf("Test failed -> Initial balance: %f; Withdrew: %f, Expected balance: %f; Actual balance: %f", test.initialBalance, test.withdrawal, test.expectedBalance, result)
		}
	}
}
