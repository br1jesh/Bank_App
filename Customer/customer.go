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
		fmt.Println("First name cannot be empty.")
		return nil
	}
	if lastName == "" {
		fmt.Println(" last name cannot be empty.")
		return nil
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
		fmt.Println("Only Admin create new users.")
		return nil
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
	if !c.IsActive {
		fmt.Println("Contact is inactive; update skipped.")
		return
	}
	if !c.IsRoleAdmin() {
		fmt.Println("Only admin can create a new bankaccount.")
		return
	}
	switch param {
	case "FName":
		c.updateCustomerFirstName(value)
	case "LName":
		c.updateCustomerLastName(value)
	case "IsActive":
		c.updateCustomerIsActive(value)
	default:
		fmt.Println("Unknown parameter:", param)
	}
}

func (c *Customer) updateCustomerFirstName(value interface{}) {
	if str, ok := value.(string); ok && str != "" {
		c.FirstName = str
		fmt.Println("First name updated successfully.")
	} else {
		fmt.Println("updateFirstName: invalid string")
	}
}

func (c *Customer) updateCustomerLastName(value interface{}) {
	if str, ok := value.(string); ok && str != "" {
		c.LastName = str
		fmt.Println("Last name updated successfully.")
	} else {
		fmt.Println("updateLastName: invalid string")
	}
}

func (c *Customer) updateCustomerIsActive(value interface{}) {
	if status, ok := value.(bool); ok {
		c.IsActive = status
		fmt.Println("IsActive status changed to:", status)
	} else {
		fmt.Println("updateIsActiveStatus: invalid bool")
	}
}

func DeleteCustomer(admin *Customer, customerId int) {
	defer handleCustomerPanic("DeleteCustomer")
	if !admin.IsRoleAdmin() {
		fmt.Println("Only Admin can delete customer.")
		return
	}
	cust := GetCustomerById(customerId)
	if cust == nil {
		fmt.Println("Customer not found")
		return
	}

	for _, acc := range Account.GetAllAccounts() {
		if acc.CustomerId == customerId {
			fmt.Println("Cannot delete. Customer has active accounts.")
			return
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
		fmt.Println("Only Admin can create a new bank.")
		return nil
	}
	return Bank.NewBank(fullName)
}

func (c *Customer) ViewBankByID(bankId int) (*Bank.Bank, error) {
	if !c.IsRoleAdmin() {
		return nil, fmt.Errorf("only admin can view bank details")
	}
	return Bank.GetBankById(bankId)
}

func (c *Customer) ViewAllBanks(page, size int) ([]*Bank.Bank, error) {
	if !c.IsRoleAdmin() {
		return nil, fmt.Errorf("only admin can view all banks")
	}
	return Bank.GetBanksPaginated(page, size), nil
}

func (c *Customer) UpdateBank(bankId int, fullName string) error {
	if !c.IsRoleAdmin() {
		return fmt.Errorf("only admin can update a bank")
	}

	bank, err := Bank.GetBankById(bankId)
	if err != nil {
		return err
	}

	return bank.UpdateBank(fullName)
}

func (c *Customer) DeleteBank(bankId int) error {
	if !c.IsRoleAdmin() {
		return fmt.Errorf("only admin can delete a bank")
	}
	return Bank.DeleteBank(bankId)
}




// -------------------------------------Account
func (c *Customer) NewAccount(customerId int, firstName, lastName string, bankId int, bankName string) *Account.Account {
	if !c.IsRoleUser() {
		fmt.Println("Only Customer can create a new bankaccount.")
		return nil
	}
	acc := Account.NewAccount(customerId, firstName, lastName, bankId, bankName)
	c.Accounts = append(c.Accounts, acc) 
	return acc
}


func (c *Customer) DepositToAccount(accountNo int, amount float32) *Account.Account {
	if !c.IsRoleUser() {
		fmt.Println("Only Customer can deposit.")
		return nil
	}
	acc := Account.GetAccountByNo(accountNo)
	if acc == nil {
		fmt.Println("Account not found.")
		return nil
	}
	acc.Deposit(accountNo, amount)
	return acc
}

func (c *Customer) Withdraw(accountNo int, amount float32) *Account.Account {
	if !c.IsRoleUser() {
		fmt.Println("Only Customer can deposit.")
		return nil
	}
	acc := Account.GetAccountByNo(accountNo)
	if acc == nil {
		fmt.Println("Account not found.")
		return nil
	}
	acc.Withdraw(accountNo, amount)
	return acc
}

func (c *Customer) TransferBetweenAccounts(fromAccNo, toAccNo int, amount float32) {
	defer handleCustomerPanic("TransferBetweenAccounts ")

	if !c.IsRoleUser() {
		fmt.Println("Only Customer can transfer between their accounts.")
		return
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
		fmt.Println("Account(s) not found for this customer.")
		return
	}
	fromAcc.TransferBetweenAccounts(toAcc, amount)
}

func (c *Customer) DeleteAccount(accountNo int) error {
	if !c.IsRoleUser() {
		fmt.Println("Only customers can delete their accounts.")
		return fmt.Errorf("only customers can delete their accounts")
	}
	return Account.DeleteAccount(accountNo)
}
