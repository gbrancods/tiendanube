package webhooks

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "webhooks",
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("webhooks/%d", id),
	}.Do()
	return
}

func create(w []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "webhooks",
		Body:    bytes.NewReader(w),
	}.Do()
	return
}

func update(w []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("webhooks/%d", id),
		Body:    bytes.NewReader(w),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("webhooks/%d", id),
	}.Do()
	return
}
