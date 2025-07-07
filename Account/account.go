package Account

import (
	transaction "Banking_App/Transaction"
	"errors"
	"fmt"
	"time"
)

func handleAccountPanic(context string) {
	if r := recover(); r != nil {
		fmt.Println("Recovered panic from", context, r)
	}
}

type Account struct {
	AccountNo  int
	BankId     int
	CustomerId int
	Balance    float32
	FullName string
	Passbook   []*transaction.Transaction
}

var (
	account_Id    = 1
	transactionId = 1
	Accounts      []*Account
)

func NewAccount(customerId int, firstName, lastName string, bankId int) *Account {
	defer handleAccountPanic("NewAccount")

	a := &Account{
		AccountNo:  account_Id,
		BankId:     bankId,
		CustomerId: customerId,
		Balance:    1000,
		FullName: firstName + " " + lastName,
	}
	account_Id++
	Accounts = append(Accounts, a)

	a.AddTransaction("Deposit", 1000, "Account opening")

	fmt.Println("Account created:", a.AccountNo, "Account created:", a.FullName, "Initial Balance: Rs.", a.Balance)
	return a
}

func GetAllAccounts() []*Account {
	return Accounts
}

func (a *Account) Deposit(amount float32) error {
	defer handleAccountPanic("Deposit")
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.Balance += amount
	a.AddTransaction("Deposit", amount, "Deposited")
	fmt.Println("Deposited Rs.", amount, "to Account", a.AccountNo, "New Balance: Rs.", a.Balance)
	return nil
}

func (a *Account) Withdraw(amount float32) error {
	defer handleAccountPanic("Withdraw")
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	a.AddTransaction("Withdraw", amount, "Withdrawn")
	fmt.Println("Withdrawn Rs.", amount, "from Account", a.AccountNo, "New Balance: Rs.", a.Balance)
	return nil
}

func (a *Account) AddTransaction(txType string, amount float32, note string) {
	t := &transaction.Transaction{
		TxID:      transactionId,
		Type:      txType,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	transactionId++
	a.Passbook = append(a.Passbook, t)
}

func (a *Account) PrintPassbook() {
	fmt.Println("Passbook for Account:", a.AccountNo)
	for _, tx := range a.Passbook {
		fmt.Printf("[%s] %s of Rs.%.2f ",tx.Timestamp.Format("2006-01-02 15:04:05"), tx.Type, tx.Amount)
	}
}
