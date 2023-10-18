package main

import "fmt"

type Account struct {
	name          string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *Account) Withdraw(value float64) string {
	if value > 0 && value > c.balance {
		return "Can't withdraw this value"
	}
	c.balance -= value
	return fmt.Sprintf("Withdrawal of $%.2f successful", value)
}

func main() {
	account := new(Account)
	account.name = "Caio Baeta"
	account.agencyNumber = 1234
	account.accountNumber = 5678
	account.balance = 1000.00

	fmt.Println(account)
	fmt.Println(account.Withdraw(800))
	fmt.Println(account)
	fmt.Println(account.Withdraw(800))
	fmt.Println(account)
}
