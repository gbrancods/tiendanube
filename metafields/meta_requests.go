package metafields

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getByOwnerResource(o string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("metafields/%s", o),
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("metafields/%d", id),
	}.Do()
	return
}

func create(m []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "coupons",
		Body:    bytes.NewReader(m),
	}.Do()
	return
}

func update(m []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("metafields/%d", id),
		Body:    bytes.NewReader(m),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("metafields/%d", id),
	}.Do()
	return
}
