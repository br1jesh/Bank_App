package Account

import (
	ledger "Banking_App/Ledger"
	transaction "Banking_App/Transaction"
	"fmt"
	"time"
)

type Account struct {
	AccountNo  int
	BankId     int
	BankName   string
	CustomerId int
	Balance    float32
	FullName   string
	IsActive   bool
	Passbook   []*transaction.Transaction
}

var (
	account_Id    = 1001
	transactionId = 1
	allAccounts   []*Account
)

func NewAccount(customerId int, firstName, lastName string, bankId int, bankName string) *Account {

	a := &Account{
		AccountNo:  account_Id,
		BankId:     bankId,
		BankName:   bankName,
		CustomerId: customerId,
		Balance:    1000,
		IsActive:   true,
		FullName:   firstName + " " + lastName,
	}
	account_Id++
	allAccounts = append(allAccounts, a)

	a.AddTransaction("Deposit", 1000, "Account opening")
	fmt.Println("Account created:", a.AccountNo, "Name:", a.FullName, "Initial Balance: Rs.", a.Balance)
	return a
}

func GetAllAccounts() []*Account {
	return allAccounts
}

func GetAccountByNo(accountNo int) *Account {
	for _, acc := range allAccounts {
		if acc.AccountNo == accountNo {
			return acc
		}
	}
	return nil
}

func (a *Account) Deposit(accountNo int, amount float32) *Account {

	acc := GetAccountByNo(accountNo)
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive")
		return acc
	}
	if !acc.IsActive {
		fmt.Println("Deposit failed: Account", accountNo, "is inactive.")
		return nil
	}
	acc.Balance += amount
	acc.AddTransaction("Deposit", amount, "Deposited")
	fmt.Println("Deposited Rs.", amount, "to Account", acc.AccountNo, "New Balance: Rs.", acc.Balance)
	return nil
}

func (a *Account) Withdraw(accountNo int, amount float32) *Account {

	if !a.IsActive {
		fmt.Println("Withdrawal failed: Account", a.AccountNo, "is inactive.")
		return nil
	}
	if amount <= 0 {
		return nil
	}
	if amount > a.Balance {
		return nil
	}

	a.Balance -= amount
	a.AddTransaction("Withdraw", amount, "Withdrawn")
	fmt.Println("Withdrawn Rs.", amount, "from Account", a.AccountNo, "New Balance: Rs.", a.Balance)
	return nil
}

func (a *Account) TransferBetweenAccountsSelf(toAcc *Account, amount float32) {

	if !a.IsActive {
		fmt.Println("Transfer failed: From Account", a.AccountNo, "is inactive.")
		return
	}

	if amount <= 0 {
		fmt.Println("Transfer amount must be positive")
		return
	}

	if a.AccountNo == toAcc.AccountNo {
		fmt.Println("Cannot transfer to the same account.")
		return
	}

	if a.Balance < amount {
		fmt.Println("Insufficient balance in from account.")
		return
	}

	a.Balance -= amount
	a.AddTransaction("Transfer Out", amount, fmt.Sprintf("Transferred to Account %d", toAcc.AccountNo))

	toAcc.Balance += amount
	toAcc.AddTransaction("Transfer In", amount, fmt.Sprintf("Received from Account %d", a.AccountNo))

	ledger.AddTransferEntry(
		a.AccountNo, a.BankName,
		toAcc.AccountNo, toAcc.BankName,
		amount,
	)

	fmt.Println("Transferred Rs.", amount, " from Account ", a.AccountNo, " to Account ", toAcc.AccountNo)
}

func TransferBetweenAccounts(fromAccNo, toAccNo int, amount float32) {

	fromAcc := GetAccountByNo(fromAccNo)
	toAcc := GetAccountByNo(toAccNo)

	if fromAcc == nil || toAcc == nil {
		fmt.Println("One or both accounts not found.")
		return
	}

	if !fromAcc.IsActive || !toAcc.IsActive {
		fmt.Println("One or both accounts are inactive.")
		return
	}

	if amount <= 0 {
		fmt.Println("Transfer amount must be positive.")
		return
	}

	if fromAcc.Balance < amount {
		fmt.Println("Insufficient balance.")
		return
	}

	fromAcc.Balance -= amount
	fromAcc.AddTransaction("Transfer Out", amount, fmt.Sprintf("To Account %d", toAcc.AccountNo))

	toAcc.Balance += amount
	toAcc.AddTransaction("Transfer In", amount, fmt.Sprintf("From Account %d", fromAcc.AccountNo))

	ledger.AddTransferEntry(
		fromAcc.AccountNo, fromAcc.BankName,
		toAcc.AccountNo, toAcc.BankName,
		amount,
	)

	fmt.Println("Transferred Rs.", amount, "from Account", fromAcc.AccountNo, "to Account", toAcc.AccountNo)
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
	if !a.IsActive {
		fmt.Println("Cannot print passbook: Account", a.AccountNo, "is inactive.")
		return
	}
	fmt.Println("Passbook for Account:", a.AccountNo)
	for _, tx := range a.Passbook {
		fmt.Printf("[%s] %s of Rs.%.2f\n",
			tx.Timestamp.Format("2006-01-02 15:04:05"), tx.Type, tx.Amount)
	}
}

func GetAccountsPaginated(page, size int) []*Account {
	start := (page - 1) * size
	end := start + size

	if start >= len(allAccounts) {
		return []*Account{}
	}

	if end > len(allAccounts) {
		end = len(allAccounts)
	}
	return allAccounts[start:end]
}

func DeleteAccount(accountNo int) error {
	acc := GetAccountByNo(accountNo)
	if acc == nil {
		return fmt.Errorf("account not found")
	}
	if !acc.IsActive {
		return fmt.Errorf("account already inactive")
	}
	acc.IsActive = false
	fmt.Println("Account", accountNo, "deleted .")
	return nil
}
