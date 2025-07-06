package Bank

import (
	"Banking_App/Account"
	"Banking_App/Customer"
	"fmt"
)

type Bank struct {
	BankId       int
	FullName     string
	Abbreviation string
}

var (
	bank_Id = 1
	banks   []*Bank
)

func NewBank(c *Customer.Customer, fullName, abbreviation string) *Bank {
	if !c.IsRoleAdmin() {
		fmt.Println("Only Admin can create a new Bank.")
		return nil
	}
	b := &Bank{
		BankId:       bank_Id,
		FullName:     fullName,
		Abbreviation: abbreviation,
	}
	bank_Id++
	banks = append(banks, b)

	fmt.Println("Bank Created:", fullName, abbreviation, "ID:", b.BankId)
	return b
}

func (b *Bank) GetAllPassbook(c *Customer.Customer) {
	if !c.IsRoleAdmin() {
		fmt.Println("Only Admin can view all user passbooks.")
		return
	}

	fmt.Println("Bank Name:", b.FullName, "| Bank Id:", b.BankId)
	found := false
	for _, acc := range Account.Accounts {
		if acc.BankId == b.BankId {
			found = true
			customer := Customer.GetCustomerById(acc.CustomerId)
			custName := "Unknown"
			if customer != nil {
				custName = customer.FirstName + " " + customer.LastName
			}
			fmt.Println("AccountNo:", acc.AccountNo, "| CustomerID:", acc.CustomerId, "| Name:", custName, "| Balance:", acc.Balance)
		}
	}
	if !found {
		fmt.Println("No accounts found for this User.")
	}
}
