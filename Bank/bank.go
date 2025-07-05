package Bank

import (
	"Banking_App/Customer"
	"fmt"
)

type Bank struct {
	BankId      int
	FullName    string
	Abbreviation string
}

var (
	bank_Id = 1
	banks       []*Bank
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
