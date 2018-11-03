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
	}

	for _, test := range withdrawalTests {
		acc := NewAccount(test.initialBalance)
		acc.Withdraw(test.withdrawal)
		if result := acc.Balance(); result != test.expectedBalance {
			t.Errorf("Test failed -> Initial balance: %f; Withdrew: %f, Expected balance: %f; Actual balance: %f", test.initialBalance, test.withdrawal, test.expectedBalance, result)
		}
	}
}
