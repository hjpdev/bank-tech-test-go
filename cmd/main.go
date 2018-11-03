package main

import (
	"fmt"

	"github.com/hjpugh/bank-tech-test-go/pkg/account"
)

func main() {
	acc := account.NewAccount(100.00)
	acc.Withdraw(26.27)
	fmt.Println(acc.Balance())
}
