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
