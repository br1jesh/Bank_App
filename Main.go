package main

import (
	"Banking_App/Customer"
	"Banking_App/Ledger"
	"fmt"
)

func main() {
	admin := Customer.NewAdmin("Ajay", "Shah")
	user1 := admin.NewUser("Brijesh", "Mavani")
	user2 := admin.NewUser("Jay", "Shah")

	b1 := admin.NewBank("State Bank India")
	b2 := admin.NewBank("Housing Development Finance Corporation")
	b3 := admin.NewBank("Yes Bank")
	admin.UpdateBank(b1.BankID, "Industrial Credit and Investment Corporation of India")
	fmt.Println(b1)

	acc1 := user1.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, b1.BankID, "SBI")
	acc5 := user1.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, b1.BankID, "SBI")
	acc2 := user1.NewAccount(user1.CustomerId, user1.FirstName, user1.LastName, b2.BankID, "HDFC")
	acc3 := user2.NewAccount(user2.CustomerId, user2.FirstName, user2.LastName, b1.BankID, "SBI")
	acc4 := user2.NewAccount(user2.CustomerId, user2.FirstName, user2.LastName, b3.BankID, "YB")

	user1.DepositToAccount(acc1.AccountNo, 2000)
	user1.DepositToAccount(acc2.AccountNo, 9900)
	user2.DepositToAccount(acc3.AccountNo, 1500)
	user2.DepositToAccount(acc4.AccountNo, 100)
	user1.DepositToAccount(acc5.AccountNo, 9900)
	user1.Withdraw(acc1.AccountNo, 500)

	user1.TransferBetweenAccountsSelf(1001, 1002, 9)
	user1.TransferBetweenAccountsSelf(acc1.AccountNo, acc2.AccountNo, 100)
	user2.TransferBetweenAccountsSelf(acc3.AccountNo, acc4.AccountNo, 100)
	user1.TransferBetweenAccounts(1001, acc3.AccountNo, 499)

	acc1.PrintPassbook()

	fmt.Println("\n------------  ---")
	user2.ViewBalances()

	fmt.Println("\n------------------------")
	b1.PrintAllPassbooks()

	Ledger.PrintDetailedSummary()
}
