package main

import (
	"Banking_App/Account"
	"Banking_App/Bank"
	"Banking_App/Customer"
	"fmt"
)

func main() {
	admin := Customer.NewAdmin("Brijesh", "Mavani")
	user := Customer.NewUser(admin, "Jaydeep", "Patel")

	bank := Bank.NewBank(admin, "Yes Bank", "YB")

	account1 := Account.NewAccount(user, bank)
	account2 := Account.NewAccount(user, bank)

	fmt.Println(account1)
	fmt.Println(account2)
	printCustomerAccounts(user)
}

func printCustomerAccounts(user *Customer.Customer) {
	for _, acc := range Account.Accounts {
		if acc.CustomerId == user.CustomerId {
			println("Account:", acc.AccountNo, "Balance:", acc.Balance)
		}
	}
	println("Total balance for", user.FirstName, user.LastName, ":", user.TotalBalance)
}
