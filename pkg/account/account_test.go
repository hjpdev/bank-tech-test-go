package account

import (
	"sync"
	"testing"

	"github.com/shopspring/decimal"
)

func TestAccountZeroBalance(t *testing.T) {
	acc := NewAccount(0.00)
	if acc.Balance().Equal(decimal.NewFromFloat(0.00)) == false {
		t.Errorf("Test failed -> Expected balance to be zero, but was %c", acc.Balance())
	}
}

/*
func BenchmarkAccount(b *testing.B) {
	acc := NewAccount(float64(b.N))
	for i := 0; i < b.N; i++ {
		acc.Withdraw(1.00)
	}
	if acc.Balance().Equal(decimal.NewFromFloat(0.00)) == false {
		b.Errorf("Balance wasn't zero: %v", acc.Balance())
	}
}
*/

const WORKERS = 10

func BenchmarkWithdrawals(b *testing.B) {
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	acc := NewAccount(float64(b.N))

	gbpPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initilised - zero values is ready to use
	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		// Lets WaitGroup know you're adding a goroutine
		wg.Add(1)
		// Spawns off founder worker as a closure
		go func() {
			// Marks worker done when function finishes
			defer wg.Done()
			for i := 0; i < gbpPerFounder; i++ {
				acc.Withdraw(1.0)
			}
		}()
	}
	// Waits for all workers to finish
	wg.Wait()

	if acc.Balance().Equal(decimal.NewFromFloat(0.00)) == false {
		b.Errorf("Balance wasn't zero: %v; Iterations: %d", acc.Balance(), b.N)
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
