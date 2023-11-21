package coupons

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "coupons",
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("coupons/%d", id),
	}.Do()
	return
}

func create(c []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "coupons",
		Body:    bytes.NewReader(c),
	}.Do()
	return
}

func update(c []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("coupons/%d", id),
		Body:    bytes.NewReader(c),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("coupons/%d", id),
	}.Do()
	return
}
