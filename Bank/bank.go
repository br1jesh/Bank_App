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

func (b *Bank) UpdateBank(param string, value interface{}) error {
	defer handleBankPanic("UpdateBank")
	if param == "" {
		return errors.New("parameter cannot be empty")
	}
	switch param {
	case "FullName":
		strVal, ok := value.(string)
		if !ok || strVal == "" {
			return errors.New("invalid fullname")
		}
		b.FullName = strVal
		b.Abbreviation = getAbbreviation(strVal)
		fmt.Println("Bank fullname updated :", b.FullName, b.Abbreviation)
	default:
		return errors.New("invalid parameter for update")
	}
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
	delete(allBanks, bankId)
	fmt.Println("Bank with ID", bankId, "deleted successfully")
	return nil
}

func (b *Bank) PrintAllPassbooks() {
	defer handleBankPanic("PrintAllPassbooks")
	fmt.Println("Bank:", b.FullName, "| ID:", b.BankID)
	if len(b.Accounts) == 0 {
		fmt.Println("No accounts found for this bank.")
		return
	}
	for _, acc := range b.Accounts {
		fmt.Println("Account No:", acc.AccountNo, " Customer ID:", acc.CustomerId, " Balance:", acc.Balance)
	}
}
