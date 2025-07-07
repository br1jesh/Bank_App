package main

import (
	"Banking_App/Account"
	"Banking_App/Bank"
	"Banking_App/Customer"
	"fmt"
)

func main() {
	bank1 := Bank.NewBank("State Bank India")
	bank2 := Bank.NewBank("Housing Development Finance Corporation ")

	fmt.Println("Bank 1:", *bank1)
	fmt.Println("Bank 2:", *bank2)

	admin := Customer.NewAdmin("Jay", "shah")

	user1 := Customer.NewUser(admin, "jain", "Smith")
	user2 := Customer.NewUser(admin, "Brijesh", "Mavani")

	acc1 := Account.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, bank1.BankID)
	acc2 := Account.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, bank2.BankID)
	acc3 := Account.NewAccount(user2.CustomerId, user2.FirstName, user2.LastName, bank1.BankID)

	acc1.Deposit(2000)
	acc1.Withdraw(500)
	acc2.Deposit(1500)
	acc3.Withdraw(100)

	user1.ViewBalances()
	user2.ViewBalances()

	user1.TransferBetweenAccounts(acc1.AccountNo, acc2.AccountNo, 700)

	acc1.PrintPassbook()
	acc2.PrintPassbook()

	fmt.Println("Printing all passbooks of bank1")
	bank1.PrintAllPassbooks()
	Customer.DeleteCustomer(admin, user1.CustomerId)

	fmt.Println("Deleting user2 accounts first...")
	for _, acc := range Account.GetAllAccounts() {
		if acc.CustomerId == user2.CustomerId {
			acc.Withdraw(acc.Balance)
		}
	}
	Customer.DeleteCustomer(admin, user2.CustomerId)

	Bank.DeleteBank(bank1.BankID)

}
