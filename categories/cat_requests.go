package categories

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "categories",
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("categories/%d", id),
	}.Do()
	return
}

func create(c []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "categories",
		Body:    bytes.NewReader(c),
	}.Do()
	return
}

func update(c []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("categories/%d", id),
		Body:    bytes.NewReader(c),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("categories/%d", id),
	}.Do()
	return
}
