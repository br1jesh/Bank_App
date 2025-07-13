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
var SentDetails = make(map[string]map[string]float32)
var ReceivedDetails = make(map[string]map[string]float32)

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

	if SentDetails[fromBankName] == nil {
		SentDetails[fromBankName] = make(map[string]float32)
	}
	SentDetails[fromBankName][toBankName] -= amount

	if ReceivedDetails[toBankName] == nil {
		ReceivedDetails[toBankName] = make(map[string]float32)
	}
	ReceivedDetails[toBankName][fromBankName] += amount

}


func PrintDetailedSummary() {
	fmt.Println("\n========== Detailed Bank Transfer Summary ==========")
	for bank := range BankBalances {
		fmt.Printf("Bank: %s\n", bank)

	
		if targets, exists := SentDetails[bank]; exists && len(targets) > 0 {
			fmt.Println("  Sent money to:")
			for toBank, amt := range targets {
				fmt.Printf("    %s: %.2f\n", toBank, amt)
			}
		} else {
			fmt.Println("  Sent money to: None")
		}

		if sources, exists := ReceivedDetails[bank]; exists && len(sources) > 0 {
			fmt.Println("  Received money from:")
			for fromBank, amt := range sources {
				fmt.Printf(" %s: +%.2f\n", fromBank, amt)
			}
		} else {
			fmt.Println("  Received money from: None")
		}
		fmt.Println()
	}
}
