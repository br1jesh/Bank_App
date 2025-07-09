package Bank

import (
	"Banking_App/Account"
	"errors"
	"fmt"
	"strings"
)

func handleBankPanic(context string) {
	if r := recover(); r != nil {
		fmt.Println("Recovered panic from", context, r)
	}
}

var (
	bankId   = 0
	allBanks = make(map[int]*Bank)
)

type Bank struct {
	BankID       int
	FullName     string
	Abbreviation string
	Accounts     []*Account.Account
	IsActive     bool
}

func NewBank(fullName string) *Bank {
	if fullName == "" {
		return nil
	}
	bankId++
	bank := &Bank{
		BankID:       bankId,
		FullName:     fullName,
		Abbreviation: getAbbreviation(fullName),
		IsActive:     true,
	}
	allBanks[bank.BankID] = bank
	fmt.Println("Bank created:", bank.FullName, "(", bank.Abbreviation, ") with ID:", bank.BankID)
	return bank
}

func getAbbreviation(input string) string {
	words := strings.Fields(input)
	var firstLetters []string
	for _, word := range words {
		if len(word) > 0 {
			firstLetters = append(firstLetters, string(word[0]))
		}
	}
	return strings.Join(firstLetters, "")
}

func GetBankById(bankId int) (*Bank, error) {
	defer handleBankPanic("GetBank")
	bank, exists := allBanks[bankId]
	if !exists {
		return nil, errors.New("bank not found")
	}
	return bank, nil
}

func (b *Bank) UpdateBank(fullName string) error {
	defer handleBankPanic("UpdateBank")

	if fullName == "" {
		return errors.New("fullname cannot be empty")
	}

	b.FullName = fullName
	b.Abbreviation = getAbbreviation(fullName)
	fmt.Println("Bank updated:", b.BankID, b.FullName, b.Abbreviation)

	return nil
}

func DeleteBank(bankId int) error {
	defer handleBankPanic("DeleteBank")

	bank, exists := allBanks[bankId]
	if !exists {
		return errors.New("bank not found")
	}
	if len(bank.Accounts) > 0 {
		return errors.New("cannot delete bank with active accounts")
	}

	bank.IsActive = false
	fmt.Println("Bank with ID", bankId, " deleted")
	return nil
}

func (b *Bank) PrintAllPassbooks() {
	defer handleBankPanic("PrintAllPassbook")
	fmt.Println("Bank:", b.FullName, "| ID:", b.BankID)
	if len(b.Accounts) == 0 {
		fmt.Println("No accounts found for this bank.")
		return
	}
	for _, acc := range b.Accounts {
		fmt.Println("Account No:", acc.AccountNo, " Customer ID:", acc.CustomerId, " Balance:", acc.Balance)
	}
}

func GetBanksPaginated(page, size int) []*Bank {
	defer handleBankPanic("GetBanksPaginated")

	var bankList []*Bank
	for _, bank := range allBanks {
		bankList = append(bankList, bank)
	}

	start := (page - 1) * size
	end := start + size

	if start >= len(bankList) {
		return []*Bank{}
	}

	if end > len(bankList) {
		end = len(bankList)
	}

	return bankList[start:end]
}
