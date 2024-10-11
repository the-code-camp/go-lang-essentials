package main

import "fmt"

/**
TODO: Find abstractions and compose interfaces to fix the commented code

This is more of an interface composition exploration exercise
*/

// A current account can be used to deposit and withdraw
type CurrentAccount struct {
	AccountNumber int
	Balance       float64
}

func (s *CurrentAccount) Deposit(amount float64) {
	s.Balance += amount
}

func (s *CurrentAccount) Withdraw(amount float64) {
	if s.Balance >= amount {
		s.Balance -= amount
	} else {
		fmt.Println("Insufficient balance.")
	}
}

func (s *CurrentAccount) GenerateStatement() {
	fmt.Printf("\nCurrent Account Statement\nAccount Number: %d\nBalance: $%.2f\n", s.AccountNumber, s.Balance)
}

// A saving account can be used to deposit and withdraw
type SavingsAccount struct {
	AccountNumber int
	Balance       float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.Balance += amount
}

func (s *SavingsAccount) Withdraw(amount float64) {
	if s.Balance >= amount {
		s.Balance -= amount
	} else {
		fmt.Println("Insufficient balance.")
	}
}

func (s *SavingsAccount) GenerateStatement() {
	fmt.Printf("\nSavings Account Statement\nAccount Number: %d\nBalance: $%.2f\n", s.AccountNumber, s.Balance)
}

// A fixed deposit account is not withdrawable
type FixedDepositAccount struct {
	AccountNumber int
	Balance       float64
}

func (s *FixedDepositAccount) Deposit(amount float64) {
	s.Balance += amount
}

func (s *FixedDepositAccount) GenerateStatement() {
	fmt.Printf("\nFixed deposit Account Statement\nAccount Number: %d\nBalance: $%.2f\n", s.AccountNumber, s.Balance)
}

func (s SavingsAccount) String() string {
	return fmt.Sprintf("Account Id: %d having balance: %.2f", s.AccountNumber, s.Balance)
}

func main() {
	// Create instances of both accounts
	savings := &SavingsAccount{AccountNumber: 123, Balance: 1000}
	current := &CurrentAccount{AccountNumber: 999, Balance: 3000}
	fd := &FixedDepositAccount{AccountNumber: 456, Balance: 0}

	fmt.Println(savings)

	amount := 100.0
	withdrawFrom(savings, amount) //TODO: fix this
	withdrawFrom(current, amount) //TODO: fix this

	printStatement(fd)      //TODO: fix this
	printStatement(current) //TODO: fix this
	printStatement(savings) //TODO: fix this

	depositTo(savings, amount) //TODO: fix this
	depositTo(current, amount) //TODO: fix this

	printStatement(current) //TODO: fix this
	printStatement(savings) //TODO: fix this
	printStatement(fd)      //TODO: fix this

	withdrawAndPrintStatement(savings, amount) //TODO: fix this
	withdrawAndPrintStatement(current, amount) //TODO: fix this
}

type WithdrawablePrintable interface {
	Withdrawable
	StatementPrinter
}

func withdrawAndPrintStatement(x WithdrawablePrintable, amount float64) {
	x.Withdraw(amount)
	x.GenerateStatement()
}

type Withdrawable interface {
	Withdraw(float64)
}

func withdrawFrom(x Withdrawable, amount float64) {
	x.Withdraw(amount)
}

type StatementPrinter interface {
	GenerateStatement()
}

func printStatement(x StatementPrinter) {
	x.GenerateStatement()
}

type Deposit interface {
	Deposit(float64)
}

func depositTo(x Deposit, amount float64) {
	x.Deposit(amount)
}

/**
// TODO: fix below functions
func withdrawAndPrintStatement(x ?, amount float64) {
}

func withdrawFrom(x ?, amount float64) {
}

func depositTo(x ?, amount float64) {
}

func printStatement(x ?) {
}
*/
