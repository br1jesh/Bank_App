package Customer

import (
	"Banking_App/Account"
	"Banking_App/Bank"
	"fmt"
)

func handleCustomerPanic(context string) {
	if r := recover(); r != nil {
		fmt.Println("Recovered panic from", context, r)
	}
}

type Role int

const (
	User Role = iota
	Admin
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case User:
		return "User"
	default:
		return "Unknown"
	}
}

type Customer struct {
	CustomerId int
	FirstName  string
	LastName   string
	Role       Role
	IsActive   bool
	Accounts   []*Account.Account
}

var (
	customerIdCounter = 1
	customers         []*Customer
)

func newCustomer(firstName, lastName string, role Role) *Customer {
	defer handleCustomerPanic("newCustomer")

	if firstName == "" {
		panic("First name cannot be empty.")
	}
	if lastName == "" {
		panic("Last name cannot be empty.")
	}

	c := &Customer{
		CustomerId: customerIdCounter,
		FirstName:  firstName,
		LastName:   lastName,
		Role:       role,
		IsActive:   true,
	}
	customerIdCounter++
	customers = append(customers, c)

	fmt.Println("Created:", role.String(), "-", firstName, lastName, "| ID:", c.CustomerId)
	return c
}

func NewAdmin(firstName, lastName string) *Customer {
	return newCustomer(firstName, lastName, Admin)
}

func (c *Customer) NewUser(firstName, lastName string) *Customer {
	if !c.IsRoleAdmin() {
		panic("Only Admin can create new users.")
	}
	return newCustomer(firstName, lastName, User)
}

func (c *Customer) IsRoleAdmin() bool {
	return c.Role == Admin
}
func (c *Customer) IsRoleUser() bool {
	return c.Role == User
}

func GetCustomerById(id int) *Customer {
	for _, cust := range customers {
		if cust.CustomerId == id {
			return cust
		}
	}
	return nil
}

func (c *Customer) ViewMyAccounts() {
	defer handleCustomerPanic("ViewMyAccounts")
	fmt.Println("Accounts for Customer:", c.FirstName, c.LastName, "| ID:", c.CustomerId)
	hasAccounts := false
	for _, acc := range Account.GetAllAccounts() {
		if acc.CustomerId == c.CustomerId {
			hasAccounts = true
			fmt.Printf("  AccountNo: %d | Balance: Rs. %.2f | IsActive: %t\n",
				acc.AccountNo, acc.Balance, acc.IsActive)
		}
	}
	if !hasAccounts {
		fmt.Println("  No accounts found for this customer.")
	}
}

func (c *Customer) ViewAccountById(accountNo int) {
	defer handleCustomerPanic("ViewAccountById")
	acc := Account.GetAccountByNo(accountNo)
	if acc == nil {
		fmt.Println("Account not found.")
		return
	}
	if acc.CustomerId != c.CustomerId {
		fmt.Println("Unauthorized: This account does not belong to you.")
		return
	}
	fmt.Printf("AccountNo: %d | Balance: Rs. %.2f | IsActive: %t\n",
		acc.AccountNo, acc.Balance, acc.IsActive)
}

func (c *Customer) GetTotalBalance() float32 {
	defer handleCustomerPanic("GetTotalBalance")
	var total float32
	for _, acc := range Account.GetAllAccounts() {
		if acc.CustomerId == c.CustomerId {
			total += acc.Balance
		}
	}
	return total
}

func (c *Customer) ViewBalances() {
	defer handleCustomerPanic("ViewBalances")
	fmt.Println("Customer: ", c.FirstName, c.LastName, "ID: ", c.CustomerId)
	fmt.Println("Accounts & Balances:")
	hasAccounts := false
	for _, acc := range Account.GetAllAccounts() {
		if acc.CustomerId == c.CustomerId {
			hasAccounts = true
			fmt.Println("  AccountNo:", acc.AccountNo, " Balance: Rs.", acc.Balance)
		}
	}
	if !hasAccounts {
		fmt.Println("  No accounts found.")
	}
	fmt.Println("Total Balanceall accounts: Rs", c.GetTotalBalance())
}

func (c *Customer) UpdateCustomer(param string, value interface{}) {
	defer handleCustomerPanic("UpdateCustomer")

	if !c.IsActive {
		panic("Customer is inactive; update skipped.")
	}
	if !c.IsRoleAdmin() {
		panic("Only Admin can update customer.")
	}

	switch param {
	case "FName":
		c.updateCustomerFirstName(value)
	case "LName":
		c.updateCustomerLastName(value)
	case "IsActive":
		c.updateCustomerIsActive(value)
	default:
		panic(fmt.Sprintf("Unknown parameter: %s", param))
	}
}

func (c *Customer) updateCustomerFirstName(value interface{}) {
	if str, ok := value.(string); ok && str != "" {
		c.FirstName = str
		fmt.Println("First name updated successfully.")
	} else {
		panic("updateFirstName: invalid string value")
	}
}

func (c *Customer) updateCustomerLastName(value interface{}) {
	if str, ok := value.(string); ok && str != "" {
		c.LastName = str
		fmt.Println("Last name updated successfully.")
	} else {
		panic("updateLastName: invalid string value")
	}
}

func (c *Customer) updateCustomerIsActive(value interface{}) {
	if status, ok := value.(bool); ok {
		c.IsActive = status
		fmt.Println("IsActive status changed to:", status)
	} else {
		panic("updateIsActiveStatus: invalid bool value")
	}
}

func DeleteCustomer(admin *Customer, customerId int) {
	if !admin.IsRoleAdmin() {
		panic("Only Admin can delete customer.")
	}

	cust := GetCustomerById(customerId)
	if cust == nil {
		panic("Customer not found.")
	}

	for _, acc := range Account.GetAllAccounts() {
		if acc.CustomerId == customerId {
			panic("Cannot delete. Customer has active accounts.")
		}
	}

	cust.IsActive = false
	fmt.Println("Customer:", cust.FirstName, cust.LastName, "soft deleted (IsActive set to false).")
}

func GetCustomerPaginated(page, size int) []*Customer {
	start := (page - 1) * size
	end := start + size

	if start >= len(customers) {
		return []*Customer{}
	}

	if end > len(customers) {
		end = len(customers)
	}
	return customers[start:end]
}

// ---------------------------------Bank
func (c *Customer) NewBank(fullName string) *Bank.Bank {
	if !c.IsRoleAdmin() {
		panic("Only Admin can create a new bank.")
	}
	return Bank.NewBank(fullName)
}

func (c *Customer) ViewBankByID(bankId int) (*Bank.Bank, error) {
	if !c.IsRoleAdmin() {
		panic("Only Admin can view bank details.")
	}
	return Bank.GetBankById(bankId)
}

func (c *Customer) ViewAllBanks(page, size int) ([]*Bank.Bank, error) {
	if !c.IsRoleAdmin() {
		panic("Only Admin can view all banks.")
	}
	return Bank.GetBanksPaginated(page, size), nil
}

func (c *Customer) UpdateBank(bankId int, fullName string) error {
	if !c.IsRoleAdmin() {
		panic("Only Admin can update a bank.")
	}

	bank, err := Bank.GetBankById(bankId)
	if err != nil {
		return err
	}
	return bank.UpdateBank(fullName)
}

func (c *Customer) DeleteBank(bankId int) error {
	if !c.IsRoleAdmin() {
		panic("Only Admin can delete a bank.")
	}
	return Bank.DeleteBank(bankId)
}

// -------------------------------------Account
func (c *Customer) NewAccount(customerId int, firstName, lastName string, bankId int, bankName string) *Account.Account {
	if !c.IsRoleUser() {
		panic("Only Customer can create a new account.")
	}
	acc := Account.NewAccount(customerId, firstName, lastName, bankId, bankName)
	c.Accounts = append(c.Accounts, acc)
	return acc
}

func (c *Customer) DepositToAccount(accountNo int, amount float32) *Account.Account {
	if !c.IsRoleUser() {
		panic("Only Customer can deposit.")
	}
	acc := Account.GetAccountByNo(accountNo)
	if acc == nil {
		panic("Account not found.")
	}
	acc.Deposit(accountNo, amount)
	return acc
}

func (c *Customer) Withdraw(accountNo int, amount float32) *Account.Account {
	if !c.IsRoleUser() {
		panic("Only Customer can withdraw.")
	}
	acc := Account.GetAccountByNo(accountNo)
	if acc == nil {
		panic("Account not found.")
	}
	acc.Withdraw(accountNo, amount)
	return acc
}

func (c *Customer) TransferBetweenAccountsSelf(fromAccNo, toAccNo int, amount float32) {
	defer handleCustomerPanic("TransferBetweenAccountsSelf")

	if !c.IsRoleUser() {
		panic("Only Customer can transfer between their accounts.")
	}

	var fromAcc, toAcc *Account.Account
	for _, acc := range c.Accounts {
		if acc.AccountNo == fromAccNo {
			fromAcc = acc
		}
		if acc.AccountNo == toAccNo {
			toAcc = acc
		}
	}

	if fromAcc == nil || toAcc == nil {
		panic("Account(s) not found for this customer.")
	}

	fromAcc.TransferBetweenAccountsSelf(toAcc, amount)
}

func (c *Customer) DeleteAccount(accountNo int) error {
	if !c.IsRoleUser() {
		panic("Only Customer can delete their accounts.")
	}
	return Account.DeleteAccount(accountNo)
}

func (c *Customer) TransferBetweenAccounts(fromAccNo, toAccNo int, amount float32) {
	defer handleCustomerPanic("TransferBetweenAccounts")

	if !c.IsRoleUser() {
		panic("Only Customer can transfer between accounts.")
	}

	fromAcc := Account.GetAccountByNo(fromAccNo)
	if fromAcc == nil {
		panic("From Account not found.")
	}
	if fromAcc.CustomerId != c.CustomerId {
		panic("Unauthorized: You do not own the from-account.")
	}

	Account.TransferBetweenAccounts(fromAccNo, toAccNo, amount)
}
