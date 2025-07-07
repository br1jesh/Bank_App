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
	FullName   string
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
		FullName:   firstName + " " + lastName,
	}
	account_Id++
	Accounts = append(Accounts, a)

	a.AddTransaction("Deposit", 1000, "Account opening")

	fmt.Println("Account created:", a.AccountNo, "Name:", a.FullName, "Initial Balance: Rs.", a.Balance)
	return a
}

func GetAllAccounts() []*Account {
	return Accounts
}

func GetAccountByNo(accountNo int) (*Account, error) {
	for _, acc := range Accounts {
		if acc.AccountNo == accountNo {
			return acc, nil
		}
	}
	return nil, errors.New("account not found")
}

func DepositToAccount(accountNo int, amount float32) {
	defer handleAccountPanic("DepositToAccount")
	acc, err := GetAccountByNo(accountNo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive")
		return
	}
	acc.Balance += amount
	acc.AddTransaction("Deposit", amount, "Deposited")
	fmt.Println("Deposited Rs.", amount, "to Account", acc.AccountNo, "New Balance: Rs.", acc.Balance)
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
		fmt.Printf("[%s] %s of Rs.%.2f\n",
			tx.Timestamp.Format("2006-01-02 15:04:05"), tx.Type, tx.Amount)
	}
}

func TransferBetweenAccounts(customerId, fromAccNo, toAccNo int, amount float32) {
	defer handleAccountPanic("TransferBetweenAccounts")

	if amount <= 0 {
		fmt.Println("Transfer amount must be positive")
		return
	}

	var fromAcc, toAcc *Account
	for _, acc := range Accounts {
		if acc.CustomerId == customerId {
			if acc.AccountNo == fromAccNo {
				fromAcc = acc
			}
			if acc.AccountNo == toAccNo {
				toAcc = acc
			}
		}
	}

	if fromAcc == nil ||toAcc == nil  {
		fmt.Println("From account not found.")
		return
	}
	if fromAccNo == toAccNo {
		fmt.Println("Cannot transfer to the same account.")
		return
	}
	if fromAcc.Balance < amount {
		fmt.Println("Insufficient balance in from account.")
		return
	}

	fromAcc.Balance -= amount
	fromAcc.AddTransaction("Transfer Out", amount, fmt.Sprintf("Transferred to Account %d", toAcc.AccountNo))

	toAcc.Balance += amount
	toAcc.AddTransaction("Transfer In", amount, fmt.Sprintf("Received from Account %d", fromAcc.AccountNo))

	fmt.Println("Transferred Rs.", amount, "from Account", fromAcc.AccountNo, "to Account", toAcc.AccountNo)
}


func GetAccountsPaginated(page, size int) []*Account {
	start := (page - 1) * size
	end := start + size

	if start >= len(Accounts) {
		return []*Account{}
	}

	if end > len(Accounts) {
		end = len(Accounts)
	}
	return Accounts[start:end]
}