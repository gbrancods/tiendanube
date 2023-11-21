package productvariant

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll(p int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/%d/variants", p),
	}.Do()
	return
}

func getByID(p, v int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("products/%d/variants/%d", p, v),
	}.Do()
	return
}

func create(v []byte, p int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("products/%d/variants", p),
		Body:    bytes.NewReader(v),
	}.Do()
	return
}

func update(v []byte, p int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("products/%d/variants", p),
		Body:    bytes.NewReader(v),
	}.Do()
	return
}

func patch(v []byte, p int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PATCH",
		Destiny: fmt.Sprintf("products/%d/variants", p),
		Body:    bytes.NewReader(v),
	}.Do()
	return
}

func updateById(v []byte, p, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("products/%d/variants/%d", p, id),
		Body:    bytes.NewReader(v),
	}.Do()
	return
}

func delete(p, id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("products/%d/variants/%d", p, id),
	}.Do()
	return
}

func stockVariant(p int, v []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("products/%d/variants/stock", p),
		Body:    bytes.NewReader(v),
	}.Do()
	return
}
