package pvcustomfields

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func create(v []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "products/variants/custom-fields",
		Body:    bytes.NewReader(v),
	}.Do()
	return
}

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "products/variants/custom-fields",
	}.Do()
	return
}

func update(v []byte, id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("products/variants/custom-fields/%s", id),
		Body:    bytes.NewReader(v),
	}.Do()
	return
}

func getByVariant(v int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/variants/%d/custom-fields", v),
	}.Do()
	return
}

func getOwners(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/variants/custom-fields/%s/owners", id),
	}.Do()
	return
}

func updateValue(v []byte, id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("products/variants/%s/custom-fields/values", id),
		Body:    bytes.NewReader(v),
	}.Do()
	return
}

func delete(id string) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("products/variants/custom-fields/%s", id),
	}.Do()
	return
}

func getByID(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/variants/custom-fields/%s", id),
	}.Do()
	return
}
