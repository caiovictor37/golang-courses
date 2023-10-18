package main

import (
	"fmt"
	"golang-courses/alura/2-object-oriented-go/3-returns-packages-visibility/accounts"
)

func main() {
	account := new(accounts.Account)
	account.Name = "Caio Baeta"
	account.AgencyNumber = 1234
	account.AccountNumber = 5678
	account.Balance = 1000.00

	fmt.Println(account)
	fmt.Println(account.Withdraw(800))
	fmt.Println(account)
	fmt.Println(account.Withdraw(800))
	fmt.Println(account)
	fmt.Println(account.Deposit(800))
	fmt.Println(account)
	fmt.Println(account.Deposit(-800))
	fmt.Println(account)

	account2 := new(accounts.Account)
	account2.Name = "Test User"
	account2.AgencyNumber = 4321
	account2.AccountNumber = 8765
	account2.Balance = 1000.00

	fmt.Println(account.TransferToAccount(account2, 500))
	fmt.Println(account)
	fmt.Println(account2)
	fmt.Println(account2.TransferToAccount(account, 200))
	fmt.Println(account)
	fmt.Println(account2)
	fmt.Println(account2.TransferToAccount(account, -200))
	fmt.Println(account)
	fmt.Println(account2)
	fmt.Println(account2.TransferToAccount(account, 20000))
	fmt.Println(account)
	fmt.Println(account2)
}
