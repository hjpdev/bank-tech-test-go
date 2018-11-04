# bank-tech-test-go
This is an attempt to do a tech test completed a few weeks earlier in Java, but this time in GoLang.
Props to Brendon Hogger and his tutorial [here](https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial) for helping me get started.

### Acceptance Criteria

**Given** a client makes a deposit of 1000 on 10-01-2012  
**And** a deposit of 2000 on 13-01-2012  
**And** a withdrawal of 500 on 14-01-2012  
**When** she prints her bank statement  
**Then** she would see

```
  date || credit || debit || balance
  14/01/2012 || || 500.00 || 2500.00
  13/01/2012 || 2000.00 || || 3000.00
  10/01/2012 || 1000.00 || || 1000.00
```

To run the main file: ``` go run ./cmd/main.go ```
To run the tests, go to the folder that contains the tests and: ``` go test -v ```
To run the benchmark tests, again go to the folder and: ``` go test -bench . ```

## User Stories

```
  As a customer,
  so I know how much money I have,
  I want to be able to see my balance. ✔

  As a customer,
  so I can increase my balance,
  I want to be able to make deposits.

  As a customer,
  so I can spend my money,
  I want to be able to make withdrawals. ✔
  
  As a customer,
  so I can pay for unpredicted expenditures,
  I want an overdraft to tide me over.
  
  As a customer,
  so I don't exceed my overdraft,
  I want to be stopped withdrawing over this amount.

  As a customer,
  so I can see previous transactions,
  I want to be able to view my statement.
  
  As a bank,
  so we can incentivise customers to open accounts,
  we want the ability to open them with money already in.
```
