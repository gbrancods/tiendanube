package cart

import (
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func deleteItem(c, li int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("carts/%d/line-items/%d", c, li),
	}.Do()
	return
}

func unsetCoupon(ca, co int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("carts/%d/coupons/%d", ca, co),
	}.Do()
	return
}
