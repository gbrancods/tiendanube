// A Customer of the store. Customer accounts store contact
// information for the customer, saving logged-in customers
// the trouble of having to provide it at every checkout.
package customer

import (
	"encoding/json"
)

func defaultCustomer() CustomerOpts {
	return CustomerOpts{}
}

// Name of the Customer
func SetName(name string) OptCus {
	return func(c *CustomerOpts) {
		c.Name = name
	}
}

// E-mail of the Customer
func SetEmail(email string) OptCus {
	return func(c *CustomerOpts) {
		c.Email = email
	}
}

// List of shipping addresses for the Customer
func SetAddress(addresses []Address) OptCus {
	return func(c *CustomerOpts) {
		c.Addresses = addresses
	}
}

// Send an email to notify the customer of their registration
func SetEmailInvite(boolean bool) OptCus {
	return func(c *CustomerOpts) {
		c.SendEmailInvite = boolean
	}
}

// User's password
func SetPassword(password string) OptCus {
	return func(c *CustomerOpts) {
		c.Password = password
	}
}

// Initialize a customer instance
func New(opts ...OptCus) *Customer {

	dc := defaultCustomer()

	for _, fn := range opts {
		fn(&dc)
	}

	return &Customer{
		CustomerOpts: dc,
	}
}

// Create a new Customer
func (c Customer) Create() (r Customer, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := create(b)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Modify an existing Customer
func (c Customer) Update(couponID int) (r Customer, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := update(b, couponID)
	json.Unmarshal(br, &r)
	return
}

// Receive a list of all Customers.
func GetAll() (cl []Customer, err error) {
	br, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &cl)
	return
}

// Receive a single Customer
func GetByID(customerID int) (r Customer, err error) {
	br, err := getByID(customerID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}
