package main

import "fmt"

type Account struct {
	name          string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *Account) Withdraw(value float64) string {
	if value < 0 || value > c.balance {
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

	// New objects with pointers
	var account1 *Account
	account1 = new(Account)
	// account1 := new(Account) // Declaração direta de ponteiro (2 linhas acima)
	account1.name = "Caio Baeta"
	account1.agencyNumber = 1234
	account1.accountNumber = 5678
	account1.balance = 1000.00

	var account2 *Account
	account2 = new(Account)
	// account2 := new(Account) // Declaração direta de ponteiro (2 linhas acima)
	account2.name = "Caio Baeta"
	account2.agencyNumber = 1234
	account2.accountNumber = 5678
	account2.balance = 1000.00
	fmt.Println(account1)               // Pointer
	fmt.Println(account2)               // Pointer
	fmt.Println(*account1)              // Value of pointer
	fmt.Println(*account2)              // Value of pointer
	fmt.Println(&account1)              // Address of pointer
	fmt.Println(&account2)              // Address of pointer
	fmt.Println(account1 == account2)   // false
	fmt.Println(&account1 == &account2) // false
	fmt.Println(*account1 == *account2) // true
}
