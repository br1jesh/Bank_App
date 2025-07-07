package Customer

import (
	"Banking_App/Account"
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

func NewUser(creator *Customer, firstName, lastName string) *Customer {
	if !creator.IsRoleAdmin() {
		fmt.Println("Only Admin create new users.")
		return nil
	}
	return newCustomer(firstName, lastName, User)
}

func (c *Customer) IsRoleAdmin() bool {
	return c.Role == Admin
}

func GetCustomerById(id int) *Customer {
	for _, cust := range customers {
		if cust.CustomerId == id {
			return cust
		}
	}
	return nil
}

func GetAllCustomers() []*Customer {
	return customers
}

func DeleteCustomer(admin *Customer, customerId int) {
	defer handleCustomerPanic("DeleteCustomer")
	if !admin.IsRoleAdmin() {
		fmt.Println("Only Admin delete customer.")
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

	for i, u := range customers {
		if u.CustomerId == customerId {
			customers = append(customers[:i], customers[i+1:]...)
			fmt.Println("Deleted customer:", cust.FirstName, cust.LastName)
			return
		}
	}
	fmt.Println("Customer deletion UnSuccesful.")
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
