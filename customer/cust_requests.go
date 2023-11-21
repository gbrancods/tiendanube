package customer

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "customers",
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("customers/%d", id),
	}.Do()
	return
}

func create(c []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "customers",
		Body:    bytes.NewReader(c),
	}.Do()
	return
}

func update(c []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("customers/%d", id),
		Body:    bytes.NewReader(c),
	}.Do()
	return
}
