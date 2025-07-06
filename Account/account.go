package Account

import (
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

func NewAccount(c *Customer.Customer, bankId int) *Account {
	a := &Account{
		AccountNo:  account_Id,
		BankId:     bankId,
		CustomerId: c.CustomerId,
		Balance:    1000,
	}
	account_Id++
	Accounts = append(Accounts, a)

	fmt.Println("Account created:", a.AccountNo, "for Customer:", c.FirstName, c.LastName, "with initial balance 1000")
	return a
}

// func(a *Account) DepositeMoney (Money int){
   
// }


