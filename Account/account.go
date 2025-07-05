package Account

import (
	"Banking_App/Bank"
	"Banking_App/Customer"
	"fmt"
)

type Account struct {
	AccountNo  int
	BankId     int
	CustomerId int
	Balance    int
}

var (
	account_Id = 1
	Accounts   []*Account
)

func NewAccount(customer *Customer.Customer, bank *Bank.Bank) *Account {
	a := &Account{
		AccountNo:  account_Id,
		BankId:     bank.BankId,
		CustomerId: customer.CustomerId,
		Balance:    1000,
	}
	account_Id++
	Accounts = append(Accounts, a)
	customer.TotalBalance += 1000

	fmt.Println("Account created:", a.AccountNo, "for Customer:", customer.FirstName, customer.LastName, "with initial balance 1000")
	return a
}
