package paymentprovider

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func create(p []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "payment_providers",
		Body:    bytes.NewReader(p),
	}.Do()
	return
}

func update(p []byte, id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("payment_providers/%s", id),
		Body:    bytes.NewReader(p),
	}.Do()
	return
}

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "payment_providers",
	}.Do()
	return
}

func getByID(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("paymentproviders/%s", id),
	}.Do()
	return
}

func delete(id string) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("paymentproviders/%s", id),
	}.Do()
	return
}
