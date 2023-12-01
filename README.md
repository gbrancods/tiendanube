# Tiendanube Go SDK 
### A complete sdk for the Tiendanube *(beta version)*

#### All examples are available in the package [tiendanube-consumer](https://github.com/gbrancods/tiendanube-consumer)
#### Get the package running: `go get github.com/gbrancods/tiendanube`
#### Example of how to use:

```go
// All documentation is imported, hover over imports or functions for more details
package tiendanube

import (
	"os"

	"github.com/gbrancods/tiendanube/auth"
	"github.com/gbrancods/tiendanube/customer"
)

func main() {
	id, _ := strconv.Atoi(os.Getenv("USER_ID"))

	//Starting
	auth.Access{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		TokenType:   "bearer",
		UserID:      id,
		APIurl:      "api.nuvemshop.com.br",
	}.Start()

	customerRunner()
}

func customerRunner() {
	c := CreateCustomer()
	UpdateCustomer(c)
	GetCustomerByID(c.ID)
	GetAllCustomers()
}

func CreateCustomer() (c customer.Customer) {

	c, err := customer.New(
		customer.SetName("Customer Name"),
		customer.SetEmail("example2@mail.com"),
		customer.SetAddress(
			[]customer.Address{
				{
					Address:  "Customer Street",
					City:     "Customer city",
					Country:  "BR",
					Locality: "Customer District",
					Number:   "999",
					Phone:    "+00 00 00000-0000",
					Province: "Customer Province",
					Zipcode:  "Customer Zipcode",
				},
			},
		),
		customer.SetEmailInvite(true),
		customer.SetPassword("CustomerSuperSecretPass"),
		customer.SetExtra(customer.Extra{
			NumberOfChildren: "0",
			Gender:           "male",
		}),
	).Create()

	if err != nil {
		panicErr("Error creating customer", err)
	}

	prettyPrint("Customer created!", c)
	return
}

func UpdateCustomer(c customer.Customer) {

	c.Name = "Updated Name"

	r, err := c.Update(c.ID)
	if err != nil {
		panicErr("Error updating customer", err)
	}

	prettyPrint("Customer updated!", r)
}

func GetCustomerByID(customerID int) {
	
	c, err := customer.GetByID(customerID)

	if err != nil {
		panicErr("Error getting customer", err)
	}

	prettyPrint("Customer:", c)
}

func GetAllCustomers() {
	c, err := customer.GetAll()

	if err != nil {
		panicErr("Error getting all customers", err)
	}

	prettyPrint("All Customers:", c)
}

func prettyPrint(m string, v any) {
	fmt.Println()
	fmt.Println("----------")
	fmt.Println(m)
	fmt.Printf("%+v\n", v)
}

func panicErr(m string, err error) {
	fmt.Println()
	fmt.Println("----------")
	fmt.Println(m)
	panic(err)
}

```
