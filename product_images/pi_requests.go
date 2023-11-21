package productimages

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

// TODO DOC
func getByProduct(p int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/%d/images", p),
	}.Do()
	return
}

func getByID(p, i int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/%d/images/%d", p, i),
	}.Do()
	return
}

func create(i []byte, p int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("products/%d/images", p),
		Body:    bytes.NewReader(i),
	}.Do()
	return
}

func update(img []byte, p, i int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("products/%d/images/%d", p, i),
		Body:    bytes.NewReader(img),
	}.Do()
	return
}

func delete(p, i int) (err error) {
	_, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("products/%d/images/%d", p, i),
	}.Do()
	return
}
