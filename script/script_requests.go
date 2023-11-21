package script

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "scripts",
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("scripts/%d", id),
	}.Do()
	return
}

func create(s []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "scripts",
		Body:    bytes.NewReader(s),
	}.Do()
	return
}

func update(s []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("scripts/%d", id),
		Body:    bytes.NewReader(s),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("scripts/%d", id),
	}.Do()
	return
}
