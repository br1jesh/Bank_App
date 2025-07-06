package Customer

import (
	"fmt"
)

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
	CustomerId   int
	FirstName    string
	LastName     string
	Role         Role
	TotalBalance int
	IsActive     bool
}

var (
	customer_Id = 1
	customers   []*Customer
)

func newCustomer(firstName, lastName string, role Role) *Customer {
	if firstName == "" {
		return nil
	}
	if lastName == "" {
		return nil
	}
	c := &Customer{
		CustomerId:   customer_Id,
		FirstName:    firstName,
		LastName:     lastName,
		Role:         role,
		TotalBalance: 0,
		IsActive:     true,
	}
	customer_Id++
	customers = append(customers, c)
	fmt.Println("Created User:", firstName, lastName, "Role:", role.String(), "ID:", c.CustomerId)
	return c
}

func NewAdmin(firstName, lastName string) *Customer {
	return newCustomer(firstName, lastName, Admin)
}

func NewUser(creator *Customer, firstName, lastName string) *Customer {
	if !creator.IsRoleAdmin() {
		fmt.Println("Only Admin can create new users.")
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
