package paymentoptions

import "github.com/gbrancods/tiendanube/requests"

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "payment-options",
	}.Do()

	return
}
