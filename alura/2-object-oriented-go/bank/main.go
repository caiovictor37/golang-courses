package main

import (
	"fmt"
	"golang-courses/alura/2-object-oriented-go/bank/accounts"
)

func payInvoices(account accounts.Account, invoiceValue float64) string {
	return account.Withdraw(invoiceValue)
}

func main() {
	account := new(accounts.CurrentAccount)
	account.Customer.Name = "Caio Baeta"
	account.Customer.Cpf = "123.456.789-01"
	account.Customer.Profession = "Software Engineer"
	account.AgencyNumber = 1234
	account.AccountNumber = 5678
	account.Deposit(1000.00)

	fmt.Println(account)
	fmt.Println(account.Withdraw(800))
	fmt.Println(account)
	fmt.Println(account.Withdraw(800))
	fmt.Println(account)
	fmt.Println(account.Deposit(800))
	fmt.Println(account)
	fmt.Println(account.Deposit(-800))
	fmt.Println(account)

	account2 := new(accounts.CurrentAccount)
	account2.Customer.Name = "Test User"
	account2.Customer.Cpf = "987.654.321-09"
	account2.Customer.Profession = "Test Profession"
	account2.AgencyNumber = 4321
	account2.AccountNumber = 8765
	account2.Deposit(1000.00)

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

	savingsAccount := new(accounts.SavingsAccount)
	savingsAccount.Customer = account.Customer
	savingsAccount.AgencyNumber = account.AgencyNumber
	savingsAccount.AccountNumber = account.AccountNumber
	savingsAccount.Operation = 1
	fmt.Println(savingsAccount.Deposit(1000.))
	savingsAccount.Deposit(1000.)
	fmt.Println(savingsAccount.GetBalance())

	fmt.Println(payInvoices(account, 200))
	fmt.Println(payInvoices(account2, 2000))
	fmt.Println(payInvoices(savingsAccount, 500))
}
