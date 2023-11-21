// A draft order enables create orders through outside channels.
package draftorder

import (
	"encoding/json"
)

func defaultDraftOrder() DraftOrderOpts {
	return DraftOrderOpts{}
}

// Customer's name.
func SetContactName(contactName string) OptDrO {
	return func(d *DraftOrderOpts) {
		d.ContactName = contactName
	}
}

// Customer's lastname.
func SetContactLastName(contactLastName string) OptDrO {
	return func(d *DraftOrderOpts) {
		d.ContactLastName = contactLastName
	}
}

// Customer contact email.
func SetContactEmail(contactEmail string) OptDrO {
	return func(d *DraftOrderOpts) {
		d.ContactEmail = contactEmail
	}
}

// Customer contact phone.
func SetContactPhone(contactPhone string) OptDrO {
	return func(d *DraftOrderOpts) {
		d.ContactPhone = contactPhone
	}
}

// Customer document
func SetIdentification(contactIdentification string) OptDrO {
	return func(d *DraftOrderOpts) {
		d.ContactIdentification = contactIdentification
	}
}

// Draft Order's payment status. Possible values are "unpaid", "pending_confirmation" and "paid".
func SetPaymentStatus(paymentStatus string) OptDrO {
	return func(d *DraftOrderOpts) {
		d.PaymentStatus = paymentStatus
	}
}

// TODO - not working
// Sales channel name.
//func SetSaleChannel(saleChannel string) OptDrO {
//	return func(d *DraftOrderOpts) {
//		d.SaleChannel = saleChannel
//	}
//}

// Store owner's note about the draft order.
func SetNote(note string) OptDrO {
	return func(d *DraftOrderOpts) {
		d.Note = note
	}
}

// Draft Order's products list (Product).
func SetProducts(products []Product) OptDrO {
	return func(d *DraftOrderOpts) {
		d.Products = products
	}
}

// Discount amount applied to the draft order.
func SetDiscount(discount string) OptDrO {
	return func(d *DraftOrderOpts) {
		d.Discount = discount
	}
}

// Shipping information (Shipping).
func SetShippingAddress(shippingAddress ShippingAddress) OptDrO {
	return func(d *DraftOrderOpts) {
		d.ShippingAddress = shippingAddress
	}
}

// Initialize a new draft order instance
func New(opts ...OptDrO) *DraftOrder {
	dc := defaultDraftOrder()
	for _, fn := range opts {
		fn(&dc)
	}
	return &DraftOrder{
		DraftOrderOpts: dc,
	}
}

// Create a draft order.
func (d DraftOrder) Create() (r DraftOrder, err error) {
	b, err := json.Marshal(d)
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

// Confirm a draft order and converts it to an order. Return an Order.
func ConfirmCreation(draftOrderID int) (r DraftOrder, err error) {
	b, err := confirmCreation(draftOrderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// List all draft orders.
func GetAll() (r []DraftOrder, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Get a draft order.
func GetByID(draftOrderID int) (r DraftOrder, err error) {
	b, err := getByID(draftOrderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Delete a draft order.
func Delete(draftOrderID int) (err error) {
	err = delete(draftOrderID)
	return
}
