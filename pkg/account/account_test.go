package account

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestAccountZeroBalance(t *testing.T) {
	acc := NewAccount(0.00)
	if acc.Balance().Equal(decimal.NewFromFloat(0.00)) == false {
		t.Errorf("Test failed -> Expected balance to be zero, but was %c", acc.Balance())
	}
}

func TestAccountWithdraw(t *testing.T) {
	acc := NewAccount(100.00)
	acc.Withdraw(25.50)
	if acc.Balance().Equal(decimal.NewFromFloat(74.50)) == false {
		t.Errorf("Test failed -> Expected balance to be 74.50, but was %v", acc.Balance())
	}
}

func TestTableWithdraw(t *testing.T) {
	var withdrawalTests = []struct {
		initialBalance  float64
		withdrawal      float64
		expectedBalance decimal.Decimal
	}{
		{100.00, 50.00, decimal.NewFromFloat(50.00)},
		{100.00, 34.65, decimal.NewFromFloat(65.35)},
		{100.00, 78.21, decimal.NewFromFloat(21.79)},
		{100.00, 22.34, decimal.NewFromFloat(77.66)},
		{100.00, 56.94, decimal.NewFromFloat(43.06)},
		{100.00, 11.45, decimal.NewFromFloat(88.55)},
	}

	for _, test := range withdrawalTests {
		acc := NewAccount(test.initialBalance)
		acc.Withdraw(test.withdrawal)
		if result := acc.Balance(); result.Equal(test.expectedBalance) == false {
			t.Errorf("Test failed -> Initial balance: %f; Withdrew: %f, Expected balance: %v; Actual balance: %v", test.initialBalance, test.withdrawal, test.expectedBalance, result)
		}
	}
}
