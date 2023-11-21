package transaction

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func create(t []byte, id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%s/transactions", id),
		Body:    bytes.NewReader(t),
	}.Do()
	return
}

func createEvent(tr []byte, o, t string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%s/transactions/%s/events", o, t),
		Body:    bytes.NewReader(tr),
	}.Do()
	return
}

func getAll(o string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%s/transactions", o),
	}.Do()
	return
}

func getByID(o, t string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%s/transactions/%s", o, t),
	}.Do()
	return
}
