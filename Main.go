package main

import (
	"Banking_App/Account"
	"Banking_App/Bank"
	"Banking_App/Customer"
	"fmt"
)

func main() {

	b1 := Bank.NewBank("State Bank India")
	b2 := Bank.NewBank(" Housing Development Finance Corporation")

	admin := Customer.NewAdmin("Ajay", "shah")
	user1 := Customer.NewUser(admin, "Brijesh", "Mavani")
	user2 := Customer.NewUser(admin, "Jay", "Shah")

	acc1 := Account.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, b1.BankID)
	acc2 := Account.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, b2.BankID)
	acc3 := Account.NewAccount(user2.CustomerId, user2.FirstName, user2.LastName, b1.BankID)

	Account.DepositToAccount(acc1.AccountNo, 2000)
	Account.DepositToAccount(acc3.AccountNo, 1500)

	if err := acc2.Withdraw(500); err != nil {
		fmt.Println("Withdraw error:", err)
	}
	Account.TransferBetweenAccounts(user1.CustomerId, acc1.AccountNo, acc2.AccountNo, 1000)
	acc1.PrintPassbook()

	fmt.Println("\n------------ ---")
	user2.ViewBalances()

	fmt.Println("\n--- ------------------")
	b1.PrintAllPassbooks()

	fmt.Println("\n----------------- ---")
	Customer.DeleteCustomer(admin, user1.CustomerId)

}
