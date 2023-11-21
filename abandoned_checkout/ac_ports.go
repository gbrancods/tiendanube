// The abandoned checkout is created when the customer reaches checkout's seconds step, but for some reason it does not finish the process.
//
// Note: The application must have read_orders or write_orders permission (depends on whether they just want to do GET or POST).
//
// Note: A cart can be created up to 6 hours after the purchase has been abandoned.
//
// Note: The carts are accessible for 30 days after being created.
//
// Note: A purchase is only converted into an abandoned cart if the customer has arrived to checkout's second step.
package abandonedcheckout

import (
	"encoding/json"
)

// List all Abandoned Checkouts.
func GetAll() (r []Checkout, err error) {
	br, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Returns a specific abandoned checkout.
func GetByID(checkoutID int) (r Checkout, err error) {
	br, err := getByID(checkoutID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Add a discount coupon to the cart.
func (c Coupon) Create(checkoutID int) (r Checkout, err error) {
	j, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := createCoupon(j, checkoutID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}
