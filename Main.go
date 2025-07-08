package main

import (
	"Banking_App/Account"
	"Banking_App/Customer"
	"Banking_App/Ledger"
	"fmt"
)

func main() {
	admin := Customer.NewAdmin("Ajay", "Shah")
	user1 := Customer.NewUser(admin, "Brijesh", "Mavani")
	user2 := Customer.NewUser(admin, "Jay", "Shah")

	b1 := admin.NewBank("State Bank India")
	b2 := admin.NewBank("Housing Development Finance Corporation")
	b3 := admin.NewBank("Yes Bank")

	
	acc1 := user1.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, b1.BankID, "SBI")
	acc2 := user1.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, b2.BankID, "HDFC")
	acc3 := user2.NewAccount(user2.CustomerId, user2.FirstName, user2.LastName, b1.BankID, "SBI")
	acc4 := user2.NewAccount(user2.CustomerId, user2.FirstName, user2.LastName, b3.BankID, "YB")

	Account.DepositToAccount(acc1.AccountNo, 2000)
	Account.DepositToAccount(acc3.AccountNo, 1500)
	Account.DepositToAccount(acc4.AccountNo, 100)

	if err := acc2.Withdraw(500); err != nil {
		fmt.Println("Withdraw error:", err)
	}

	Account.TransferBetweenAccounts(user1.CustomerId, acc1.AccountNo, acc2.AccountNo, 1000)
	Account.TransferBetweenAccounts(user1.CustomerId, acc1.AccountNo, acc2.AccountNo, 100)
	Account.TransferBetweenAccounts(user2.CustomerId, acc1.AccountNo, acc4.AccountNo, 100)

	
	acc1.PrintPassbook()

	fmt.Println("\n------------  ---")
	user2.ViewBalances()

	fmt.Println("\n------------------------")
	b1.PrintAllPassbooks()

	fmt.Println("\n--- Deleting User1---")
	Customer.DeleteCustomer(admin, user1.CustomerId)

	fmt.Println("\n--- Ledger Balances ---")
	Ledger.PrintBankBalances()
}
