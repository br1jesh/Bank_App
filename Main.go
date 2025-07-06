package main

import (
	"Banking_App/Account"
	"Banking_App/Bank"
	"Banking_App/Customer"
)

func main() {
	admin := Customer.NewAdmin("Brijesh", "Mavani")
	hdfcBank := Bank.NewBank(admin, "HDFC Bank", "HDFC")


	user1 := Customer.NewUser(admin, "Amit", "Shah")
	user2 := Customer.NewUser(admin, "ajay", "shah")

	Account.NewAccount(user1, hdfcBank.BankId)
	Account.NewAccount(user2, hdfcBank.BankId)

	hdfcBank.GetAllPassbook(admin)
	hdfcBank.GetAllPassbook(user1)
}
