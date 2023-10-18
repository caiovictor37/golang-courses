package accounts

import "fmt"

type Account struct {
	Name          string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *Account) Withdraw(value float64) string {
	if value < 0 || value > c.Balance {
		return "Can't withdraw this value"
	}
	c.Balance -= value
	return fmt.Sprintf("Withdrawal of $%.2f successful", value)
}

func (c *Account) Deposit(value float64) string {
	if value < 0 {
		return "Can't deposit this value"
	}
	c.Balance += value
	return fmt.Sprintf("Deposit of $%.2f successful", value)
}

func (c *Account) TransferToAccount(account *Account, value float64) string {
	if value < 0 || value > c.Balance {
		return "Can't transfer this value"
	}
	account.Balance += value
	c.Balance -= value
	return fmt.Sprintf("Transfer of $%.2f from "+c.Name+" to "+account.Name+" successful", value)
}
