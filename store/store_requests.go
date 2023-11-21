package store

import "github.com/gbrancods/tiendanube/requests"

func getInfo() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "store",
	}.Do()
	return
}
