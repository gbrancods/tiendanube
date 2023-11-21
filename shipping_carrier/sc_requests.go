package shippingcarrier

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

//TODO VERIFICAR DOCS PLURAL

func create(s []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "shipping_carriers",
		Body:    bytes.NewReader(s),
	}.Do()
	return
}

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "shipping_carriers",
	}.Do()
	return
}

func getByID(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("shipping_carriers/%d", id),
	}.Do()
	return
}

func update(s []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("shipping_carriers/%d", id),
		Body:    bytes.NewReader(s),
	}.Do()
	return
}

func delete(id int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("shipping_carriers/%d", id),
	}.Do()
	return
}

func createOption(s []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("shipping_carriers/%d/options", id),
		Body:    bytes.NewReader(s),
	}.Do()
	return
}

func getOptions(id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("shipping_carriers/%d/options", id),
	}.Do()
	return
}

func getOptionByID(id, oid int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("shipping_carriers/%d/options/%d", id, oid),
	}.Do()
	return
}

func updateOption(s []byte, c, o int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("shipping_carriers/%d/options/%d", c, o),
		Body:    bytes.NewReader(s),
	}.Do()
	return
}

func deleteOption(c, o int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("shipping_carriers/%d/options/%d", c, o),
	}.Do()
	return
}

func getAllFulfillments(o int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%d/fulfillments", o),
	}.Do()
	return
}

func getFulfillmentByID(o, f int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%d/fulfillments/%d", o, f),
	}.Do()
	return
}

func createFulfillment(s []byte, id int) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%d/fulfillments", id),
		Body:    bytes.NewReader(s),
	}.Do()
	return
}

func deleteFulfillment(o, f int) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("orders/%d/fulfillments/%d", o, f),
	}.Do()
	return
}
