package main

import "fmt"

type Account struct {
	name          string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	account1 := Account{
		name:          "Caio Baeta",
		agencyNumber:  1234,
		accountNumber: 5678,
		balance:       1000.00,
	}
	account2 := Account{
		"Caio Baeta",
		1234,
		5678,
		1000.00,
	}
	fmt.Println(account1)
	fmt.Println(account2)
}
