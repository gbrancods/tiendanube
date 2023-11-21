package abandonedcheckout

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "checkouts",
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("checkouts/%d", id),
	}.Do()
	return
}

func createCoupon(c []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("checkouts/%d/coupons", id),
		Body:    bytes.NewReader(c),
	}.Do()
	return
}
