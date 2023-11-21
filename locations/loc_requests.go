// A Location of the store where physical goods reside.
package locations

import (
	"bytes"
	"fmt"

	"github.com/gbrancods/tiendanube/requests"
)

func getAll() (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: "locations",
	}.Do()
	return
}

func getByID(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "GET",
		Destiny: fmt.Sprintf("locations/%s", id),
	}.Do()
	return
}

func create(l []byte) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "POST",
		Destiny: "locations",
		Body:    bytes.NewReader(l),
	}.Do()
	return
}

func update(l []byte, id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "PUT",
		Destiny: fmt.Sprintf("locations/%s", id),
		Body:    bytes.NewReader(l),
	}.Do()
	return
}

func delete(id string) (err error) {
	_, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("locations/%s", id),
	}.Do()
	return
}

func getInventoryLevels(id string) (r []byte, err error) {
	r, err = requests.Request{
		Method:  "DELETE",
		Destiny: fmt.Sprintf("locations/%s/inventory-levels", id),
	}.Do()
	return
}
