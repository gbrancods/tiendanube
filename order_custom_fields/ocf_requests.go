package ordercustomfields

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func create(c []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "orders/custom-fields",
		Body:    bytes.NewReader(c),
	}.Do()
	return
}

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "orders/custom-field",
	}.Do()

	return
}

func update(c []byte, id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("orders/custom-fields/%s", id),
		Body:    bytes.NewReader(c),
	}.Do()

	return
}

func getByOrder(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%s/custom-fields", id),
	}.Do()

	return
}

func getByOwner(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/custom-fields/%s/owners", id),
	}.Do()

	return
}

func getByID(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/custom-fields/%s", id),
	}.Do()

	return
}

func updateValue(v []byte, id int) (err error) {
	_, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("orders/%d/custom-fields/values", id),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("orders/custom-fields/%d", id),
	}.Do()
	return
}
