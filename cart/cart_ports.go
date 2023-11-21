// The Cart resource allows the manipulation of shopping carts generated in a storefront.
// You need the write_orders scope in order to call these endpoints.
package cart

// Remove a line item from a Cart by its line item id.
func DeleteCartItem(cartID, itemID int) (err error) {
	err = deleteItem(cartID, itemID)
	return
}

// Unset a Coupon from a Cart by its coupon id.
func UnsetCouponCart(cartID, couponID int) (err error) {
	err = unsetCoupon(cartID, couponID)
	return
}
