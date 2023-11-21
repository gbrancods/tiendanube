package draftorder

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "draft_orders",
	}.Do()
	fmt.Println(string(r))
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("draft_orders/%d", id),
	}.Do()
	return
}

func create(d []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "draft_orders",
		Body:    bytes.NewReader(d),
	}.Do()
	return
}

func confirmCreation(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("draft_orders/%d/confirm", id),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("draft_orders/%d", id),
	}.Do()
	return
}
