package accounts

import (
	"fmt"
	"golang-courses/alura/2-object-oriented-go/bank/customers"
)

type SavingsAccount struct {
	Customer                               customers.Customer
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) Withdraw(value float64) string {
	if value < 0 || value > c.balance {
		return "Can't withdraw this value"
	}
	c.balance -= value
	return fmt.Sprintf("Withdrawal of $%.2f successful", value)
}

func (c *SavingsAccount) Deposit(value float64) string {
	if value < 0 {
		return "Can't deposit this value"
	}
	c.balance += value
	return fmt.Sprintf("Deposit of $%.2f successful", value)
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
