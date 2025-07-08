package Ledger

import (
	"fmt"
	"time"
)

type LedgerEntry struct {
	FromBank_Id   int
	FromBank_Name string
	ToBank_ID     int
	ToBank_Name   string
	Amount        float32
	Timestamp     time.Time
}

var LedgerBook []LedgerEntry
var BankBalances = make(map[string]float32)

func AddTransferEntry(fromBankId int, fromBankName string, toBankId int, toBankName string, amount float32) {
	entry := LedgerEntry{
		FromBank_Id:   fromBankId,
		FromBank_Name: fromBankName,
		ToBank_ID:     toBankId,
		ToBank_Name:   toBankName,
		Amount:        amount,
		Timestamp:     time.Now(),
	}
	LedgerBook = append(LedgerBook, entry)

	BankBalances[fromBankName] -= amount
	BankBalances[toBankName] += amount

	fmt.Printf("%s sent %.2f to %s\n", fromBankName, amount, toBankName)
	PrintBankBalances()
}

func PrintBankBalances() {
	fmt.Println("Current Balances:")
	for bank, balance := range BankBalances {
		fmt.Printf("%s: %.2f\n", bank, balance)
	}
	fmt.Println()
}
