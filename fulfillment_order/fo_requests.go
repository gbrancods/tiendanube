package fulfillmentorder

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAllFulfillmentOrders(o string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders", o),
	}.Do()
	return
}

func getFulfillmenteOrderByID(o, f string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders/%s", o, f),
	}.Do()
	return
}

func deleteFulfillmentOrder(o, f string) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders/%s", o, f),
	}.Do()
	return
}

func patchFulfillmentOrder(b []byte, o, f string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PATCH",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders/%s", o, f),
		Body:    bytes.NewReader(b),
	}.Do()
	return
}

func trackingPost(b []byte, o, f string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders/%s/tracking-events", o, f),
		Body:    bytes.NewReader(b),
	}.Do()
	return
}

func trackingPut(b []byte, o, f, t string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders/%s/tracking-events/%s", o, f, t),
		Body:    bytes.NewReader(b),
	}.Do()
	return
}

func trackingDelete(o, f, t string) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders/%s/tracking-events/%s", o, f, t),
	}.Do()
	return
}

func getAllTrackings(o, f string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders/%s/tracking-events", o, f),
	}.Do()
	return
}

func getTrackingByID(o, f, t string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("orders/%s/fulfillment-orders/%s/tracking-events/%s", o, f, t),
	}.Do()
	return
}
