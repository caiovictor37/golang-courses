package accounts

import (
	"fmt"
	"golang-courses/alura/2-object-oriented-go/bank/customers"
)

type CurrentAccount struct {
	Customer      customers.Customer
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	if value < 0 || value > c.balance {
		return "Can't withdraw this value"
	}
	c.balance -= value
	return fmt.Sprintf("Withdrawal of $%.2f successful", value)
}

func (c *CurrentAccount) Deposit(value float64) string {
	if value < 0 {
		return "Can't deposit this value"
	}
	c.balance += value
	return fmt.Sprintf("Deposit of $%.2f successful", value)
}

func (c *CurrentAccount) TransferToAccount(account *CurrentAccount, value float64) string {
	if value < 0 || value > c.balance {
		return "Can't transfer this value"
	}
	account.balance += value
	c.balance -= value
	return fmt.Sprintf("Transfer of $%.2f from "+c.Customer.Name+" to "+account.Customer.Name+" successful", value)
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
