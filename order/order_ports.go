// An order is created when a customer completes the checkout process. Orders also can be created through the API.
package order

import (
	"encoding/json"
)

func defaultOrder() OrderOpts {
	return OrderOpts{}
}

// The order currency code (ISO 4217 format). The default is the store currency.
func SetCurrency(currency string) OptOrd {
	return func(c *OrderOpts) {
		c.Currency = currency
	}
}

func SetLanguage(language string) OptOrd {
	return func(c *OrderOpts) {
		c.Language = language
	}
}

func SetGateway(gateway string) OptOrd {
	return func(c *OrderOpts) {
		c.Gateway = gateway
	}
}

func SetPaymentStatus(paymentStatus string) OptOrd {
	return func(c *OrderOpts) {
		c.PaymentStatus = paymentStatus
	}
}

func SetStatus(status string) OptOrd {
	return func(c *OrderOpts) {
		c.Status = status
	}
}

//TODO not found in example
//func SetFulfillment(fulfillment string) OptOrd {
//	return func(c *OrderOpts) {
//		c.F = status
//	}
//}

func SetProducts(products []Product) OptOrd {
	return func(c *OrderOpts) {
		c.Products = products
	}
}

func SetTotal(total string) OptOrd {
	return func(c *OrderOpts) {
		c.Total = total
	}
}

// TODO not found in example
//func SetInventoryBehavier(inventoryBehavier string) OptOrd {
//	return func(c *OrderOpts) {
//		c.Inve = total
//	}
//}

func SetCustomer(customer Customer) OptOrd {
	return func(c *OrderOpts) {
		c.Customer = customer
	}
}

//func SetNote(note ) OptOrd {
//	return func(c *OrderOpts) {
//		c.Note = note
//	}
//}

// Initialize a order instance
func New(opts ...OptOrd) *Order {
	dc := defaultOrder()
	for _, fn := range opts {
		fn(&dc)
	}
	return &Order{
		OrderOpts: dc,
	}
}

// Receive a list of all Orders. Make sure to check out our recommendations on best practices for retrieving orders information in our FAQ section below.
func GetAll() (r []Order, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Order
func GetByID(orderID string) (r Order, err error) {
	b, err := getByID(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Create an Order.
func (o Order) Create() (r Order, err error) {
	b, err := json.Marshal(o)
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

// Change an Order's attributes (just owner_note for now) and/or update an Order's status
func (o Order) Update(orderID string) (r Order, err error) {
	b, err := json.Marshal(o)
	if err != nil {
		return
	}
	br, err := update(b, orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Close an Order
func Close(orderID string) (r Order, err error) {
	br, err := close(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Re-open a closed Order
func Open(orderID string) (r Order, err error) {
	br, err := open(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Pack an Order
func Pack(orderID string) (r Order, err error) {
	br, err := pack(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Fulfill an Order
func Fulfill(orderID string) (r Order, err error) {
	br, err := fulfill(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Cancel an Order
func Calcel(orderID string) (r Order, err error) {
	br, err := cancel(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}
