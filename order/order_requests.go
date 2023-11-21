package order

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "orders",
	}.Do()
	return
}

func getByID(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%s", id),
	}.Do()
	return
}

func create(o []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "orders",
		Body:    bytes.NewReader(o),
	}.Do()
	return
}

func update(o []byte, id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("orders/%s", id),
		Body:    bytes.NewReader(o),
	}.Do()
	return
}

// Pay

func close(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%s/close", id),
	}.Do()
	return
}

func open(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%s/open", id),
	}.Do()
	return
}

func pack(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%s/pack", id),
	}.Do()
	return
}

func fulfill(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%s/fulfill", id),
	}.Do()
	return
}

func cancel(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%s/cancel", id),
	}.Do()
	return
}
