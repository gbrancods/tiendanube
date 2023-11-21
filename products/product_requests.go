package products

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "products",
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/%d", id),
	}.Do()
	return
}

func getBySKU(sku string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/sku/%s", sku),
	}.Do()
	return
}

func create(p []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "products",
		Body:    bytes.NewReader(p),
	}.Do()
	return
}

func update(p []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: "products",
		Body:    bytes.NewReader(p),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: "products",
	}.Do()
	return
}

func stockPrice(sp []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PATCH",
		Destiny: "products",
		Body:    bytes.NewReader(sp),
	}.Do()
	return
}
